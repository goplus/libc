package libc

import unsafe "unsafe"

func Trunc(x float64) float64 {
	type _cgoa_18_trunc struct {
		f float64
	}
	var u _cgoa_18_trunc
	u.f = x
	var e int32 = int32(*(*uint64)(unsafe.Pointer(&u))>>int32(52)&uint64(2047)) - int32(1023) + int32(12)
	var m uint64
	if e >= 64 {
		return x
	}
	if e < int32(12) {
		e = int32(1)
	}
	m = 18446744073709551615 >> e
	if *(*uint64)(unsafe.Pointer(&u))&m == uint64(0) {
		return x
	}
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
	*(*uint64)(unsafe.Pointer(&u)) &= ^m
	return u.f
}
