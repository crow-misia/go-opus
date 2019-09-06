/*
 * Copyright (c) 2019 Zenichi Amano
 *
 * This file is part of go-opus, which is MIT licensed.
 * See http://opensource.org/licenses/MIT
 */

package opus

/*
#include <opus/opus.h>
*/
import "C"

const (
	BitrateAuto int32 = C.OPUS_AUTO
	BitrateMax  int32 = C.OPUS_BITRATE_MAX
)

type Application int32

const (
	// ApplicationVoip defined "Best for most VoIP/videoconference applications where listening quality and intelligibility matter most"
	ApplicationVoip Application = C.OPUS_APPLICATION_VOIP
	// ApplicationAudio defined "Best for broadcast/high-fidelity application where the decoded audio should be as close as possible to the input"
	ApplicationAudio = C.OPUS_APPLICATION_AUDIO
	// ApplicationRestrictedLowdelay defined "Only use when lowest-achievable latency is what matters most. Voice-optimized modes cannot be used"
	ApplicationRestrictedLowdelay = C.OPUS_APPLICATION_RESTRICTED_LOWDELAY
)

type Vbr int32

const (
	VbrHard    = Vbr(0) // Hard VBR
	VbrDefault = Vbr(1) // VBR (default)
)

type Channel int32

const (
	ChannelMono   = Channel(1) // Mono
	ChannelStereo = Channel(2) // Stereo
)

type Signal int32

const (
	SignalAuto  = Signal(C.OPUS_AUTO)         // Auto/default setting
	SignalVoice = Signal(C.OPUS_SIGNAL_VOICE) // Signal being encoded is voice
	SignalMusic = Signal(C.OPUS_SIGNAL_MUSIC) // Signal being encoded is music
)

type Bandwidth int32

const (
	BandwidthAuto          = Bandwidth(C.OPUS_AUTO)                    // Auto/default setting
	BandwidthNarrowband    = Bandwidth(C.OPUS_BANDWIDTH_NARROWBAND)    // 4 kHz bandpass
	BandwidthMediumBand    = Bandwidth(C.OPUS_BANDWIDTH_MEDIUMBAND)    // 6 kHz bandpass
	BandwidthWideband      = Bandwidth(C.OPUS_BANDWIDTH_WIDEBAND)      // 8 kHz bandpass
	BandwidthSuperwideband = Bandwidth(C.OPUS_BANDWIDTH_SUPERWIDEBAND) // 12 kHz bandpass
	BandwidthFullband      = Bandwidth(C.OPUS_BANDWIDTH_FULLBAND)      // 20 kHz bandpass
)

type Framesize int32

const (
	FramesizeArg   = Framesize(C.OPUS_FRAMESIZE_ARG)    // Select frame size from the argument (default)
	Framesize25ms  = Framesize(C.OPUS_FRAMESIZE_2_5_MS) // Use 2.5 ms frames
	Framesize5ms   = Framesize(C.OPUS_FRAMESIZE_5_MS)   // Use 5 ms frames
	Framesize10ms  = Framesize(C.OPUS_FRAMESIZE_10_MS)  // Use 10 ms frames
	Framesize20ms  = Framesize(C.OPUS_FRAMESIZE_20_MS)  // Use 20 ms frames
	Framesize40ms  = Framesize(C.OPUS_FRAMESIZE_40_MS)  // Use 40 ms frames
	Framesize60ms  = Framesize(C.OPUS_FRAMESIZE_60_MS)  // Use 60 ms frames
	Framesize80ms  = Framesize(C.OPUS_FRAMESIZE_80_MS)  // Use 80 ms frames
	Framesize100ms = Framesize(C.OPUS_FRAMESIZE_100_MS) // Use 100 ms frames
	Framesize120ms = Framesize(C.OPUS_FRAMESIZE_120_MS) // Use 120 ms frames
)
