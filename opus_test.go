/*
 * Copyright (c) 2019 Zenichi Amano
 *
 * This file is part of libsoundio, which is MIT licensed.
 * See http://opensource.org/licenses/MIT
 */

package opus

import (
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	if ver := Version(); !strings.HasPrefix(ver, "libopus 1.") {
		t.Errorf("Unexpected linked libopus version: %s", ver)
	}
}

func TestErrorString(t *testing.T) {
	// I scooped this -1 up from opus_defines.h, it's OPUS_BAD_ARG. Not pretty,
	// but it's better than not testing at all. Again, accessing #defines from
	// CGO is not possible.
	if ErrorBadArg.Error() != "invalid argument" {
		t.Errorf("Expected \"invalid argument\" error message for error code %d: %v", ErrorBadArg, ErrorBadArg)
	}
}
