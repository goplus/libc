package libc

import unsafe "unsafe"

func Asinh(x float64) float64 {
	type _cgoa_18_asinh struct {
		f float64
	}
	var u _cgoa_18_asinh
	u.f = x
	var e uint32 = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(52) & uint64(2047))
	var s uint32 = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(63))
	*(*uint64)(unsafe.Pointer(&u)) &= 9223372036854775807
	x = u.f
	if e >= uint32(1049) {
		x = Log(x) + 0.69314718055994529
	} else if e >= uint32(1024) {
		x = Log(float64(int32(2))*x + float64(int32(1))/(Sqrt(x*x+float64(int32(1)))+x))
	} else if e >= uint32(997) {
		x = Log1p(x + x*x/(Sqrt(x*x+float64(int32(1)))+float64(int32(1))))
	} else {
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
	}
	return func() float64 {
		if s != 0 {
			return -x
		} else {
			return x
		}
	}()
}
