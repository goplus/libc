package libc

import unsafe "unsafe"

var s1pio2_cgo18_sincosf float64 = float64(int32(1)) * 1.5707963267948966
var s2pio2_cgo19_sincosf float64 = float64(int32(2)) * 1.5707963267948966
var s3pio2_cgo20_sincosf float64 = float64(int32(3)) * 1.5707963267948966
var s4pio2_cgo21_sincosf float64 = float64(int32(4)) * 1.5707963267948966

func Sincosf(x float32, sin *float32, cos *float32) {
	var y float64
	var s float32
	var c float32
	var ix uint32
	var n uint32
	var sign uint32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_22_sincosf{x}))
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
					fp_force_evalf(func() float32 {
						if ix < uint32(1048576) {
							return x / 1.329228e+36
						} else {
							return x + 1.329228e+36
						}
					}())
				} else if 4 == 8 {
					fp_force_eval(float64(func() float32 {
						if ix < uint32(1048576) {
							return x / 1.329228e+36
						} else {
							return x + 1.329228e+36
						}
					}()))
				} else {
					fp_force_evall(float64(func() float32 {
						if ix < uint32(1048576) {
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
			*sin = x
			*cos = float32(1)
			return
		}
		*sin = __sindf(float64(x))
		*cos = __cosdf(float64(x))
		return
	}
	if ix <= uint32(1081824209) {
		if ix <= uint32(1075235811) {
			if sign != 0 {
				*sin = -__cosdf(float64(x) + s1pio2_cgo18_sincosf)
				*cos = __sindf(float64(x) + s1pio2_cgo18_sincosf)
			} else {
				*sin = __cosdf(s1pio2_cgo18_sincosf - float64(x))
				*cos = __sindf(s1pio2_cgo18_sincosf - float64(x))
			}
			return
		}
		*sin = -__sindf(func() float64 {
			if sign != 0 {
				return float64(x) + s2pio2_cgo19_sincosf
			} else {
				return float64(x) - s2pio2_cgo19_sincosf
			}
		}())
		*cos = -__cosdf(func() float64 {
			if sign != 0 {
				return float64(x) + s2pio2_cgo19_sincosf
			} else {
				return float64(x) - s2pio2_cgo19_sincosf
			}
		}())
		return
	}
	if ix <= uint32(1088565717) {
		if ix <= uint32(1085271519) {
			if sign != 0 {
				*sin = __cosdf(float64(x) + s3pio2_cgo20_sincosf)
				*cos = -__sindf(float64(x) + s3pio2_cgo20_sincosf)
			} else {
				*sin = -__cosdf(float64(x) - s3pio2_cgo20_sincosf)
				*cos = __sindf(float64(x) - s3pio2_cgo20_sincosf)
			}
			return
		}
		*sin = __sindf(func() float64 {
			if sign != 0 {
				return float64(x) + s4pio2_cgo21_sincosf
			} else {
				return float64(x) - s4pio2_cgo21_sincosf
			}
		}())
		*cos = __cosdf(func() float64 {
			if sign != 0 {
				return float64(x) + s4pio2_cgo21_sincosf
			} else {
				return float64(x) - s4pio2_cgo21_sincosf
			}
		}())
		return
	}
	if ix >= uint32(2139095040) {
		*sin = func() (_cgo_ret float32) {
			_cgo_addr := &*cos
			*_cgo_addr = x - x
			return *_cgo_addr
		}()
		return
	}
	n = uint32(__rem_pio2f(x, &y))
	s = __sindf(y)
	c = __cosdf(y)
	switch n & uint32(3) {
	case uint32(0):
		*sin = s
		*cos = c
		break
	case uint32(1):
		*sin = c
		*cos = -s
		break
	case uint32(2):
		*sin = -s
		*cos = -c
		break
	case uint32(3):
		fallthrough
	default:
		*sin = -c
		*cos = s
		break
	}
}

type _cgoz_22_sincosf struct {
	_f float32
}
