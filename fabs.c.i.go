package libc

import unsafe "unsafe"

func Fabs(x float64) float64 {
	type _cgoa_229 struct {
		f float64
	}
	var u _cgoa_229
	u.f = x
	*(*uint64)(unsafe.Pointer(&u)) &= 9223372036854775807
	return u.f
}
