package sort

import (
	"github.com/facebookgo/ensure"
	"testing"
)

func TestBytes(t *testing.T) {
	a := []byte{1, 3, 3, 7, 2}
	{
		cp := make([]byte, len(a))
		copy(cp, a)
		Bytes(cp, false)
		ensure.DeepEqual(t, cp, []byte{1, 2, 3, 3, 7})
	}
	{
		cp := make([]byte, len(a))
		copy(cp, a)
		Bytes(cp, true)
		ensure.DeepEqual(t, cp, []byte{7, 3, 3, 2, 1})
	}
}

func TestInt8s(t *testing.T) {
	a := []int8{1, 3, -4, 3, 7, 2}
	{
		cp := make([]int8, len(a))
		copy(cp, a)
		Int8s(cp, false)
		ensure.DeepEqual(t, cp, []int8{-4, 1, 2, 3, 3, 7})
	}
	{
		cp := make([]int8, len(a))
		copy(cp, a)
		Int8s(cp, true)
		ensure.DeepEqual(t, cp, []int8{7, 3, 3, 2, 1, -4})
	}
}

func TestUInt8s(t *testing.T) {
	a := []uint8{1, 3, 3, 7, 2}
	{
		cp := make([]uint8, len(a))
		copy(cp, a)
		UInt8s(cp, false)
		ensure.DeepEqual(t, cp, []uint8{1, 2, 3, 3, 7})
	}
	{
		cp := make([]uint8, len(a))
		copy(cp, a)
		UInt8s(cp, true)
		ensure.DeepEqual(t, cp, []uint8{7, 3, 3, 2, 1})
	}
}

func TestInts(t *testing.T) {
	a := []int{1, 3, -4, 3, 7, 2}
	{
		cp := make([]int, len(a))
		copy(cp, a)
		Ints(cp, false)
		ensure.DeepEqual(t, cp, []int{-4, 1, 2, 3, 3, 7})
	}
	{
		cp := make([]int, len(a))
		copy(cp, a)
		Ints(cp, true)
		ensure.DeepEqual(t, cp, []int{7, 3, 3, 2, 1, -4})
	}
}

func TestUInts(t *testing.T) {
	a := []uint{1, 3, 4, 3, 7, 2}
	{
		cp := make([]uint, len(a))
		copy(cp, a)
		UInts(cp, false)
		ensure.DeepEqual(t, cp, []uint{1, 2, 3, 3, 4, 7})
	}
	{
		cp := make([]uint, len(a))
		copy(cp, a)
		UInts(cp, true)
		ensure.DeepEqual(t, cp, []uint{7, 4, 3, 3, 2, 1})
	}
}

func TestInt32s(t *testing.T) {
	a := []int32{1, 3, -4, 3, 7, 2}
	{
		cp := make([]int32, len(a))
		copy(cp, a)
		Int32s(cp, false)
		ensure.DeepEqual(t, cp, []int32{-4, 1, 2, 3, 3, 7})
	}
	{
		cp := make([]int32, len(a))
		copy(cp, a)
		Int32s(cp, true)
		ensure.DeepEqual(t, cp, []int32{7, 3, 3, 2, 1, -4})
	}
}

func TestUInt32s(t *testing.T) {
	a := []uint32{1, 3, 4, 3, 7, 2}
	{
		cp := make([]uint32, len(a))
		copy(cp, a)
		UInt32s(cp, false)
		ensure.DeepEqual(t, cp, []uint32{1, 2, 3, 3, 4, 7})
	}
	{
		cp := make([]uint32, len(a))
		copy(cp, a)
		UInt32s(cp, true)
		ensure.DeepEqual(t, cp, []uint32{7, 4, 3, 3, 2, 1})
	}
}

func TestInt64s(t *testing.T) {
	a := []int64{1, 3, -4, 3, 7, 2}
	{
		cp := make([]int64, len(a))
		copy(cp, a)
		Int64s(cp, false)
		ensure.DeepEqual(t, cp, []int64{-4, 1, 2, 3, 3, 7})
	}
	{
		cp := make([]int64, len(a))
		copy(cp, a)
		Int64s(cp, true)
		ensure.DeepEqual(t, cp, []int64{7, 3, 3, 2, 1, -4})
	}
}

func TestUInt64s(t *testing.T) {
	a := []uint64{1, 3, 4, 3, 7, 2}
	{
		cp := make([]uint64, len(a))
		copy(cp, a)
		UInt64s(cp, false)
		ensure.DeepEqual(t, cp, []uint64{1, 2, 3, 3, 4, 7})
	}
	{
		cp := make([]uint64, len(a))
		copy(cp, a)
		UInt64s(cp, true)
		ensure.DeepEqual(t, cp, []uint64{7, 4, 3, 3, 2, 1})
	}
}

func TestFloat32s(t *testing.T) {
	a := []float32{1.0, 3.0, -4.0, 3.0, 7.0, 2.0}
	{
		cp := make([]float32, len(a))
		copy(cp, a)
		Float32s(cp, false)
		ensure.DeepEqual(t, cp, []float32{-4.0, 1.0, 2.0, 3.0, 3.0, 7.0})
	}
	{
		cp := make([]float32, len(a))
		copy(cp, a)
		Float32s(cp, true)
		ensure.DeepEqual(t, cp, []float32{7.0, 3.0, 3.0, 2.0, 1.0, -4.0})
	}
}

func TestFloat64s(t *testing.T) {
	a := []float64{1.0, 3.0, -4.0, 3.0, 7.0, 2.0}
	{
		cp := make([]float64, len(a))
		copy(cp, a)
		Float64s(cp, false)
		ensure.DeepEqual(t, cp, []float64{-4.0, 1.0, 2.0, 3.0, 3.0, 7.0})
	}
	{
		cp := make([]float64, len(a))
		copy(cp, a)
		Float64s(cp, true)
		ensure.DeepEqual(t, cp, []float64{7.0, 3.0, 3.0, 2.0, 1.0, -4.0})
	}
}

func TestStrings(t *testing.T) {
	a := []string{"a", "aa", "z", "bbb", "ba"}
	{
		cp := make([]string, len(a))
		copy(cp, a)
		Strings(cp, false)
		ensure.DeepEqual(t, cp, []string{"a", "aa", "ba", "bbb", "z"})
	}
	{
		cp := make([]string, len(a))
		copy(cp, a)
		Strings(cp, true)
		ensure.DeepEqual(t, cp, []string{"z", "bbb", "ba", "aa", "a"})
	}
}
