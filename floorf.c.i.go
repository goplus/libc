package libc

import unsafe "unsafe"

func Floorf(x float32) float32 {
	type _cgoa_18_floorf struct {
		f float32
	}
	var u _cgoa_18_floorf
	u.f = x
	var e int32 = int32(*(*uint32)(unsafe.Pointer(&u))>>int32(23)&uint32(255)) - int32(127)
	var m uint32
	if e >= int32(23) {
		return x
	}
	if e >= int32(0) {
		m = uint32(int32(8388607) >> e)
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
		if *(*uint32)(unsafe.Pointer(&u))>>int32(31) != 0 {
			*(*uint32)(unsafe.Pointer(&u)) += m
		}
		*(*uint32)(unsafe.Pointer(&u)) &= ^m
	} else {
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
		if *(*uint32)(unsafe.Pointer(&u))>>int32(31) == uint32(0) {
			*(*uint32)(unsafe.Pointer(&u)) = uint32(0)
		} else if *(*uint32)(unsafe.Pointer(&u))<<int32(1) != 0 {
			u.f = float32(-1.0)
		}
	}
	return u.f
}
