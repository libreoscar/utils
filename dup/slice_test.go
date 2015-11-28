package dup

import (
	"github.com/facebookgo/ensure"
	"testing"
)

func TestByteSlice(t *testing.T) {
	ensure.DeepEqual(t, ByteSlice([]byte{1, 3, 2}), []byte{1, 3, 2})

	var nilBS []byte
	ensure.DeepEqual(t, ByteSlice(nilBS), nilBS)
}

func TestInt8Slice(t *testing.T) {
	ensure.DeepEqual(t, Int8Slice([]int8{1, 3, 2}), []int8{1, 3, 2})

	var nilBS []int8
	ensure.DeepEqual(t, Int8Slice(nilBS), nilBS)
}

func TestUInt8Slice(t *testing.T) {
	ensure.DeepEqual(t, UInt8Slice([]uint8{1, 3, 2}), []uint8{1, 3, 2})

	var nilBS []uint8
	ensure.DeepEqual(t, UInt8Slice(nilBS), nilBS)
}

func TestIntSlice(t *testing.T) {
	ensure.DeepEqual(t, IntSlice([]int{1, 3, 2}), []int{1, 3, 2})

	var nilBS []int
	ensure.DeepEqual(t, IntSlice(nilBS), nilBS)
}

func TestUIntSlice(t *testing.T) {
	ensure.DeepEqual(t, UIntSlice([]uint{1, 3, 2}), []uint{1, 3, 2})

	var nilBS []uint
	ensure.DeepEqual(t, UIntSlice(nilBS), nilBS)
}

func TestInt32Slice(t *testing.T) {
	ensure.DeepEqual(t, Int32Slice([]int32{1, 3, 2}), []int32{1, 3, 2})

	var nilBS []int32
	ensure.DeepEqual(t, Int32Slice(nilBS), nilBS)
}

func TestUInt32Slice(t *testing.T) {
	ensure.DeepEqual(t, UInt32Slice([]uint32{1, 3, 2}), []uint32{1, 3, 2})

	var nilBS []uint32
	ensure.DeepEqual(t, UInt32Slice(nilBS), nilBS)
}

func TestInt64Slice(t *testing.T) {
	ensure.DeepEqual(t, Int64Slice([]int64{1, 3, 2}), []int64{1, 3, 2})

	var nilBS []int64
	ensure.DeepEqual(t, Int64Slice(nilBS), nilBS)
}

func TestUInt64Slice(t *testing.T) {
	ensure.DeepEqual(t, UInt64Slice([]uint64{1, 3, 2}), []uint64{1, 3, 2})

	var nilBS []uint64
	ensure.DeepEqual(t, UInt64Slice(nilBS), nilBS)
}

func TestFloat32Slice(t *testing.T) {
	ensure.DeepEqual(t, Float32Slice([]float32{1, 3, 2}), []float32{1, 3, 2})

	var nilBS []float32
	ensure.DeepEqual(t, Float32Slice(nilBS), nilBS)
}

func TestFloat64Slice(t *testing.T) {
	ensure.DeepEqual(t, Float64Slice([]float64{1, 3, 2}), []float64{1, 3, 2})

	var nilBS []float64
	ensure.DeepEqual(t, Float64Slice(nilBS), nilBS)
}

func TestStringSlice(t *testing.T) {
	ensure.DeepEqual(t, StringSlice([]string{"1", "333", "22"}), []string{"1", "333", "22"})

	var nilBS []string
	ensure.DeepEqual(t, StringSlice(nilBS), nilBS)
}
