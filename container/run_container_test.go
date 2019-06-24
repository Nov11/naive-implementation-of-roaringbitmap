package container

import (
	"github.com/Nov11/naive-implementation-of-roaringbitmap/util"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func makeRunContainer() *RunContainer {
	//s := make([]uint16, 4096)
	//ret := (*RunContainer)(unsafe.Pointer(&s[0]))
	//return ret
	return new(RunContainer)
}

//Add exist get
func TestAdd_RunContainer(t *testing.T) {
	run := makeRunContainer()

	var value uint64 = 1
	for value <= math.MaxUint32 {
		v := uint16(value)
		assert.Equal(t, true, run.Add(v))
		assert.Equal(t, uint16(1), run.cnt)
		assert.Equal(t, true, run.Exists(v))
		assert.Equal(t, true, run.Del(v))
		assert.Equal(t, uint16(0), run.cnt)
		value <<= 1
	}
}

func TestAddBit_RunContainer(t *testing.T) {
	run := makeRunContainer()

	run.Add(18)
	assert.Equal(t, uint16(0x01), run.cnt)
	for i := 0; i < len(run.value); i++ {
		binary := run.value[i]
		if i == 0 {
			assert.Equal(t, uint16(18), binary)
		} else {
			assert.Equal(t, 0, util.BitCount(binary))
		}
	}
}

func TestDelBit_RunContainer(t *testing.T) {
	run := makeRunContainer()

	run.Add(18)
	run.Del(18)
	assert.Equal(t, uint16(0), run.cnt)

	for i := 0; i < len(run.value); i++ {
		binary := run.value[i]
		assert.Equal(t, 0, util.BitCount(binary))
	}

	run.Del(777)
	for i := 0; i < len(run.value); i++ {
		binary := run.value[i]
		assert.Equal(t, 0, util.BitCount(binary))
	}
}
