package libc

import unsafe "unsafe"

func Fabsf(x float32) float32 {
	type _cgoa_251 struct {
		f float32
	}
	var u _cgoa_251
	u.f = x
	*(*uint32)(unsafe.Pointer(&u)) &= uint32(2147483647)
	return u.f
}
