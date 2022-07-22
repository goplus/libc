package libc

import unsafe "unsafe"

func Cosh(x float64) float64 {
	type _cgoa_18_cosh struct {
		f float64
	}
	var u _cgoa_18_cosh
	u.f = x
	var w uint32
	var t float64
	*(*uint64)(unsafe.Pointer(&u)) &= 9223372036854775807
	x = u.f
	w = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(32))
	if w < uint32(1072049730) {
		if w < uint32(1045430272) {
			for {
				if 8 == 4 {
					fp_force_evalf(float32(x + float64(1.329228e+36)))
				} else if 8 == 8 {
					fp_force_eval(x + float64(1.329228e+36))
				} else {
					fp_force_evall(float64(x + float64(1.329228e+36)))
				}
				if true {
					break
				}
			}
			return float64(int32(1))
		}
		t = Expm1(x)
		return float64(int32(1)) + t*t/(float64(int32(2))*(float64(int32(1))+t))
	}
	if w < uint32(1082535490) {
		t = Exp(x)
		return 0.5 * (t + float64(int32(1))/t)
	}
	t = __expo2(x, 1)
	return t
}
