package container

//8KB bits for value 0 ~ 2^16 - 1
type BitMap struct {
	value [4096]uint16
}

const limit = 4096

func (bitmap *BitMap) exists(v uint16) bool {
	if v >= limit {
		return false
	}

	index := v / 16
	var mask uint16 = 1 << (v % 16)

	return (bitmap.value[index] & mask) != 0
}

func (bitmap *BitMap) add(v uint16) bool {
	if v >= limit {
		return false
	}

	index := v / 16
	var mask uint16 = 1 << (v % 16)

	ret := (bitmap.value[index] & mask) == 0

	bitmap.value[index] |= mask

	return ret
}

func (bitmap *BitMap) remove(v uint16) bool {
	if v >= limit {
		return false
	}

	index := v / 16
	var mask uint16 = 1 << (v % 16)

	ret := (bitmap.value[index] & mask) == 1

	bitmap.value[index] &= ^mask

	return ret
}

func (bitmap *BitMap) convert(target CONTAINTER_TYPE) *Container {
	switch target {
	case BITMAP:
	case SORTED_ARRAY:
	case RUN:
	}
	return nil
}
