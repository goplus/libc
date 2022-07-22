package libc

import unsafe "unsafe"

func Modff(x float32, iptr *float32) float32 {
	type _cgoa_18_modff struct {
		f float32
	}
	var u _cgoa_18_modff
	u.f = x
	var mask uint32
	var e int32 = int32(*(*uint32)(unsafe.Pointer(&u))>>int32(23)&uint32(255)) - int32(127)
	if e >= int32(23) {
		*iptr = x
		if e == int32(128) && *(*uint32)(unsafe.Pointer(&u))<<int32(9) != uint32(0) {
			return x
		}
		*(*uint32)(unsafe.Pointer(&u)) &= uint32(2147483648)
		return u.f
	}
	if e < int32(0) {
		*(*uint32)(unsafe.Pointer(&u)) &= uint32(2147483648)
		*iptr = u.f
		return x
	}
	mask = uint32(int32(8388607) >> e)
	if *(*uint32)(unsafe.Pointer(&u))&mask == uint32(0) {
		*iptr = x
		*(*uint32)(unsafe.Pointer(&u)) &= uint32(2147483648)
		return u.f
	}
	*(*uint32)(unsafe.Pointer(&u)) &= ^mask
	*iptr = u.f
	return x - u.f
}
