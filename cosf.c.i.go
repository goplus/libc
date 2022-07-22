package libc

import unsafe "unsafe"

var c1pio2_cgo18_cosf float64 = float64(int32(1)) * 1.5707963267948966
var c2pio2_cgo19_cosf float64 = float64(int32(2)) * 1.5707963267948966
var c3pio2_cgo20_cosf float64 = float64(int32(3)) * 1.5707963267948966
var c4pio2_cgo21_cosf float64 = float64(int32(4)) * 1.5707963267948966

func Cosf(x float32) float32 {
	var y float64
	var ix uint32
	var n uint32
	var sign uint32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_22_cosf{x}))
		if true {
			break
		}
	}
	sign = ix >> int32(31)
	ix &= uint32(2147483647)
	if ix <= uint32(1061752794) {
		if ix < uint32(964689920) {
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
			return float32(1)
		}
		return __cosdf(float64(x))
	}
	if ix <= uint32(1081824209) {
		if ix > uint32(1075235811) {
			return -__cosdf(func() float64 {
				if sign != 0 {
					return float64(x) + c2pio2_cgo19_cosf
				} else {
					return float64(x) - c2pio2_cgo19_cosf
				}
			}())
		} else if sign != 0 {
			return __sindf(float64(x) + c1pio2_cgo18_cosf)
		} else {
			return __sindf(c1pio2_cgo18_cosf - float64(x))
		}
	}
	if ix <= uint32(1088565717) {
		if ix > uint32(1085271519) {
			return __cosdf(func() float64 {
				if sign != 0 {
					return float64(x) + c4pio2_cgo21_cosf
				} else {
					return float64(x) - c4pio2_cgo21_cosf
				}
			}())
		} else if sign != 0 {
			return __sindf(float64(-x) - c3pio2_cgo20_cosf)
		} else {
			return __sindf(float64(x) - c3pio2_cgo20_cosf)
		}
	}
	if ix >= uint32(2139095040) {
		return x - x
	}
	n = uint32(__rem_pio2f(x, &y))
	switch n & uint32(3) {
	case uint32(0):
		return __cosdf(y)
	case uint32(1):
		return __sindf(-y)
	case uint32(2):
		return -__cosdf(y)
	default:
		return __sindf(y)
	}
	return 0
}

type _cgoz_22_cosf struct {
	_f float32
}
