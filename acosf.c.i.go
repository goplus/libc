package libc

import unsafe "unsafe"

var pio2_hi_cgos__acosf float32 = float32(1.5707962513)
var pio2_lo_cgos__acosf float32 = float32(7.5497894159e-8)
var pS0_cgos__acosf float32 = float32(0.16666586696999999)
var pS1_cgos__acosf float32 = float32(-0.042743422091000002)
var pS2_cgos__acosf float32 = float32(-0.0086563630029999998)
var qS1_cgos__acosf float32 = float32(-0.7066296339)

func R_cgos__acosf(z float32) float32 {
	var p float32
	var q float32
	p = z * (pS0_cgos__acosf + z*(pS1_cgos__acosf+z*pS2_cgos__acosf))
	q = 1 + z*qS1_cgos__acosf
	return p / q
}
func Acosf(x float32) float32 {
	var z float32
	var w float32
	var s float32
	var c float32
	var df float32
	var hx uint32
	var ix uint32
	for {
		hx = *(*uint32)(unsafe.Pointer(&_cgoz_18_acosf{x}))
		if true {
			break
		}
	}
	ix = hx & uint32(2147483647)
	if ix >= uint32(1065353216) {
		if ix == uint32(1065353216) {
			if hx>>int32(31) != 0 {
				return float32(int32(2))*pio2_hi_cgos__acosf + 7.52316385e-37
			}
			return float32(int32(0))
		}
		return float32(int32(0)) / (x - x)
	}
	if ix < uint32(1056964608) {
		if ix <= uint32(847249408) {
			return pio2_hi_cgos__acosf + 7.52316385e-37
		}
		return pio2_hi_cgos__acosf - (x - (pio2_lo_cgos__acosf - x*R_cgos__acosf(x*x)))
	}
	if hx>>int32(31) != 0 {
		z = (float32(int32(1)) + x) * 0.5
		s = Sqrtf(z)
		w = R_cgos__acosf(z)*s - pio2_lo_cgos__acosf
		return float32(int32(2)) * (pio2_hi_cgos__acosf - (s + w))
	}
	z = (float32(int32(1)) - x) * 0.5
	s = Sqrtf(z)
	for {
		hx = *(*uint32)(unsafe.Pointer(&_cgoz_19_acosf{s}))
		if true {
			break
		}
	}
	for {
		df = *(*float32)(unsafe.Pointer(&_cgoz_20_acosf{hx & uint32(4294963200)}))
		if true {
			break
		}
	}
	c = (z - df*df) / (s + df)
	w = R_cgos__acosf(z)*s + c
	return float32(int32(2)) * (df + w)
}

type _cgoz_18_acosf struct {
	_f float32
}
type _cgoz_19_acosf struct {
	_f float32
}
type _cgoz_20_acosf struct {
	_i uint32
}
