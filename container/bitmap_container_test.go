package container

import (
	"github.com/Nov11/naive-implementation-of-roaringbitmap/util"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func makeBitMapContainer() *BitMapContainer {
	ret := new(BitMapContainer)
	//ret.value = make([]uint16, 4096)
	return ret
}

//add exist get
func TestAdd_BitMap(t *testing.T) {
	bitmap := makeBitMapContainer()

	var value uint64 = 1
	for value <= math.MaxUint32 {
		v := uint16(value)
		assert.Equal(t, true, bitmap.add(v))
		assert.Equal(t, true, bitmap.exists(v))
		assert.Equal(t, true, bitmap.del(v))
		value <<= 1
	}
}

func TestAddBit_BitMap(t *testing.T) {
	bitmap := makeBitMapContainer()

	bitmap.add(18)

	for i := 0; i < 4096; i++ {
		binary := bitmap.value[i]
		if i != 1 {
			assert.Equal(t, 0, util.BitCount(binary))
		} else {
			assert.Equal(t, uint16(0x04), binary)
		}
	}
}

func TestDelBit_BitMap(t *testing.T) {
	bitmap := makeBitMapContainer()

	bitmap.add(18)
	bitmap.del(18)

	for i := 0; i < 4096; i++ {
		binary := bitmap.value[i]
		assert.Equal(t, 0, util.BitCount(binary))
	}

	bitmap.del(777)
	for i := 0; i < 4096; i++ {
		binary := bitmap.value[i]
		assert.Equal(t, 0, util.BitCount(binary))
	}
}
