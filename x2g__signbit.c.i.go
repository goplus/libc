package libc

import unsafe "unsafe"

func X__signbit(x float64) int32 {
	type _cgoa_18___signbit struct {
		d float64
	}
	var y _cgoa_18___signbit
	y.d = x
	return int32(*(*uint64)(unsafe.Pointer(&y)) >> int32(63))
}
