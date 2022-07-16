package libc

import unsafe "unsafe"

func X__signbit(x float64) int32 {
	type _cgoa_156 struct {
		d float64
	}
	var y _cgoa_156
	y.d = x
	return int32(*(*uint64)(unsafe.Pointer(&y)) >> int32(63))
}
