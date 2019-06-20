package container

//8KB bits for value 0 ~ 2^16 - 1
type BitMapContainer struct {
	//value [4096]uint16
	value []uint16
}

func (bitmap *BitMapContainer) exists(v uint16) bool {
	index := v / 16
	var mask uint16 = 1 << (v % 16)

	return (bitmap.value[index] & mask) != 0
}

func (bitmap *BitMapContainer) add(v uint16) bool {
	index := v / 16
	var mask uint16 = 1 << (v % 16)

	ret := (bitmap.value[index] & mask) == 0

	bitmap.value[index] |= mask

	return ret
}

func (bitmap *BitMapContainer) del(v uint16) bool {
	index := v / 16
	var mask uint16 = 1 << (v % 16)

	ret := (bitmap.value[index] & mask) != 0

	bitmap.value[index] &= ^mask

	return ret
}

func (bitmap *BitMapContainer) convert(target CONTAINTER_TYPE) *Container {
	switch target {
	case BitmapContainerType:
	case SortedArrayContainerType:
	case RunContainerType:
	}
	return nil
}
