package container

//8KB bits for Value 0 ~ 2^16 - 1
type BitMapContainer struct {
	Value *[4096]uint16
	//Value []uint16
}

func MakeBitMapContainer() *BitMapContainer {
	ret := new(BitMapContainer)
	ret.Value = &[4096]uint16{}
	return ret
}

func (bitmap *BitMapContainer) Exists(v uint16) bool {
	index := v / 16
	var mask uint16 = 1 << (v % 16)

	return (bitmap.Value[index] & mask) != 0
}

func (bitmap *BitMapContainer) Add(v uint16) bool {
	index := v / 16
	var mask uint16 = 1 << (v % 16)

	ret := (bitmap.Value[index] & mask) == 0

	bitmap.Value[index] |= mask

	return ret
}

func (bitmap *BitMapContainer) Del(v uint16) bool {
	index := v / 16
	var mask uint16 = 1 << (v % 16)

	ret := (bitmap.Value[index] & mask) != 0

	bitmap.Value[index] &= ^mask

	return ret
}

func (bitmap *BitMapContainer) convert(target TypeContainer) *Container {
	switch target {
	case BitmapContainerType:
	case SortedArrayContainerType:
	case RunContainerType:
	}
	return nil
}
