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
import "errors"

// Error is Opus error
type Error int32

const (
	ErrorOk             = Error(C.OPUS_OK)               // No error
	ErrorBadArg         = Error(C.OPUS_BAD_ARG)          // One or more invalid/out of range arguments
	ErrorBufferTooSmall = Error(C.OPUS_BUFFER_TOO_SMALL) // The mode struct passed is invalid
	ErrorInternalError  = Error(C.OPUS_INTERNAL_ERROR)   // An internal error was detected
	ErrorInvalidPacket  = Error(C.OPUS_INVALID_PACKET)   // The compressed data passed is corrupted
	ErrorUnimplemented  = Error(C.OPUS_UNIMPLEMENTED)    // Invalid/unsupported request number
	ErrorInvalidState   = Error(C.OPUS_INVALID_STATE)    // An encoder or decoder structure is invalid or already freed
	ErrorAllocFail      = Error(C.OPUS_ALLOC_FAIL)       // Memory allocation has failed
)

func (e Error) Error() string {
	return C.GoString(C.opus_strerror(C.int(e)))
}

func convertToError(err C.int) error {
	if err == C.OPUS_OK {
		return nil
	}
	return Error(err)
}

var errorNoDataSupplied = errors.New("no data supplied")
var errorTargetBufferEmpty = errors.New("target buffer empty")
