package container

import (
	"errors"
	"unsafe"
)

//8KB bits for Value 0 ~ 2^16 - 1
type BitMapContainer struct {
	Value *[4096]uint16
}

func MakeBitMapContainer() *BitMapContainer {
	ret := new(BitMapContainer)
	ret.Value = &[4096]uint16{}
	return ret
}

func FromBinaryArray(input []byte) (*BitMapContainer, error) {
	if len(input) != 8192 {
		return nil, errors.New("input must be of length 8192")
	}

	ret := new(BitMapContainer)
	ret.Value = (*[4096]uint16)(unsafe.Pointer(&input[0]))
	return ret, nil
}

func (bitmap *BitMapContainer) ToBinaryArray() []byte {
	bytes := (*[8192]byte)(unsafe.Pointer(&(bitmap.Value[0])))
	return bytes[:]
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
