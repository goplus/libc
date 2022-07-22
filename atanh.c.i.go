package libc

import unsafe "unsafe"

func Atanh(x float64) float64 {
	type _cgoa_18_atanh struct {
		f float64
	}
	var u _cgoa_18_atanh
	u.f = x
	var e uint32 = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(52) & uint64(2047))
	var s uint32 = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(63))
	var y float64
	*(*uint64)(unsafe.Pointer(&u)) &= 9223372036854775807
	y = u.f
	if e < uint32(1022) {
		if e < uint32(991) {
			if e == uint32(0) {
				for {
					if 4 == 4 {
						fp_force_evalf(float32(y))
					} else if 4 == 8 {
						fp_force_eval(float64(float32(y)))
					} else {
						fp_force_evall(float64(float32(y)))
					}
					if true {
						break
					}
				}
			}
		} else {
			y = 0.5 * Log1p(float64(int32(2))*y+float64(int32(2))*y*y/(float64(int32(1))-y))
		}
	} else {
		y = 0.5 * Log1p(float64(int32(2))*(y/(float64(int32(1))-y)))
	}
	return func() float64 {
		if s != 0 {
			return -y
		} else {
			return y
		}
	}()
}
