package libc

import unsafe "unsafe"

func Fabs(x float64) float64 {
	type _cgoa_18_fabs struct {
		f float64
	}
	var u _cgoa_18_fabs
	u.f = x
	*(*uint64)(unsafe.Pointer(&u)) &= 9223372036854775807
	return u.f
}
