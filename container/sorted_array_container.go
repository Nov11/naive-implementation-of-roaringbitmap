package container

import (
	"unsafe"
)

//8KB at most. elements are uint16 integers
type SortedArrayContainer struct {
	cnt   uint16
	value [4095]uint16
	//Value []uint16
}

const sortedArrayCntLimit = 4096 - 1

func (array *SortedArrayContainer) Exists(v uint16) bool {
	ret := array.existsImp(v)
	return array.alreadyExists(ret, v)
}

func (array *SortedArrayContainer) alreadyExists(idx uint16, v uint16) bool {
	return idx != array.cnt && array.value[idx] == v
}

func (array *SortedArrayContainer) existsImp(v uint16) uint16 {
	if array.cnt == 0 {
		return array.cnt
	}

	var b uint16 = 0
	e := array.cnt

	for b < e {
		mid := b + (e-b)/2
		if array.value[mid] == v {
			return mid
		} else if array.value[mid] < v {
			b = mid + 1
		} else {
			e = mid
		}
	}

	return e
}

func (array *SortedArrayContainer) Add(v uint16) bool {
	idx := array.existsImp(v)
	if array.alreadyExists(idx, v) {
		return false
	}

	if array.cnt+1 > sortedArrayCntLimit {
		bitmap := (*BitMapContainer)(unsafe.Pointer(array.convert(BitmapContainerType)))
		return bitmap.Add(v)
	}

	target := uint16(idx)
	for last := array.cnt; last > target; last-- {
		array.value[last] = array.value[last-1]
	}
	array.value[target] = v

	array.cnt++
	return true
}

func (array *SortedArrayContainer) Del(v uint16) bool {
	if array.cnt == 0 {
		return false
	}

	idx := array.existsImp(v)
	if !array.alreadyExists(idx, v) {
		return false
	}

	target := uint16(idx)

	for i := target; i+1 < array.cnt; i++ {
		array.value[i] = array.value[i+1]
	}
	array.value[array.cnt-1] = 0
	array.cnt--
	return true
}

func (array *SortedArrayContainer) convert(target TypeContainer) *Container {
	switch target {
	case BitmapContainerType:
	case SortedArrayContainerType:
	case RunContainerType:
	}
	return nil
}
