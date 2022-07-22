package libc

import unsafe "unsafe"

func Truncf(x float32) float32 {
	type _cgoa_18_truncf struct {
		f float32
	}
	var u _cgoa_18_truncf
	u.f = x
	var e int32 = int32(*(*uint32)(unsafe.Pointer(&u))>>int32(23)&uint32(255)) - int32(127) + int32(9)
	var m uint32
	if e >= 32 {
		return x
	}
	if e < int32(9) {
		e = int32(1)
	}
	m = 4294967295 >> e
	if *(*uint32)(unsafe.Pointer(&u))&m == uint32(0) {
		return x
	}
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
	*(*uint32)(unsafe.Pointer(&u)) &= ^m
	return u.f
}
