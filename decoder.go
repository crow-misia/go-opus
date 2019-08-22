/*
 * Copyright (c) 2019 Zenichi Amano
 *
 * This file is part of libsoundio, which is MIT licensed.
 * See http://opensource.org/licenses/MIT
 */

package opus

/*
#include <opus/opus.h>

// Decoder CTLs

int decoder_set_gain(OpusDecoder *ctx, opus_int32 x) {
	return opus_decoder_ctl(ctx, OPUS_SET_GAIN(x));
}

int decoder_get_gain(OpusDecoder *ctx, opus_int32 *x) {
	return opus_decoder_ctl(ctx, OPUS_GET_GAIN(x));
}

int decoder_get_last_packet_duration(OpusDecoder *ctx, opus_int32 *x) {
	return opus_decoder_ctl(ctx, OPUS_GET_LAST_PACKET_DURATION(x));
}

int decoder_get_pitch(OpusDecoder *ctx, opus_int32 *x) {
	return opus_decoder_ctl(ctx, OPUS_GET_PITCH(x));
}

// Generi CTLs

int decoder_get_reset_state(OpusDecoder *ctx) {
	return opus_decoder_ctl(ctx, OPUS_RESET_STATE);
}

int decoder_get_final_range(OpusDecoder *ctx, opus_uint32 *x) {
	return opus_decoder_ctl(ctx, OPUS_GET_FINAL_RANGE(x));
}

int decoder_get_bandwidth(OpusDecoder *ctx, opus_int32 *x) {
	return opus_decoder_ctl(ctx, OPUS_GET_BANDWIDTH(x));
}

int decoder_get_sample_rate(OpusDecoder *ctx, opus_int32 *x) {
	return opus_decoder_ctl(ctx, OPUS_GET_SAMPLE_RATE(x));
}

int decoder_set_phase_inversion_disabled(OpusDecoder *ctx, opus_int32 x) {
	return opus_decoder_ctl(ctx, OPUS_SET_PHASE_INVERSION_DISABLED(x));
}

int decoder_get_phase_inversion_disabled(OpusDecoder *ctx, opus_int32 *x) {
	return opus_decoder_ctl(ctx, OPUS_GET_PHASE_INVERSION_DISABLED(x));
}

int decoder_get_in_dtx(OpusDecoder *ctx, opus_int32 *x) {
#ifdef OPUS_GET_IN_DTX
	return opus_decoder_ctl(ctx, OPUS_GET_IN_DTX(x));
#else
	return OPUS_UNIMPLEMENTED;
#endif
}
*/
import "C"
import (
	"unsafe"
)

// Decoder is Opus decoder
type Decoder struct {
	ptr *C.struct_OpusDecoder
}

func NewDecoder(sampleRate int, channels Channel) (*Decoder, error) {
	structSize := C.opus_decoder_get_size(C.int(channels))
	tmp := make([]byte, structSize)
	ptr := (*C.OpusDecoder)(unsafe.Pointer(&tmp[0]))

	err := C.opus_decoder_init(ptr, C.opus_int32(sampleRate), C.int(channels))
	return &Decoder{
		ptr: ptr,
	}, convertToError(err)
}

func (d *Decoder) Decode(data []byte, pcm []int16, decodeFec bool) (int, error) {
	lenData := len(data)
	if lenData == 0 {
		return 0, errorNoDataSupplied
	}
	capPcm := cap(pcm)
	if capPcm == 0 {
		return 0, errorTargetBufferEmpty
	}

	n := C.opus_decode(d.ptr, (*C.uchar)(&data[0]), C.opus_int32(lenData), (*C.opus_int16)(&pcm[0]), C.int(capPcm), boolToInt(decodeFec))
	if n < 0 {
		return 0, convertToError(n)
	}
	return int(n), nil
}

func (d *Decoder) DecodeFloat(data []byte, pcm []float32, decodeFec bool) (int, error) {
	lenData := len(data)
	if lenData == 0 {
		return 0, errorNoDataSupplied
	}
	capPcm := cap(pcm)
	if capPcm == 0 {
		return 0, errorTargetBufferEmpty
	}

	n := C.opus_decode_float(d.ptr, (*C.uchar)(&data[0]), C.opus_int32(lenData), (*C.float)(&pcm[0]), C.int(capPcm), boolToInt(decodeFec))
	if n < 0 {
		return 0, convertToError(n)
	}
	return int(n), nil
}

// Decoder CTLs

func (d *Decoder) SetGain(gain int) error {
	return convertToError(C.decoder_set_gain(d.ptr, C.opus_int32(gain)))
}

func (d *Decoder) Gain() (int, error) {
	var gain C.opus_int32
	err := convertToError(C.decoder_get_gain(d.ptr, &gain))
	return int(gain), err
}

func (d *Decoder) LastPacketDuration() (int, error) {
	var duration C.opus_int32
	err := convertToError(C.decoder_get_last_packet_duration(d.ptr, &duration))
	return int(duration), err
}

func (d *Decoder) Pitch() (int, error) {
	var pitch C.opus_int32
	err := convertToError(C.decoder_get_pitch(d.ptr, &pitch))
	return int(pitch), err
}

// Generic CTLs

func (d *Decoder) ResetState() error {
	return convertToError(C.decoder_get_reset_state(d.ptr))
}

func (d *Decoder) FinalRange() (uint, error) {
	var finalRange C.opus_uint32
	err := convertToError(C.decoder_get_final_range(d.ptr, &finalRange))
	return uint(finalRange), err
}

func (d *Decoder) Bandwidth() (Bandwidth, error) {
	var bandwidth C.opus_int32
	err := convertToError(C.decoder_get_bandwidth(d.ptr, &bandwidth))
	return Bandwidth(bandwidth), err
}

func (d *Decoder) SampleRate() (int, error) {
	var sampleRate C.opus_int32
	err := convertToError(C.decoder_get_sample_rate(d.ptr, &sampleRate))
	return int(sampleRate), err
}

func (d *Decoder) SetPhaseInversionDisabled(disabled bool) error {
	return convertToError(C.decoder_set_phase_inversion_disabled(d.ptr, boolToOpusInt32(disabled)))
}

func (d *Decoder) PhaseInversionDisabled() (bool, error) {
	var disabled C.opus_int32
	err := convertToError(C.decoder_get_phase_inversion_disabled(d.ptr, &disabled))
	return opusInt32ToBool(disabled), err
}

func (d *Decoder) InDtx() (bool, error) {
	var inDtx C.opus_int32
	err := convertToError(C.decoder_get_in_dtx(d.ptr, &inDtx))
	return opusInt32ToBool(inDtx), err
}
