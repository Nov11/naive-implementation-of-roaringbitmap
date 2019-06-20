package container

import (
	"unsafe"
)

//8KB at most. elements are uint16 integers
type SortedArrayContainer struct {
	cnt   uint16
	value [4095]uint16
}

const sortedArrayCntLimit = 4096 - 1

func (array *SortedArrayContainer) exists(v uint16) bool {
	ret := array.existsImp(v)
	return alreadyExists(ret)
}

func alreadyExists(idx int32) bool {
	return idx != -1
}

func (array *SortedArrayContainer) existsImp(v uint16) int32 {
	if array.cnt == 0 {
		return -1
	}

	var b uint16 = 0
	e := array.cnt

	for b < e {
		mid := b + (e-b)/2
		if array.value[mid] == v {
			return int32(mid)
		} else if array.value[mid] < v {
			b = mid + 1
		} else {
			e = mid
		}
	}

	return -1
}

func (array *SortedArrayContainer) add(v uint16) bool {
	idx := array.existsImp(v)
	if alreadyExists(idx) {
		return false
	}

	if array.cnt+1 > sortedArrayCntLimit {
		bitmap := (*BitMapContainer)(unsafe.Pointer(array.convert(BitmapContainerType)))
		return bitmap.add(v)
	}

	target := uint16(idx)
	for last := array.cnt; last > target; last-- {
		array.value[last] = array.value[last-1]
	}
	array.value[target] = v

	array.cnt++
	return true
}

func (array *SortedArrayContainer) del(v uint16) bool {
	if array.cnt == 0 {
		return false
	}

	idx := array.existsImp(v)
	if !alreadyExists(idx) {
		return false
	}

	target := uint16(idx)

	for i := target; i+1 < array.cnt; i++ {
		array.value[i] = array.value[i+1]
	}
	array.cnt--
	return true
}

func (array *SortedArrayContainer) convert(target CONTAINTER_TYPE) *Container {
	switch target {
	case BitmapContainerType:
	case SortedArrayContainerType:
	case RunContainerType:
	}
	return nil
}
