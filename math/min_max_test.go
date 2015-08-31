package math

import (
	"github.com/facebookgo/ensure"
	"testing"
)

func TestMinMax(t *testing.T) {
	ensure.DeepEqual(t, ByteMin(byte(4), byte(6)), byte(4))
	ensure.DeepEqual(t, ByteMax(byte(4), byte(6)), byte(6))

	ensure.DeepEqual(t, Int8Min(int8(-4), int8(6)), int8(-4))
	ensure.DeepEqual(t, Int8Max(int8(-4), int8(6)), int8(6))

	ensure.DeepEqual(t, UInt8Min(uint8(4), uint8(6)), uint8(4))
	ensure.DeepEqual(t, UInt8Max(uint8(4), uint8(6)), uint8(6))

	ensure.DeepEqual(t, IntMin(int(-4), int(6)), int(-4))
	ensure.DeepEqual(t, IntMax(int(-4), int(6)), int(6))

	ensure.DeepEqual(t, UIntMin(uint(4), uint(6)), uint(4))
	ensure.DeepEqual(t, UIntMax(uint(4), uint(6)), uint(6))

	ensure.DeepEqual(t, Int32Min(int32(-4), int32(6)), int32(-4))
	ensure.DeepEqual(t, Int32Max(int32(-4), int32(6)), int32(6))

	ensure.DeepEqual(t, UInt32Min(uint32(4), uint32(6)), uint32(4))
	ensure.DeepEqual(t, UInt32Max(uint32(4), uint32(6)), uint32(6))

	ensure.DeepEqual(t, Int64Min(int64(-4), int64(6)), int64(-4))
	ensure.DeepEqual(t, Int64Max(int64(-4), int64(6)), int64(6))

	ensure.DeepEqual(t, UInt64Min(uint64(4), uint64(6)), uint64(4))
	ensure.DeepEqual(t, UInt64Max(uint64(4), uint64(6)), uint64(6))

	ensure.DeepEqual(t, Float32Min(float32(-4), float32(6)), float32(-4))
	ensure.DeepEqual(t, Float32Max(float32(-4), float32(6)), float32(6))

	ensure.DeepEqual(t, Float64Min(float64(-4), float64(6)), float64(-4))
	ensure.DeepEqual(t, Float64Max(float64(-4), float64(6)), float64(6))

	ensure.DeepEqual(t, StringMin("abc", "def"), "abc")
	ensure.DeepEqual(t, StringMax("abc", "def"), "def")
}
