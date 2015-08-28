package dup

func ByteSlice(data []byte) (ret []byte) {
	if data == nil {
		return
	}
	ret = make([]byte, len(data))
	copy(ret, data)
	return
}

func IntSlice(data []int) (ret []int) {
	if data == nil {
		return
	}
	ret = make([]int, len(data))
	copy(ret, data)
	return
}

func UIntSlice(data []uint) (ret []uint) {
	if data == nil {
		return
	}
	ret = make([]uint, len(data))
	copy(ret, data)
	return
}

func Int32Slice(data []int32) (ret []int32) {
	if data == nil {
		return
	}
	ret = make([]int32, len(data))
	copy(ret, data)
	return
}

func UInt32Slice(data []uint32) (ret []uint32) {
	if data == nil {
		return
	}
	ret = make([]uint32, len(data))
	copy(ret, data)
	return
}

func Int64Slice(data []int64) (ret []int64) {
	if data == nil {
		return
	}
	ret = make([]int64, len(data))
	copy(ret, data)
	return
}

func UInt64Slice(data []uint64) (ret []uint64) {
	if data == nil {
		return
	}
	ret = make([]uint64, len(data))
	copy(ret, data)
	return
}

func StringSlice(data []string) (ret []string) {
	if data == nil {
		return
	}
	ret = make([]string, len(data))
	copy(ret, data)
	return
}
