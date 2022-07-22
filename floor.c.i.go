package libc

import unsafe "unsafe"

var toint_cgo18_floor float64 = float64(int32(1)) / 2.2204460492503131e-16

func Floor(x float64) float64 {
	type _cgoa_19_floor struct {
		f float64
	}
	var u _cgoa_19_floor
	u.f = x
	var e int32 = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(52) & uint64(2047))
	var y float64
	if e >= 1075 || x == float64(int32(0)) {
		return x
	}
	if *(*uint64)(unsafe.Pointer(&u))>>int32(63) != 0 {
		y = x - toint_cgo18_floor + toint_cgo18_floor - x
	} else {
		y = x + toint_cgo18_floor - toint_cgo18_floor - x
	}
	if e <= 1022 {
		for {
			if 8 == 4 {
				fp_force_evalf(float32(y))
			} else if 8 == 8 {
				fp_force_eval(y)
			} else {
				fp_force_evall(float64(y))
			}
			if true {
				break
			}
		}
		return float64(func() int32 {
			if *(*uint64)(unsafe.Pointer(&u))>>int32(63) != 0 {
				return -1
			} else {
				return int32(0)
			}
		}())
	}
	if y > float64(int32(0)) {
		return x + y - float64(int32(1))
	}
	return x + y
}
