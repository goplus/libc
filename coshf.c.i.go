package libc

import unsafe "unsafe"

func Coshf(x float32) float32 {
	type _cgoa_18_coshf struct {
		f float32
	}
	var u _cgoa_18_coshf
	u.f = x
	var w uint32
	var t float32
	*(*uint32)(unsafe.Pointer(&u)) &= uint32(2147483647)
	x = u.f
	w = *(*uint32)(unsafe.Pointer(&u))
	if w < uint32(1060205079) {
		if w < uint32(964689920) {
			for {
				if 4 == 4 {
					fp_force_evalf(x + 1.329228e+36)
				} else if 4 == 8 {
					fp_force_eval(float64(x + 1.329228e+36))
				} else {
					fp_force_evall(float64(x + 1.329228e+36))
				}
				if true {
					break
				}
			}
			return float32(int32(1))
		}
		t = Expm1f(x)
		return float32(int32(1)) + t*t/(float32(int32(2))*(float32(int32(1))+t))
	}
	if w < uint32(1118925335) {
		t = Expf(x)
		return 0.5 * (t + float32(int32(1))/t)
	}
	t = __expo2f(x, 1.0)
	return t
}
