package container

import (
	"errors"
	"github.com/Nov11/naive-implementation-of-roaringbitmap/util"
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

func (bitmap *BitMapContainer) countRuns() uint16 {
	var ret uint16 = 0

	var w uint16 = 0
	var n uint16 = 0

	for i := 0; i+1 < len(bitmap.Value); i++ {
		w = bitmap.Value[i]
		n = bitmap.Value[i+1]
		ret += uint16(util.BitCount((w<<1)&^w)) + uint16(w>>15&^n)
	}

	w = n
	ret += uint16(util.BitCount((w<<1)&^w)) + uint16(w>>15)

	return ret
}
