package container

import (
	"unsafe"
)

//at most 2047 runs. 2 uint16 per run. first uint16 is start value. second is continuous count - 1
type RunContainer struct {
	cnt   uint16
	value [4095]uint16
}

const runCntLimit = 2047

func (run *RunContainer) exists(v uint16) bool {
	return alreadyExists(run.existsImp(v))
}

func (run *RunContainer) existsImp(v uint16) int32 {
	if run.cnt == 0 {
		return -1
	}

	//2 uint16 serves as one unit
	var b uint16 = 0
	e := run.cnt / 2

	for b < e {
		mid := b + (e-b)/2

		lower := run.value[mid*2]
		high := run.value[mid*2] + run.value[mid*2+1]
		if v >= lower && v <= high {
			return int32(mid)
		} else if v < lower {
			e = mid
		} else {
			b = mid + 1
		}
	}

	return -1
}

func (run *RunContainer) add(v uint16) bool {
	idx := run.existsImp(v)
	if alreadyExists(idx) {
		return false
	}

	//can be added to previous / next run
	target := uint16(idx)
	//previous
	if target > 0 {
		previous := target - 1
		lower := run.value[previous*2]
		high := run.value[previous*2+1] + lower
		if high+1 == v {
			run.value[previous*2+1]++
			return true
		}
	} else {
		lower := run.value[target*2]
		high := run.value[target*2+1] + lower
		if v+1 == lower {
			run.value[target*2]--
			run.value[target*2+1]++
			return true
		}

		if v == high+1 {
			run.value[target*2+1]++
			return true
		}
	}

	if run.cnt+1 > runCntLimit {
		bitmap := (*BitMapContainer)(unsafe.Pointer(run.convert(BitmapContainerType)))
		return bitmap.add(v)
	}

	for i := run.cnt * 2; i > 2*target+1; i-- {
		run.value[i] = run.value[i-2]
		run.value[i+1] = run.value[i-2]
	}

	run.value[target*2] = v
	run.value[target*2+1] = 1
	run.cnt++
	return true
}

func (run *RunContainer) remove(v uint16) bool {
	idx := run.existsImp(v)

	if !alreadyExists(idx) {
		return false
	}

	target := uint16(idx)

	return run.doRemove(target, v)
}

func (run *RunContainer) doRemove(idx uint16, v uint16) bool {
	lower := run.value[idx*2]
	high := run.value[idx*2+1] + lower

	edge := false
	size := high - lower + 1
	if lower == v || high == v {
		if lower == v {
			run.value[idx*2]++
		}
		edge = true
		size--
		if size > 0 {
			run.value[idx*2+1]--
		}

		//if run length == 0, remove this run
		if edge && size == 0 {
			for i := idx; i+1 < run.cnt; i++ {
				run.value[idx*2] = run.value[idx*2+2]
				run.value[idx*2+1] = run.value[idx*2+3]
			}
			run.cnt--
		}
		return true
	}

	//v is in the middle

	//make room for this new run
	for i := run.cnt; i-1 > idx; i-- {
		run.value[i*2] = run.value[(i-1)*2]
		run.value[i*2+1] = run.value[(i-1)*2+1]
	}

	//adjust value[idx]
	run.value[idx*2+1] = v - 1 - lower

	run.value[(idx+1)*2] = v + 1
	run.value[(idx+1)*2+1] = high - (v + 1)

	run.cnt++
	return true
}

func (run *RunContainer) convert(target CONTAINTER_TYPE) *Container {
	switch target {
	case BitmapContainerType:
	case SortedArrayContainerType:
	case RunContainerType:
	}
	return nil
}
