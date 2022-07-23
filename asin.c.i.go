package libc

import unsafe "unsafe"

var pio2_hi_cgos__asin float64 = 1.5707963267948966
var pio2_lo_cgos__asin float64 = 6.123233995736766e-17
var pS0_cgos__asin float64 = 0.16666666666666666
var pS1_cgos__asin float64 = -0.32556581862240092
var pS2_cgos__asin float64 = 0.20121253213486293
var pS3_cgos__asin float64 = -0.040055534500679411
var pS4_cgos__asin float64 = 7.9153499428981453e-4
var pS5_cgos__asin float64 = 3.4793310759602117e-5
var qS1_cgos__asin float64 = -2.4033949117344142
var qS2_cgos__asin float64 = 2.0209457602335057
var qS3_cgos__asin float64 = -0.68828397160545329
var qS4_cgos__asin float64 = 0.077038150555901935

func R_cgos__asin(z float64) float64 {
	var p float64
	var q float64
	p = z * (pS0_cgos__asin + z*(pS1_cgos__asin+z*(pS2_cgos__asin+z*(pS3_cgos__asin+z*(pS4_cgos__asin+z*pS5_cgos__asin)))))
	q = 1 + z*(qS1_cgos__asin+z*(qS2_cgos__asin+z*(qS3_cgos__asin+z*qS4_cgos__asin)))
	return p / q
}
func Asin(x float64) float64 {
	var z float64
	var r float64
	var s float64
	var hx uint32
	var ix uint32
	for {
		hx = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_18_asin{x})) >> int32(32))
		if true {
			break
		}
	}
	ix = hx & uint32(2147483647)
	if ix >= uint32(1072693248) {
		var lx uint32
		for {
			lx = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_19_asin{x})))
			if true {
				break
			}
		}
		if ix-uint32(1072693248)|lx == uint32(0) {
			return x*pio2_hi_cgos__asin + float64(7.52316385e-37)
		}
		return float64(int32(0)) / (x - x)
	}
	if ix < uint32(1071644672) {
		if ix < uint32(1045430272) && ix >= uint32(1048576) {
			return x
		}
		return x + x*R_cgos__asin(x*x)
	}
	z = (float64(int32(1)) - Fabs(x)) * 0.5
	s = Sqrt(z)
	r = R_cgos__asin(z)
	if ix >= uint32(1072640819) {
		x = pio2_hi_cgos__asin - (float64(int32(2))*(s+s*r) - pio2_lo_cgos__asin)
	} else {
		var f float64
		var c float64
		f = s
		for {
			f = *(*float64)(unsafe.Pointer(&_cgoz_20_asin{uint64(*(*uint64)(unsafe.Pointer(&_cgoz_21_asin{f}))>>int32(32))<<int32(32) | uint64(0)}))
			if true {
				break
			}
		}
		c = (z - f*f) / (s + f)
		x = 0.5*pio2_hi_cgos__asin - (float64(int32(2))*s*r - (pio2_lo_cgos__asin - float64(int32(2))*c) - (0.5*pio2_hi_cgos__asin - float64(int32(2))*f))
	}
	if hx>>int32(31) != 0 {
		return -x
	}
	return x
}

type _cgoz_18_asin struct {
	_f float64
}
type _cgoz_19_asin struct {
	_f float64
}
type _cgoz_20_asin struct {
	_i uint64
}
type _cgoz_21_asin struct {
	_f float64
}
