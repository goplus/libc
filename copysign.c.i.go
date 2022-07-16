package libc

import unsafe "unsafe"

func Copysign(x float64, y float64) float64 {
	type _cgoa_192 struct {
		f float64
	}
	var ux _cgoa_192
	ux.f = x
	var uy _cgoa_192
	uy.f = y
	*(*uint64)(unsafe.Pointer(&ux)) &= 9223372036854775807
	*(*uint64)(unsafe.Pointer(&ux)) |= *(*uint64)(unsafe.Pointer(&uy)) & 9223372036854775808
	return ux.f
}
