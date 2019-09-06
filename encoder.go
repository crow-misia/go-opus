/*
 * Copyright (c) 2019 Zenichi Amano
 *
 * This file is part of go-opus, which is MIT licensed.
 * See http://opensource.org/licenses/MIT
 */

package opus

/*
#include <opus/opus.h>

// Encoder CTLs

int encoder_set_complexity(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_COMPLEXITY(x));
}

int encoder_get_complexity(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_COMPLEXITY(x));
}

int encoder_set_bitrate(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_BITRATE(x));
}

int encoder_get_bitrate(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_BITRATE(x));
}

int encoder_set_vbr(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_VBR(x));
}

int encoder_get_vbr(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_VBR(x));
}

int encoder_set_vbr_constraint(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_VBR_CONSTRAINT(x));
}

int encoder_get_vbr_constraint(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_VBR_CONSTRAINT(x));
}

int encoder_set_force_channels(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_FORCE_CHANNELS(x));
}

int encoder_get_force_channels(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_FORCE_CHANNELS(x));
}

int encoder_set_max_bandwidth(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_MAX_BANDWIDTH(x));
}

int encoder_get_max_bandwidth(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_MAX_BANDWIDTH(x));
}

int encoder_set_bandwidth(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_BANDWIDTH(x));
}

int encoder_set_signal(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_SIGNAL(x));
}

int encoder_get_signal(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_SIGNAL(x));
}

int encoder_set_application(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_APPLICATION(x));
}

int encoder_get_application(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_APPLICATION(x));
}

int encoder_get_lookahead(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_LOOKAHEAD(x));
}

int encoder_set_inband_fec(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_INBAND_FEC(x));
}

int encoder_get_inband_fec(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_INBAND_FEC(x));
}

int encoder_set_packet_loss_perc(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_PACKET_LOSS_PERC(x));
}

int encoder_get_packet_loss_perc(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_PACKET_LOSS_PERC(x));
}

int encoder_set_dtx(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_DTX(x));
}

int encoder_get_dtx(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_DTX(x));
}

int encoder_set_lsb_depth(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_LSB_DEPTH(x));
}

int encoder_get_lsb_depth(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_LSB_DEPTH(x));
}

int encoder_set_expert_frame_duration(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_EXPERT_FRAME_DURATION(x));
}

int encoder_get_expert_frame_duration(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_EXPERT_FRAME_DURATION(x));
}

int encoder_set_prediction_disabled(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_PREDICTION_DISABLED(x));
}

int encoder_get_prediction_disabled(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_PREDICTION_DISABLED(x));
}

// Generic CTLs

int encoder_get_reset_state(OpusEncoder *ctx) {
	return opus_encoder_ctl(ctx, OPUS_RESET_STATE);
}

int encoder_get_final_range(OpusEncoder *ctx, opus_uint32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_FINAL_RANGE(x));
}

int encoder_get_bandwidth(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_BANDWIDTH(x));
}

int encoder_get_sample_rate(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_SAMPLE_RATE(x));
}

int encoder_set_phase_inversion_disabled(OpusEncoder *ctx, opus_int32 x) {
	return opus_encoder_ctl(ctx, OPUS_SET_PHASE_INVERSION_DISABLED(x));
}

int encoder_get_phase_inversion_disabled(OpusEncoder *ctx, opus_int32 *x) {
	return opus_encoder_ctl(ctx, OPUS_GET_PHASE_INVERSION_DISABLED(x));
}

int encoder_get_in_dtx(OpusEncoder *ctx, opus_int32 *x) {
#ifdef OPUS_GET_IN_DTX
	return opus_encoder_ctl(ctx, OPUS_GET_IN_DTX(x));
#else
	return OPUS_UNIMPLEMENTED;
#endif
}
*/
import "C"
import "unsafe"

// Encoder is Opus encoder
type Encoder struct {
	// OpusEncoder pointer
	ptr *C.struct_OpusEncoder
	// channel count
	channels Channel
}

func NewEncoder(sampleRate int, channels Channel, application Application) (*Encoder, error) {
	structSize := C.opus_encoder_get_size(C.int(channels))
	tmp := make([]byte, structSize)
	ptr := (*C.OpusEncoder)(unsafe.Pointer(&tmp[0]))

	err := C.opus_encoder_init(ptr, C.opus_int32(sampleRate), C.int(channels), C.int(application))
	return &Encoder{
		ptr:      ptr,
		channels: channels,
	}, convertToError(err)
}

func (e *Encoder) Encode(pcm []int16, data []byte) (int, error) {
	lenPcm := len(pcm)
	if lenPcm == 0 {
		return 0, errorNoDataSupplied
	}
	capData := cap(data)
	if capData == 0 {
		return 0, errorTargetBufferEmpty
	}

	frameSize := lenPcm / int(e.channels)
	n := C.opus_encode(e.ptr, (*C.opus_int16)(&pcm[0]), C.int(frameSize), (*C.uchar)(&data[0]), C.opus_int32(capData))
	if n < 0 {
		return 0, convertToError(n)
	}
	return int(n), nil
}

func (e *Encoder) EncodeFloat(pcm []float32, data []byte) (int, error) {
	lenPcm := len(pcm)
	if lenPcm == 0 {
		return 0, errorNoDataSupplied
	}
	capData := cap(data)
	if capData == 0 {
		return 0, errorTargetBufferEmpty
	}

	frameSize := lenPcm / int(e.channels)
	n := C.opus_encode_float(e.ptr, (*C.float)(&pcm[0]), C.int(frameSize), (*C.uchar)(&data[0]), C.opus_int32(capData))
	if n < 0 {
		return 0, convertToError(n)
	}
	return int(n), nil
}

// Encoder CTLs

func (e *Encoder) SetComplexity(complexity int) error {
	return convertToError(C.encoder_set_complexity(e.ptr, C.opus_int32(complexity)))
}

func (e *Encoder) Complexity() (int, error) {
	var complexity C.opus_int32
	err := convertToError(C.encoder_get_complexity(e.ptr, &complexity))
	return int(complexity), err
}

func (e *Encoder) SetBitrate(bitrate int) error {
	return convertToError(C.encoder_set_bitrate(e.ptr, C.opus_int32(bitrate)))
}

func (e *Encoder) Bitrate() (int, error) {
	var bitrate C.opus_int32
	err := convertToError(C.encoder_get_bitrate(e.ptr, &bitrate))
	return int(bitrate), err
}

func (e *Encoder) SetVbr(vbr Vbr) error {
	return convertToError(C.encoder_set_vbr(e.ptr, C.opus_int32(vbr)))
}

func (e *Encoder) Vbr() (Vbr, error) {
	var vbr C.opus_int32
	err := convertToError(C.encoder_get_vbr(e.ptr, &vbr))
	return Vbr(vbr), err
}

func (e *Encoder) SetVbrConstraint(constrained bool) error {
	return convertToError(C.encoder_set_vbr_constraint(e.ptr, boolToOpusInt32(constrained)))
}

func (e *Encoder) VbrConstraint() (bool, error) {
	var constrained C.opus_int32
	err := convertToError(C.encoder_get_vbr_constraint(e.ptr, &constrained))
	return opusInt32ToBool(constrained), err
}

func (e *Encoder) SetForceChannels(channels Channel) error {
	return convertToError(C.encoder_set_force_channels(e.ptr, C.opus_int32(channels)))
}

func (e *Encoder) ForceChannels() (Channel, error) {
	var channels C.opus_int32
	err := convertToError(C.encoder_get_force_channels(e.ptr, &channels))
	return Channel(channels), err
}

func (e *Encoder) SetMaxBandwidth(bandwidth Bandwidth) error {
	return convertToError(C.encoder_set_max_bandwidth(e.ptr, C.opus_int32(bandwidth)))
}

func (e *Encoder) MaxBandwidth() (Bandwidth, error) {
	var bandwidth C.opus_int32
	err := convertToError(C.encoder_get_max_bandwidth(e.ptr, &bandwidth))
	return Bandwidth(bandwidth), err
}

func (e *Encoder) SetBandwidth(bandwidth Bandwidth) error {
	return convertToError(C.encoder_set_bandwidth(e.ptr, C.opus_int32(bandwidth)))
}

func (e *Encoder) SetSignal(signal Signal) error {
	return convertToError(C.encoder_set_signal(e.ptr, C.opus_int32(signal)))
}

func (e *Encoder) Signal() (Signal, error) {
	var signal C.opus_int32
	err := convertToError(C.encoder_get_signal(e.ptr, &signal))
	return Signal(signal), err
}

func (e *Encoder) SetApplication(application Application) error {
	return convertToError(C.encoder_set_application(e.ptr, C.opus_int32(application)))
}

func (e *Encoder) Application() (Application, error) {
	var application C.opus_int32
	err := convertToError(C.encoder_get_application(e.ptr, &application))
	return Application(application), err
}

func (e *Encoder) Lookahead() (int, error) {
	var lookahead C.opus_int32
	err := convertToError(C.encoder_get_lookahead(e.ptr, &lookahead))
	return int(lookahead), err
}

func (e *Encoder) SetInbandFec(enabled bool) error {
	return convertToError(C.encoder_set_inband_fec(e.ptr, boolToOpusInt32(enabled)))
}

func (e *Encoder) InbandFec() (bool, error) {
	var fec C.opus_int32
	err := convertToError(C.encoder_get_inband_fec(e.ptr, &fec))
	return opusInt32ToBool(fec), err
}

func (e *Encoder) SetPacketLossPerc(perc int) error {
	return convertToError(C.encoder_set_packet_loss_perc(e.ptr, C.opus_int32(perc)))
}

func (e *Encoder) PacketLossPerc() (int, error) {
	var perc C.opus_int32
	err := convertToError(C.encoder_get_packet_loss_perc(e.ptr, &perc))
	return int(perc), err
}

func (e *Encoder) SetDtx(enabled bool) error {
	return convertToError(C.encoder_set_dtx(e.ptr, boolToOpusInt32(enabled)))
}

func (e *Encoder) Dtx() (bool, error) {
	var dtx C.opus_int32
	err := convertToError(C.encoder_get_dtx(e.ptr, &dtx))
	return opusInt32ToBool(dtx), err
}

func (e *Encoder) SetLsbDepth(depth int) error {
	return convertToError(C.encoder_set_lsb_depth(e.ptr, C.opus_int32(depth)))
}

func (e *Encoder) LsbDepth() (int, error) {
	var depth C.opus_int32
	err := convertToError(C.encoder_get_lsb_depth(e.ptr, &depth))
	return int(depth), err
}

func (e *Encoder) SetExpertFrameDuration(duration Framesize) error {
	return convertToError(C.encoder_set_expert_frame_duration(e.ptr, C.opus_int32(duration)))
}

func (e *Encoder) ExpertFrameDuration() (Framesize, error) {
	var duration C.opus_int32
	err := convertToError(C.encoder_get_expert_frame_duration(e.ptr, &duration))
	return Framesize(duration), err
}

func (e *Encoder) SetPredictionDisabled(disabled bool) error {
	return convertToError(C.encoder_set_expert_frame_duration(e.ptr, boolToOpusInt32(disabled)))
}

func (e *Encoder) PredictionDisabled() (bool, error) {
	var disabled C.opus_int32
	err := convertToError(C.encoder_get_expert_frame_duration(e.ptr, &disabled))
	return opusInt32ToBool(disabled), err
}

// Generic CTLs

func (e *Encoder) ResetState() error {
	return convertToError(C.encoder_get_reset_state(e.ptr))
}

func (e *Encoder) FinalRange() (uint, error) {
	var finalRange C.opus_uint32
	err := convertToError(C.encoder_get_final_range(e.ptr, &finalRange))
	return uint(finalRange), err
}

func (e *Encoder) Bandwidth() (Bandwidth, error) {
	var bandwidth C.opus_int32
	err := convertToError(C.encoder_get_bandwidth(e.ptr, &bandwidth))
	return Bandwidth(bandwidth), err
}

func (e *Encoder) SampleRate() (int, error) {
	var sampleRate C.opus_int32
	err := convertToError(C.encoder_get_sample_rate(e.ptr, &sampleRate))
	return int(sampleRate), err
}

func (e *Encoder) SetPhaseInversionDisabled(disabled bool) error {
	return convertToError(C.encoder_set_phase_inversion_disabled(e.ptr, boolToOpusInt32(disabled)))
}

func (e *Encoder) PhaseInversionDisabled() (bool, error) {
	var disabled C.opus_int32
	err := convertToError(C.encoder_get_phase_inversion_disabled(e.ptr, &disabled))
	return opusInt32ToBool(disabled), err
}

func (e *Encoder) InDtx() (bool, error) {
	var inDtx C.opus_int32
	err := convertToError(C.encoder_get_in_dtx(e.ptr, &inDtx))
	return opusInt32ToBool(inDtx), err
}
