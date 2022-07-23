package libc

import unsafe "unsafe"

func Exp10f(x float32) float32 {
	var n float32
	var y float32 = Modff(x, &n)
	type _cgoa_18_exp10f struct {
		f float32
	}
	var u _cgoa_18_exp10f
	u.f = n
	if *(*uint32)(unsafe.Pointer(&u))>>int32(23)&uint32(255) < uint32(130) {
		if !(y != 0) {
			return *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&_cgos_p10__exp10f)))) + uintptr(int32(n)+int32(7))*4))
		}
		y = Exp2f(3.32192802 * y)
		return y * *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&_cgos_p10__exp10f)))) + uintptr(int32(n)+int32(7))*4))
	}
	return float32(Exp2(3.3219280948873622 * float64(x)))
}

var _cgos_p10__exp10f [15]float32 = [15]float32{1.00000001e-7, 9.99999997e-7, 9.99999974e-6, 9.99999974e-5, 0.00100000005, 0.00999999977, 0.100000001, float32(int32(1)), float32(10), float32(100), float32(1000), float32(1.0e+4), float32(1.0e+5), float32(1.0e+6), float32(1.0e+7)}
