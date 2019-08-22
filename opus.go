/*
 * Copyright (c) 2019 Zenichi Amano
 *
 * This file is part of libsoundio, which is MIT licensed.
 * See http://opensource.org/licenses/MIT
 */

package opus

/*
#cgo LDFLAGS: -lopus -lm
#include <opus/opus.h>
*/
import "C"

func Version() string {
	return C.GoString(C.opus_get_version_string())
}
