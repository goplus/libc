package libc

import unsafe "unsafe"

func Modf(x float64, iptr *float64) float64 {
	type _cgoa_18_modf struct {
		f float64
	}
	var u _cgoa_18_modf
	u.f = x
	var mask uint64
	var e int32 = int32(*(*uint64)(unsafe.Pointer(&u))>>int32(52)&uint64(2047)) - int32(1023)
	if e >= int32(52) {
		*iptr = x
		if e == int32(1024) && *(*uint64)(unsafe.Pointer(&u))<<int32(12) != uint64(0) {
			return x
		}
		*(*uint64)(unsafe.Pointer(&u)) &= 9223372036854775808
		return u.f
	}
	if e < int32(0) {
		*(*uint64)(unsafe.Pointer(&u)) &= 9223372036854775808
		*iptr = u.f
		return x
	}
	mask = 4503599627370495 >> e
	if *(*uint64)(unsafe.Pointer(&u))&mask == uint64(0) {
		*iptr = x
		*(*uint64)(unsafe.Pointer(&u)) &= 9223372036854775808
		return u.f
	}
	*(*uint64)(unsafe.Pointer(&u)) &= ^mask
	*iptr = u.f
	return x - u.f
}
