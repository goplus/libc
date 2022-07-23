package libc

import unsafe "unsafe"

var _cgos_pio2_hi__acos float64 = 1.5707963267948966
var _cgos_pio2_lo__acos float64 = 6.123233995736766e-17
var _cgos_pS0__acos float64 = 0.16666666666666666
var _cgos_pS1__acos float64 = -0.32556581862240092
var _cgos_pS2__acos float64 = 0.20121253213486293
var _cgos_pS3__acos float64 = -0.040055534500679411
var _cgos_pS4__acos float64 = 7.9153499428981453e-4
var _cgos_pS5__acos float64 = 3.4793310759602117e-5
var _cgos_qS1__acos float64 = -2.4033949117344142
var _cgos_qS2__acos float64 = 2.0209457602335057
var _cgos_qS3__acos float64 = -0.68828397160545329
var _cgos_qS4__acos float64 = 0.077038150555901935

func _cgos_R__acos(z float64) float64 {
	var p float64
	var q float64
	p = z * (_cgos_pS0__acos + z*(_cgos_pS1__acos+z*(_cgos_pS2__acos+z*(_cgos_pS3__acos+z*(_cgos_pS4__acos+z*_cgos_pS5__acos)))))
	q = 1 + z*(_cgos_qS1__acos+z*(_cgos_qS2__acos+z*(_cgos_qS3__acos+z*_cgos_qS4__acos)))
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
				return float64(int32(2))*_cgos_pio2_hi__acos + float64(7.52316385e-37)
			}
			return float64(int32(0))
		}
		return float64(int32(0)) / (x - x)
	}
	if ix < uint32(1071644672) {
		if ix <= uint32(1012924416) {
			return _cgos_pio2_hi__acos + float64(7.52316385e-37)
		}
		return _cgos_pio2_hi__acos - (x - (_cgos_pio2_lo__acos - x*_cgos_R__acos(x*x)))
	}
	if hx>>int32(31) != 0 {
		z = (1 + x) * 0.5
		s = Sqrt(z)
		w = _cgos_R__acos(z)*s - _cgos_pio2_lo__acos
		return float64(int32(2)) * (_cgos_pio2_hi__acos - (s + w))
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
	w = _cgos_R__acos(z)*s + c
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
