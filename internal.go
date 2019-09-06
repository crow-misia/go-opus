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

func opusInt32ToBool(v C.opus_int32) bool {
	return v == 1
}

func boolToInt(v bool) C.int {
	if v {
		return C.int(1)
	}
	return C.int(0)
}

func boolToOpusInt32(v bool) C.opus_int32 {
	if v {
		return C.opus_int32(1)
	}
	return C.opus_int32(0)
}
