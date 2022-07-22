package libc

import unsafe "unsafe"

var pio2_cgo18_asinf float64 = 1.5707963267948966
var pS0_cgo19_asinf float32 = float32(0.16666586696999999)
var pS1_cgo20_asinf float32 = float32(-0.042743422091000002)
var pS2_cgo21_asinf float32 = float32(-0.0086563630029999998)
var qS1_cgo22_asinf float32 = float32(-0.7066296339)

func R_cgo23_asinf(z float32) float32 {
	var p float32
	var q float32
	p = z * (pS0_cgo19_asinf + z*(pS1_cgo20_asinf+z*pS2_cgo21_asinf))
	q = 1 + z*qS1_cgo22_asinf
	return p / q
}
func Asinf(x float32) float32 {
	var s float64
	var z float32
	var hx uint32
	var ix uint32
	for {
		hx = *(*uint32)(unsafe.Pointer(&_cgoz_24_asinf{x}))
		if true {
			break
		}
	}
	ix = hx & uint32(2147483647)
	if ix >= uint32(1065353216) {
		if ix == uint32(1065353216) {
			return float32(float64(x)*pio2_cgo18_asinf + float64(7.52316385e-37))
		}
		return float32(int32(0)) / (x - x)
	}
	if ix < uint32(1056964608) {
		if ix < uint32(964689920) && ix >= uint32(8388608) {
			return x
		}
		return x + x*R_cgo23_asinf(x*x)
	}
	z = (float32(int32(1)) - Fabsf(x)) * 0.5
	s = Sqrt(float64(z))
	x = float32(pio2_cgo18_asinf - float64(int32(2))*(s+s*float64(R_cgo23_asinf(z))))
	if hx>>int32(31) != 0 {
		return -x
	}
	return x
}

type _cgoz_24_asinf struct {
	_f float32
}
