package libc

import unsafe "unsafe"

var pio2_hi_cgos__acos float64 = 1.5707963267948966
var pio2_lo_cgos__acos float64 = 6.123233995736766e-17
var pS0_cgos__acos float64 = 0.16666666666666666
var pS1_cgos__acos float64 = -0.32556581862240092
var pS2_cgos__acos float64 = 0.20121253213486293
var pS3_cgos__acos float64 = -0.040055534500679411
var pS4_cgos__acos float64 = 7.9153499428981453e-4
var pS5_cgos__acos float64 = 3.4793310759602117e-5
var qS1_cgos__acos float64 = -2.4033949117344142
var qS2_cgos__acos float64 = 2.0209457602335057
var qS3_cgos__acos float64 = -0.68828397160545329
var qS4_cgos__acos float64 = 0.077038150555901935

func R_cgos__acos(z float64) float64 {
	var p float64
	var q float64
	p = z * (pS0_cgos__acos + z*(pS1_cgos__acos+z*(pS2_cgos__acos+z*(pS3_cgos__acos+z*(pS4_cgos__acos+z*pS5_cgos__acos)))))
	q = 1 + z*(qS1_cgos__acos+z*(qS2_cgos__acos+z*(qS3_cgos__acos+z*qS4_cgos__acos)))
	return p / q
}
func Acos(x float64) float64 {
	var z float64
	var w float64
	var s float64
	var c float64
	var df float64
	var hx uint32
	var ix uint32
	for {
		hx = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_18_acos{x})) >> int32(32))
		if true {
			break
		}
	}
	ix = hx & uint32(2147483647)
	if ix >= uint32(1072693248) {
		var lx uint32
		for {
			lx = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_19_acos{x})))
			if true {
				break
			}
		}
		if ix-uint32(1072693248)|lx == uint32(0) {
			if hx>>int32(31) != 0 {
				return float64(int32(2))*pio2_hi_cgos__acos + float64(7.52316385e-37)
			}
			return float64(int32(0))
		}
		return float64(int32(0)) / (x - x)
	}
	if ix < uint32(1071644672) {
		if ix <= uint32(1012924416) {
			return pio2_hi_cgos__acos + float64(7.52316385e-37)
		}
		return pio2_hi_cgos__acos - (x - (pio2_lo_cgos__acos - x*R_cgos__acos(x*x)))
	}
	if hx>>int32(31) != 0 {
		z = (1 + x) * 0.5
		s = Sqrt(z)
		w = R_cgos__acos(z)*s - pio2_lo_cgos__acos
		return float64(int32(2)) * (pio2_hi_cgos__acos - (s + w))
	}
	z = (1 - x) * 0.5
	s = Sqrt(z)
	df = s
	for {
		df = *(*float64)(unsafe.Pointer(&_cgoz_20_acos{uint64(*(*uint64)(unsafe.Pointer(&_cgoz_21_acos{df}))>>int32(32))<<int32(32) | uint64(0)}))
		if true {
			break
		}
	}
	c = (z - df*df) / (s + df)
	w = R_cgos__acos(z)*s + c
	return float64(int32(2)) * (df + w)
}

type _cgoz_18_acos struct {
	_f float64
}
type _cgoz_19_acos struct {
	_f float64
}
type _cgoz_20_acos struct {
	_i uint64
}
type _cgoz_21_acos struct {
	_f float64
}
