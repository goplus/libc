package libc

import unsafe "unsafe"

func __signbitf(x float32) int32 {
	type _cgoa_18___signbitf struct {
		f float32
	}
	var y _cgoa_18___signbitf
	y.f = x
	return int32(*(*uint32)(unsafe.Pointer(&y)) >> int32(31))
}
