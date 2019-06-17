package util

import "math/bits"

func BitCount(v uint16) int {
	return bits.OnesCount16(v)
}
