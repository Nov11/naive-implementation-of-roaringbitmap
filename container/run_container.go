package container

import (
	"unsafe"
)

//at most 2047 runs. 2 uint16 per run. first uint16 is start value. second is continuous count - 1
type RunContainer struct {
	cnt   uint16
	value [4095]uint16
	//value []uint16
}

const runCntLimit = 2047

func (run *RunContainer) exists(v uint16) bool {
	return run.alreadyExists(run.existsImp(v), v)
}

func (run *RunContainer) alreadyExists(idx uint16, v uint16) bool {
	if idx >= run.cnt {
		return false
	}

	lower := run.value[idx*2]
	high := run.value[idx*2+1] + lower

	return lower <= v && v <= high
}

func (run *RunContainer) existsImp(v uint16) uint16 {
	if run.cnt == 0 {
		return run.cnt
	}

	//2 uint16 serves as one unit
	var b uint16 = 0
	e := run.cnt

	for b < e {
		mid := b + (e-b)/2

		lower := run.value[mid*2]
		high := run.value[mid*2] + run.value[mid*2+1]
		if v >= lower && v <= high {
			return mid
		} else if v < lower {
			e = mid
		} else {
			b = mid + 1
		}
	}

	return e
}

func (run *RunContainer) add(v uint16) bool {
	idx := run.existsImp(v)
	if run.alreadyExists(idx, v) {
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
	}

	if target < run.cnt {
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

	for i := run.cnt; i > target; i-- {
		run.value[2*i] = run.value[2*(i-1)]
		run.value[2*i+1] = run.value[2*(i-1)+1]
	}

	run.value[target*2] = v
	run.value[target*2+1] = 0
	run.cnt++
	return true
}

func (run *RunContainer) del(v uint16) bool {
	idx := run.existsImp(v)

	if !run.alreadyExists(idx, v) {
		return false
	}

	target := uint16(idx)

	return run.doRemove(target, v)
}

func (run *RunContainer) doRemove(idx uint16, v uint16) bool {
	lower := run.value[idx*2]
	high := run.value[idx*2+1] + lower

	size := high - lower + 1
	if lower == v || high == v {
		if lower == v {
			run.value[idx*2]++
		}

		size--
		if size > 0 {
			run.value[idx*2+1]--
		}

		//if run length == 0, remove this run
		if size == 0 {
			for i := idx; i+1 < run.cnt; i++ {
				run.value[idx*2] = run.value[(idx+1)*2]
				run.value[idx*2+1] = run.value[(idx+1)*2+1]
			}
			run.cnt--
			run.value[run.cnt*2] = 0
			run.value[run.cnt*2+1] = 0
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

func (run *RunContainer) convert(target TypeContainer) *Container {
	switch target {
	case BitmapContainerType:
	case SortedArrayContainerType:
	case RunContainerType:
	}
	return nil
}
