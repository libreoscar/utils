package dup

import (
	"github.com/facebookgo/ensure"
	"testing"
)

func TestSliceDup(t *testing.T) {
	ensure.DeepEqual(t, ByteSlice([]byte{1, 3, 2}), []byte{1, 3, 2})
}
