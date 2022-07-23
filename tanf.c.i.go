package libc

import unsafe "unsafe"

var t1pio2_cgos__tanf float64 = float64(int32(1)) * 1.5707963267948966
var t2pio2_cgos__tanf float64 = float64(int32(2)) * 1.5707963267948966
var t3pio2_cgos__tanf float64 = float64(int32(3)) * 1.5707963267948966
var t4pio2_cgos__tanf float64 = float64(int32(4)) * 1.5707963267948966

func Tanf(x float32) float32 {
	var y float64
	var ix uint32
	var n uint32
	var sign uint32
	for {
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_18_tanf{x}))
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
		return __tandf(float64(x), int32(0))
	}
	if ix <= uint32(1081824209) {
		if ix <= uint32(1075235811) {
			return __tandf(func() float64 {
				if sign != 0 {
					return float64(x) + t1pio2_cgos__tanf
				} else {
					return float64(x) - t1pio2_cgos__tanf
				}
			}(), int32(1))
		} else {
			return __tandf(func() float64 {
				if sign != 0 {
					return float64(x) + t2pio2_cgos__tanf
				} else {
					return float64(x) - t2pio2_cgos__tanf
				}
			}(), int32(0))
		}
	}
	if ix <= uint32(1088565717) {
		if ix <= uint32(1085271519) {
			return __tandf(func() float64 {
				if sign != 0 {
					return float64(x) + t3pio2_cgos__tanf
				} else {
					return float64(x) - t3pio2_cgos__tanf
				}
			}(), int32(1))
		} else {
			return __tandf(func() float64 {
				if sign != 0 {
					return float64(x) + t4pio2_cgos__tanf
				} else {
					return float64(x) - t4pio2_cgos__tanf
				}
			}(), int32(0))
		}
	}
	if ix >= uint32(2139095040) {
		return x - x
	}
	n = uint32(__rem_pio2f(x, &y))
	return __tandf(y, int32(n&uint32(1)))
}

type _cgoz_18_tanf struct {
	_f float32
}
