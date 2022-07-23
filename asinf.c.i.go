package libc

import unsafe "unsafe"

var pio2_cgos__asinf float64 = 1.5707963267948966
var pS0_cgos__asinf float32 = float32(0.16666586696999999)
var pS1_cgos__asinf float32 = float32(-0.042743422091000002)
var pS2_cgos__asinf float32 = float32(-0.0086563630029999998)
var qS1_cgos__asinf float32 = float32(-0.7066296339)

func R_cgos__asinf(z float32) float32 {
	var p float32
	var q float32
	p = z * (pS0_cgos__asinf + z*(pS1_cgos__asinf+z*pS2_cgos__asinf))
	q = 1 + z*qS1_cgos__asinf
	return p / q
}
func Asinf(x float32) float32 {
	var s float64
	var z float32
	var hx uint32
	var ix uint32
	for {
		hx = *(*uint32)(unsafe.Pointer(&_cgoz_18_asinf{x}))
		if true {
			break
		}
	}
	ix = hx & uint32(2147483647)
	if ix >= uint32(1065353216) {
		if ix == uint32(1065353216) {
			return float32(float64(x)*pio2_cgos__asinf + float64(7.52316385e-37))
		}
		return float32(int32(0)) / (x - x)
	}
	if ix < uint32(1056964608) {
		if ix < uint32(964689920) && ix >= uint32(8388608) {
			return x
		}
		return x + x*R_cgos__asinf(x*x)
	}
	z = (float32(int32(1)) - Fabsf(x)) * 0.5
	s = Sqrt(float64(z))
	x = float32(pio2_cgos__asinf - float64(int32(2))*(s+s*float64(R_cgos__asinf(z))))
	if hx>>int32(31) != 0 {
		return -x
	}
	return x
}

type _cgoz_18_asinf struct {
	_f float32
}
