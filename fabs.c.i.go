package libc

import unsafe "unsafe"

func fabs(x float64) float64 {
	type _cgoa_205 struct {
		f float64
	}
	var u _cgoa_205
	u.f = x
	*(*uint64)(unsafe.Pointer(&u)) &= 9223372036854775807
	return u.f
}
