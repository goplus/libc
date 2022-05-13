package libc

import unsafe "unsafe"

func __signbit(x float64) int32 {
	type _cgoa_97 struct {
		d float64
	}
	var y _cgoa_97
	y.d = x
	return int32(*(*uint64)(unsafe.Pointer(&y)) >> 63)
}
