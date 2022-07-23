package libc

import unsafe "unsafe"

var o_threshold_cgos__expm1f float32 = float32(88.721679687999995)
var ln2_hi_cgos__expm1f float32 = float32(0.69313812255999996)
var ln2_lo_cgos__expm1f float32 = float32(9.0580006144999996e-6)
var invln2_cgos__expm1f float32 = float32(1.4426950216000001)
var Q1_cgos__expm1f float32 = float32(-0.033333212137000003)
var Q2_cgos__expm1f float32 = float32(0.0015807170421000001)

func Expm1f(x float32) float32 {
	var y float32
	var hi float32
	var lo float32
	var c float32
	var t float32
	var e float32
	var hxs float32
	var hfx float32
	var r1 float32
	var twopk float32
	type _cgoa_18_expm1f struct {
		f float32
	}
	var u _cgoa_18_expm1f
	u.f = x
	var hx uint32 = *(*uint32)(unsafe.Pointer(&u)) & uint32(2147483647)
	var k int32
	var sign int32 = int32(*(*uint32)(unsafe.Pointer(&u)) >> int32(31))
	if hx >= uint32(1100331076) {
		if hx > uint32(2139095040) {
			return x
		}
		if sign != 0 {
			return float32(-1)
		}
		if x > o_threshold_cgos__expm1f {
			x *= float32(1.70141183e+38)
			return x
		}
	}
	if hx > uint32(1051816472) {
		if hx < uint32(1065686418) {
			if !(sign != 0) {
				hi = x - ln2_hi_cgos__expm1f
				lo = ln2_lo_cgos__expm1f
				k = int32(1)
			} else {
				hi = x + ln2_hi_cgos__expm1f
				lo = -ln2_lo_cgos__expm1f
				k = -1
			}
		} else {
			k = int32(invln2_cgos__expm1f*x + func() float32 {
				if sign != 0 {
					return -0.5
				} else {
					return 0.5
				}
			}())
			t = float32(k)
			hi = x - t*ln2_hi_cgos__expm1f
			lo = t * ln2_lo_cgos__expm1f
		}
		x = hi - lo
		c = hi - x - lo
	} else if hx < uint32(855638016) {
		if hx < uint32(8388608) {
			for {
				if 4 == 4 {
					fp_force_evalf(x * x)
				} else if 4 == 8 {
					fp_force_eval(float64(x * x))
				} else {
					fp_force_evall(float64(x * x))
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
	r1 = 1 + hxs*(Q1_cgos__expm1f+hxs*Q2_cgos__expm1f)
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
	*(*uint32)(unsafe.Pointer(&u)) = uint32((int32(127) + k) << int32(23))
	twopk = u.f
	if k < int32(0) || k > int32(56) {
		y = x - e + 1
		if k == int32(128) {
			y = y * 2 * 1.70141183e+38
		} else {
			y = y * twopk
		}
		return y - 1
	}
	*(*uint32)(unsafe.Pointer(&u)) = uint32((int32(127) - k) << int32(23))
	if k < int32(23) {
		y = (x - e + (float32(int32(1)) - u.f)) * twopk
	} else {
		y = (x - (e + u.f) + float32(int32(1))) * twopk
	}
	return y
}
