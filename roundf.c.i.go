package libc

import unsafe "unsafe"

var _cgos_toint_roundf float32 = float32(int32(1)) / 1.1920929e-7

func Roundf(x float32) float32 {
	type _cgoa_18_roundf struct {
		f float32
	}
	var u _cgoa_18_roundf
	u.f = x
	var e int32 = int32(*(*uint32)(unsafe.Pointer(&u)) >> int32(23) & uint32(255))
	var y float32
	if e >= 150 {
		return x
	}
	if *(*uint32)(unsafe.Pointer(&u))>>int32(31) != 0 {
		x = -x
	}
	if e < 126 {
		for {
			if 4 == 4 {
				fp_force_evalf(x + _cgos_toint_roundf)
			} else if 4 == 8 {
				fp_force_eval(float64(x + _cgos_toint_roundf))
			} else {
				fp_force_evall(float64(x + _cgos_toint_roundf))
			}
			if true {
				break
			}
		}
		return float32(int32(0)) * u.f
	}
	y = x + _cgos_toint_roundf - _cgos_toint_roundf - x
	if y > 0.5 {
		y = y + x - float32(int32(1))
	} else if y <= -0.5 {
		y = y + x + float32(int32(1))
	} else {
		y = y + x
	}
	if *(*uint32)(unsafe.Pointer(&u))>>int32(31) != 0 {
		y = -y
	}
	return y
}
