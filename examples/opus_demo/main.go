/*
 * Copyright (c) 2019 Zenichi Amano
 *
 * This file is part of libsoundio, which is MIT licensed.
 * See http://opensource.org/licenses/MIT
 */

package main

import (
	"encoding/binary"
	"fmt"
	soundio "github.com/crow-misia/go-libsoundio"
	"github.com/crow-misia/go-opus"
	"log"
	"math"
	"os"
)

const (
	sampleRate = 48000
	freq       = 440.0
	duration   = 120
	frameSize  = sampleRate * duration / 1000
)

var frameOffset = 0

func main() {
	err := realMain()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func realMain() error {
	log.Println(opus.Version())

	encoder, err := opus.NewEncoder(48000, opus.ChannelMono, opus.ApplicationVoip)
	if err != nil {
		return err
	}

	mono := make([]int16, frameSize*100)
	createSine(mono, sampleRate, freq)

	err = encoder.SetMaxBandwidth(opus.BandwidthFullband)
	if err != nil {
		return err
	}
	err = encoder.SetBitrate(128000)
	if err != nil {
		return err
	}

	encodedData := make([][]byte, 100)
	for i, idx := 0, 0; i < len(mono); i += frameSize {
		data := make([]byte, 1000)
		n, err := encoder.Encode(mono[i:i+frameSize], data)
		if err != nil {
			return err
		}
		encodedData[idx] = data[:n]
		idx++
	}

	decoder, err := opus.NewDecoder(sampleRate, opus.ChannelMono)
	if err != nil || decoder == nil {
		return err
	}
	newPcm := make([]float32, frameSize*100+1000)
	pcmOffset := 0
	for _, data := range encodedData {
		n, err := decoder.DecodeFloat(data, newPcm[pcmOffset:], false)
		if err != nil {
			return err
		}
		pcmOffset += n
	}
	if frameSize*100 != pcmOffset {
		return fmt.Errorf("decode pcm data size invalid %d != %d", frameSize*100, pcmOffset)
	}
	newPcm = newPcm[:pcmOffset]

	s := soundio.Create()

	err = s.Connect()
	if err != nil {
		return fmt.Errorf("error connecting: %s", err)
	}

	s.FlushEvents()

	defaultOutputDeviceIndex := s.DefaultOutputDeviceIndex()

	device := s.OutputDevice(defaultOutputDeviceIndex)

	outStream := device.NewOutStream()
	layout := soundio.ChannelLayoutGetDefault(2)
	outStream.SetLayout(layout)
	outStream.SetFormat(soundio.FormatFloat32LE)
	outStream.SetWriteCallback(func(stream *soundio.OutStream, frameCountMix int, frameCountMax int) {
		var areas *soundio.ChannelAreas

		framesLeft := frameCountMax
		for framesLeft > 0 {
			frameCount := framesLeft
			if frameCount <= 0 {
				break
			}

			areas, err = stream.BeginWrite(&frameCount)
			if err != nil {
				panic(err)
			}
			if frameCount <= 0 {
				break
			}
			if areas == nil {
				break
			}

			for frame := 0; frame < frameCount; frame++ {
				sample := newPcm[frameOffset]
				frameOffset = (frameOffset + 1) % int(pcmOffset)

				for channel := 0; channel < 2; channel++ {
					buffer := areas.Buffer(channel, frame)
					bites := math.Float32bits(sample)
					binary.LittleEndian.PutUint32(buffer, bites)
				}
			}

			err = outStream.EndWrite()
			if err != nil {
				panic(err)
			}
			framesLeft -= frameCount
		}
	})

	err = outStream.Open()
	if err != nil {
		return fmt.Errorf("error opening: %s", err)
	}
	defer outStream.Destroy()

	channels := layout.Channels()
	channelCount := layout.ChannelCount()
	for i := 0; i < channelCount; i++ {
		log.Printf("      Channel ID = %d, Name = %s", channels[i], channels[i])
	}

	err = outStream.Start()
	if err != nil {
		return fmt.Errorf("error opening: %s", err)
	}

	for {
		s.WaitEvents()
	}
}

func createSine(buffer []int16, sampleRate int, freq float64) {
	factor := 2.0 * math.Pi * freq / float64(sampleRate)
	max := float64(math.MaxInt16 - 1)
	for i := range buffer {
		buffer[i] = int16(math.Sin(float64(i)*factor) * max)
	}
}
