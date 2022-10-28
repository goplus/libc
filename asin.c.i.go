package libc

import unsafe "unsafe"

var _cgos_pio2_hi_asin float64 = 1.5707963267948966
var _cgos_pio2_lo_asin float64 = 6.123233995736766e-17
var _cgos_pS0_asin float64 = 0.16666666666666666
var _cgos_pS1_asin float64 = -0.32556581862240092
var _cgos_pS2_asin float64 = 0.20121253213486293
var _cgos_pS3_asin float64 = -0.040055534500679411
var _cgos_pS4_asin float64 = 7.9153499428981453e-4
var _cgos_pS5_asin float64 = 3.4793310759602117e-5
var _cgos_qS1_asin float64 = -2.4033949117344142
var _cgos_qS2_asin float64 = 2.0209457602335057
var _cgos_qS3_asin float64 = -0.68828397160545329
var _cgos_qS4_asin float64 = 0.077038150555901935

func _cgos_R_asin(z float64) float64 {
	var p float64
	var q float64
	p = z * (_cgos_pS0_asin + z*(_cgos_pS1_asin+z*(_cgos_pS2_asin+z*(_cgos_pS3_asin+z*(_cgos_pS4_asin+z*_cgos_pS5_asin)))))
	q = 1.0 + z*(_cgos_qS1_asin+z*(_cgos_qS2_asin+z*(_cgos_qS3_asin+z*_cgos_qS4_asin)))
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
			return x*_cgos_pio2_hi_asin + float64(7.52316385e-37)
		}
		return float64(int32(0)) / (x - x)
	}
	if ix < uint32(1071644672) {
		if ix < uint32(1045430272) && ix >= uint32(1048576) {
			return x
		}
		return x + x*_cgos_R_asin(x*x)
	}
	z = (float64(int32(1)) - Fabs(x)) * 0.5
	s = Sqrt(z)
	r = _cgos_R_asin(z)
	if ix >= uint32(1072640819) {
		x = _cgos_pio2_hi_asin - (float64(int32(2))*(s+s*r) - _cgos_pio2_lo_asin)
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
		x = 0.5*_cgos_pio2_hi_asin - (float64(int32(2))*s*r - (_cgos_pio2_lo_asin - float64(int32(2))*c) - (0.5*_cgos_pio2_hi_asin - float64(int32(2))*f))
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
