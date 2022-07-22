package libc

import unsafe "unsafe"

func Asinhf(x float32) float32 {
	type _cgoa_18_asinhf struct {
		f float32
	}
	var u _cgoa_18_asinhf
	u.f = x
	var i uint32 = *(*uint32)(unsafe.Pointer(&u)) & uint32(2147483647)
	var s uint32 = *(*uint32)(unsafe.Pointer(&u)) >> int32(31)
	*(*uint32)(unsafe.Pointer(&u)) = i
	x = u.f
	if i >= uint32(1166016512) {
		x = Logf(x) + 0.693147182
	} else if i >= uint32(1073741824) {
		x = Logf(float32(int32(2))*x + float32(int32(1))/(Sqrtf(x*x+float32(int32(1)))+x))
	} else if i >= uint32(964689920) {
		x = Log1pf(x + x*x/(Sqrtf(x*x+float32(int32(1)))+float32(int32(1))))
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
	}
	return func() float32 {
		if s != 0 {
			return -x
		} else {
			return x
		}
	}()
}
