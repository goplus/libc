package libc

import unsafe "unsafe"

func Sinhf(x float32) float32 {
	type _cgoa_18_sinhf struct {
		f float32
	}
	var u _cgoa_18_sinhf
	u.f = x
	var w uint32
	var t float32
	var h float32
	var absx float32
	h = float32(0.5)
	if *(*uint32)(unsafe.Pointer(&u))>>int32(31) != 0 {
		h = -h
	}
	*(*uint32)(unsafe.Pointer(&u)) &= uint32(2147483647)
	absx = u.f
	w = *(*uint32)(unsafe.Pointer(&u))
	if w < uint32(1118925335) {
		t = Expm1f(absx)
		if w < uint32(1065353216) {
			if w < uint32(964689920) {
				return x
			}
			return h * (float32(int32(2))*t - t*t/(t+float32(int32(1))))
		}
		return h * (t + t/(t+float32(int32(1))))
	}
	t = __expo2f(absx, float32(int32(2))*h)
	return t
}
