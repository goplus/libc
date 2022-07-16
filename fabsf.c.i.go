package libc

import unsafe "unsafe"

func Fabsf(x float32) float32 {
	type _cgoa_244 struct {
		f float32
	}
	var u _cgoa_244
	u.f = x
	*(*uint32)(unsafe.Pointer(&u)) &= uint32(2147483647)
	return u.f
}
