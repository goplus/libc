package libc

import unsafe "unsafe"

func Sinh(x float64) float64 {
	type _cgoa_18_sinh struct {
		f float64
	}
	var u _cgoa_18_sinh
	u.f = x
	var w uint32
	var t float64
	var h float64
	var absx float64
	h = float64(0.5)
	if *(*uint64)(unsafe.Pointer(&u))>>int32(63) != 0 {
		h = -h
	}
	*(*uint64)(unsafe.Pointer(&u)) &= 9223372036854775807
	absx = u.f
	w = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(32))
	if w < uint32(1082535490) {
		t = Expm1(absx)
		if w < uint32(1072693248) {
			if w < uint32(1045430272) {
				return x
			}
			return h * (float64(int32(2))*t - t*t/(t+float64(int32(1))))
		}
		return h * (t + t/(t+float64(int32(1))))
	}
	t = __expo2(absx, float64(int32(2))*h)
	return t
}
