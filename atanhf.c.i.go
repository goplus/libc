package libc

import unsafe "unsafe"

func Atanhf(x float32) float32 {
	type _cgoa_18_atanhf struct {
		f float32
	}
	var u _cgoa_18_atanhf
	u.f = x
	var s uint32 = *(*uint32)(unsafe.Pointer(&u)) >> int32(31)
	var y float32
	*(*uint32)(unsafe.Pointer(&u)) &= uint32(2147483647)
	y = u.f
	if *(*uint32)(unsafe.Pointer(&u)) < uint32(1056964608) {
		if *(*uint32)(unsafe.Pointer(&u)) < uint32(796917760) {
			if *(*uint32)(unsafe.Pointer(&u)) < uint32(8388608) {
				for {
					if 4 == 4 {
						fp_force_evalf(float32(y * y))
					} else if 4 == 8 {
						fp_force_eval(float64(float32(y * y)))
					} else {
						fp_force_evall(float64(float32(y * y)))
					}
					if true {
						break
					}
				}
			}
		} else {
			y = 0.5 * Log1pf(float32(int32(2))*y+float32(int32(2))*y*y/(float32(int32(1))-y))
		}
	} else {
		y = 0.5 * Log1pf(float32(int32(2))*(y/(float32(int32(1))-y)))
	}
	return func() float32 {
		if s != 0 {
			return -y
		} else {
			return y
		}
	}()
}
