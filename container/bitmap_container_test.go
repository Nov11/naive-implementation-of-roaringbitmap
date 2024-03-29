package container

import (
	"github.com/Nov11/naive-implementation-of-roaringbitmap/util"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

//Add exist get
func TestAdd_BitMap(t *testing.T) {
	bitmap := MakeBitMapContainer()

	var value uint64 = 1
	for value <= math.MaxUint32 {
		v := uint16(value)
		assert.Equal(t, true, bitmap.Add(v))
		assert.Equal(t, true, bitmap.Exists(v))
		assert.Equal(t, true, bitmap.Del(v))
		value <<= 1
	}
}

func TestAddBit_BitMap(t *testing.T) {
	bitmap := MakeBitMapContainer()

	bitmap.Add(18)

	for i := 0; i < 4096; i++ {
		binary := bitmap.Value[i]
		if i != 1 {
			assert.Equal(t, 0, util.BitCount(binary))
		} else {
			assert.Equal(t, uint16(0x04), binary)
		}
	}
}

func TestDelBit_BitMap(t *testing.T) {
	bitmap := MakeBitMapContainer()

	bitmap.Add(18)
	bitmap.Del(18)

	for i := 0; i < 4096; i++ {
		binary := bitmap.Value[i]
		assert.Equal(t, 0, util.BitCount(binary))
	}

	bitmap.Del(777)
	for i := 0; i < 4096; i++ {
		binary := bitmap.Value[i]
		assert.Equal(t, 0, util.BitCount(binary))
	}
}

func TestFromBinary_BitMap(t *testing.T) {
	tmp := [8192]byte{}

	for i := 0; i < 8192; i++ {
		tmp[i] = byte(i % 256)
	}

	bitmap, err := FromBinaryArray(tmp[:])
	assert.NoError(t, err, "should not raise error")

	for i := 0; i < 4096; i++ {
		expect0 := tmp[i*2]
		expect1 := tmp[i*2+1]
		assert.Equal(t, expect1, byte(bitmap.Value[i]>>8))
		assert.Equal(t, expect0, byte(bitmap.Value[i]))
	}
}

func TestFromBinary2_BitMap(t *testing.T) {
	tmp := make([]byte, 8192)

	for i := 0; i < 8192; i++ {
		tmp[i] = byte(i % 256)
	}

	bitmap, err := FromBinaryArray(tmp[:])
	assert.NoError(t, err, "should not raise error")

	for i := 0; i < 4096; i++ {
		expect0 := tmp[i*2]
		expect1 := tmp[i*2+1]
		assert.Equal(t, expect1, byte(bitmap.Value[i]>>8))
		assert.Equal(t, expect0, byte(bitmap.Value[i]))
	}
}

func TestToBinary_BitMap(t *testing.T) {
	bitmap := MakeBitMapContainer()

	ret := bitmap.ToBinaryArray()
	assert.Equal(t, 8192, len(ret))
}

func TestCountRun_Bitmap(t *testing.T) {
	bitmap := MakeBitMapContainer()

	bitmap.Value[0] = 0x8000
	bitmap.Value[1] = 0x0001

	assert.Equal(t, uint16(1), bitmap.countRuns())
}
