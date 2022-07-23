package libc

import unsafe "unsafe"

var o_threshold_cgos__expm1 float64 = 709.78271289338397
var ln2_hi_cgos__expm1 float64 = 0.69314718036912382
var ln2_lo_cgos__expm1 float64 = 1.9082149292705877e-10
var invln2_cgos__expm1 float64 = 1.4426950408889634
var Q1_cgos__expm1 float64 = -0.033333333333333132
var Q2_cgos__expm1 float64 = 0.0015873015872548146
var Q3_cgos__expm1 float64 = -7.9365075786748794e-5
var Q4_cgos__expm1 float64 = 4.0082178273293624e-6
var Q5_cgos__expm1 float64 = -2.0109921818362437e-7

func Expm1(x float64) float64 {
	var y float64
	var hi float64
	var lo float64
	var c float64
	var t float64
	var e float64
	var hxs float64
	var hfx float64
	var r1 float64
	var twopk float64
	type _cgoa_18_expm1 struct {
		f float64
	}
	var u _cgoa_18_expm1
	u.f = x
	var hx uint32 = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(32) & uint64(2147483647))
	var k int32
	var sign int32 = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(63))
	if hx >= uint32(1078159482) {
		if func() int32 {
			if 8 == 4 {
				return func() int32 {
					if X__FLOAT_BITS(float32(x))&uint32(2147483647) > uint32(2139095040) {
						return 1
					} else {
						return 0
					}
				}()
			} else {
				return func() int32 {
					if 8 == 8 {
						return func() int32 {
							if X__DOUBLE_BITS(x)&9223372036854775807 > 9218868437227405312 {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if X__fpclassifyl(float64(x)) == int32(0) {
								return 1
							} else {
								return 0
							}
						}()
					}
				}()
			}
		}() != 0 {
			return x
		}
		if sign != 0 {
			return float64(-1)
		}
		if x > o_threshold_cgos__expm1 {
			x *= float64(8.9884656743115795e+307)
			return x
		}
	}
	if hx > uint32(1071001154) {
		if hx < uint32(1072734898) {
			if !(sign != 0) {
				hi = x - ln2_hi_cgos__expm1
				lo = ln2_lo_cgos__expm1
				k = int32(1)
			} else {
				hi = x + ln2_hi_cgos__expm1
				lo = -ln2_lo_cgos__expm1
				k = -1
			}
		} else {
			k = int32(invln2_cgos__expm1*x + func() float64 {
				if sign != 0 {
					return -0.5
				} else {
					return 0.5
				}
			}())
			t = float64(k)
			hi = x - t*ln2_hi_cgos__expm1
			lo = t * ln2_lo_cgos__expm1
		}
		x = hi - lo
		c = hi - x - lo
	} else if hx < uint32(1016070144) {
		if hx < uint32(1048576) {
			for {
				if 4 == 4 {
					fp_force_evalf(float32(x))
				} else if 4 == 8 {
					fp_force_eval(float64(float32(x)))
				} else {
					fp_force_evall(float64(float32(x)))
				}
				if true {
					break
				}
			}
		}
		return x
	} else {
		k = int32(0)
	}
	hfx = 0.5 * x
	hxs = x * hfx
	r1 = 1 + hxs*(Q1_cgos__expm1+hxs*(Q2_cgos__expm1+hxs*(Q3_cgos__expm1+hxs*(Q4_cgos__expm1+hxs*Q5_cgos__expm1))))
	t = 3 - r1*hfx
	e = hxs * ((r1 - t) / (6 - x*t))
	if k == int32(0) {
		return x - (x*e - hxs)
	}
	e = x*(e-c) - c
	e -= hxs
	if k == -1 {
		return 0.5*(x-e) - 0.5
	}
	if k == int32(1) {
		if x < -0.25 {
			return -2 * (e - (x + 0.5))
		}
		return 1 + 2*(x-e)
	}
	*(*uint64)(unsafe.Pointer(&u)) = uint64(int32(1023)+k) << int32(52)
	twopk = u.f
	if k < int32(0) || k > int32(56) {
		y = x - e + 1
		if k == int32(1024) {
			y = y * 2 * 8.9884656743115795e+307
		} else {
			y = y * twopk
		}
		return y - 1
	}
	*(*uint64)(unsafe.Pointer(&u)) = uint64(int32(1023)-k) << int32(52)
	if k < int32(20) {
		y = (x - e + (float64(int32(1)) - u.f)) * twopk
	} else {
		y = (x - (e + u.f) + float64(int32(1))) * twopk
	}
	return y
}
