package libc

import unsafe "unsafe"

var s1pio2_cgos__sinf float64 = float64(int32(1)) * 1.5707963267948966
var s2pio2_cgos__sinf float64 = float64(int32(2)) * 1.5707963267948966
var s3pio2_cgos__sinf float64 = float64(int32(3)) * 1.5707963267948966
var s4pio2_cgos__sinf float64 = float64(int32(4)) * 1.5707963267948966

func Sinf(x float32) float32 {
	var y float64
	var ix uint32
	var n int32
	var sign int32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_18_sinf{x}))
		if true {
			break
		}
	}
	sign = int32(ix >> int32(31))
	ix &= uint32(2147483647)
	if ix <= uint32(1061752794) {
		if ix < uint32(964689920) {
			for {
				if 4 == 4 {
					fp_force_evalf(func() float32 {
						if ix < uint32(8388608) {
							return x / 1.329228e+36
						} else {
							return x + 1.329228e+36
						}
					}())
				} else if 4 == 8 {
					fp_force_eval(float64(func() float32 {
						if ix < uint32(8388608) {
							return x / 1.329228e+36
						} else {
							return x + 1.329228e+36
						}
					}()))
				} else {
					fp_force_evall(float64(func() float32 {
						if ix < uint32(8388608) {
							return x / 1.329228e+36
						} else {
							return x + 1.329228e+36
						}
					}()))
				}
				if true {
					break
				}
			}
			return x
		}
		return __sindf(float64(x))
	}
	if ix <= uint32(1081824209) {
		if ix <= uint32(1075235811) {
			if sign != 0 {
				return -__cosdf(float64(x) + s1pio2_cgos__sinf)
			} else {
				return __cosdf(float64(x) - s1pio2_cgos__sinf)
			}
		}
		return __sindf(func() float64 {
			if sign != 0 {
				return -(float64(x) + s2pio2_cgos__sinf)
			} else {
				return -(float64(x) - s2pio2_cgos__sinf)
			}
		}())
	}
	if ix <= uint32(1088565717) {
		if ix <= uint32(1085271519) {
			if sign != 0 {
				return __cosdf(float64(x) + s3pio2_cgos__sinf)
			} else {
				return -__cosdf(float64(x) - s3pio2_cgos__sinf)
			}
		}
		return __sindf(func() float64 {
			if sign != 0 {
				return float64(x) + s4pio2_cgos__sinf
			} else {
				return float64(x) - s4pio2_cgos__sinf
			}
		}())
	}
	if ix >= uint32(2139095040) {
		return x - x
	}
	n = __rem_pio2f(x, &y)
	switch n & int32(3) {
	case int32(0):
		return __sindf(y)
	case int32(1):
		return __cosdf(y)
	case int32(2):
		return __sindf(-y)
	default:
		return -__cosdf(y)
	}
	return 0
}

type _cgoz_18_sinf struct {
	_f float32
}
