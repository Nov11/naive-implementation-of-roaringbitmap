package container

import (
	"github.com/Nov11/naive-implementation-of-roaringbitmap/util"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func makeSortedArrayContainer() *SortedArrayContainer {
	//s := [4096]uint16{}
	//pointer := unsafe.Pointer(&s[0])
	//ret := (*SortedArrayContainer)(pointer)
	return new(SortedArrayContainer)
}

//add exist get
func TestAdd_SortedArray(t *testing.T) {
	sortedArray := makeSortedArrayContainer()

	var value uint64 = 1
	for value <= math.MaxUint32 {
		v := uint16(value)
		assert.Equal(t, true, sortedArray.add(v))
		assert.Equal(t, uint16(1), sortedArray.cnt)
		assert.Equal(t, true, sortedArray.exists(v))
		assert.Equal(t, true, sortedArray.del(v))
		assert.Equal(t, uint16(0), sortedArray.cnt)
		value <<= 1
	}
}

func TestAddBit_SortedArray(t *testing.T) {
	sortedArray := makeSortedArrayContainer()

	sortedArray.add(18)
	assert.Equal(t, uint16(0x01), sortedArray.cnt)
	for i := 0; i < len(sortedArray.value); i++ {
		binary := sortedArray.value[i]
		if i == 0 {
			assert.Equal(t, uint16(18), binary)
		} else {
			assert.Equal(t, 0, util.BitCount(binary))
		}
	}
}

func TestDelBit_SortedArray(t *testing.T) {
	sortedArray := makeSortedArrayContainer()

	sortedArray.add(18)
	sortedArray.del(18)
	assert.Equal(t, uint16(0), sortedArray.cnt)

	for i := 0; i < len(sortedArray.value); i++ {
		binary := sortedArray.value[i]
		assert.Equal(t, 0, util.BitCount(binary))
	}

	sortedArray.del(777)
	for i := 0; i < len(sortedArray.value); i++ {
		binary := sortedArray.value[i]
		assert.Equal(t, 0, util.BitCount(binary))
	}
}

//todo: add elements to 4096 and test container conversion
