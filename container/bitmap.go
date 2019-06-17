package container

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