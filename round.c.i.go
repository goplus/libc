package libc

import unsafe "unsafe"

var _cgos_toint_round float64 = float64(int32(1)) / 2.2204460492503131e-16

func Round(x float64) float64 {
	type _cgoa_18_round struct {
		f float64
	}
	var u _cgoa_18_round
	u.f = x
	var e int32 = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(52) & uint64(2047))
	var y float64
	if e >= 1075 {
		return x
	}
	if *(*uint64)(unsafe.Pointer(&u))>>int32(63) != 0 {
		x = -x
	}
	if e < 1022 {
		for {
			if 8 == 4 {
				fp_force_evalf(float32(x + _cgos_toint_round))
			} else if 8 == 8 {
				fp_force_eval(x + _cgos_toint_round)
			} else {
				fp_force_evall(float64(x + _cgos_toint_round))
			}
			if true {
				break
			}
		}
		return float64(int32(0)) * u.f
	}
	y = x + _cgos_toint_round - _cgos_toint_round - x
	if y > 0.5 {
		y = y + x - float64(int32(1))
	} else if y <= -0.5 {
		y = y + x + float64(int32(1))
	} else {
		y = y + x
	}
	if *(*uint64)(unsafe.Pointer(&u))>>int32(63) != 0 {
		y = -y
	}
	return y
}
