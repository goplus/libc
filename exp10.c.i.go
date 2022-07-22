package libc

import unsafe "unsafe"

func Exp10(x float64) float64 {
	var n float64
	var y float64 = Modf(x, &n)
	type _cgoa_19_exp10 struct {
		f float64
	}
	var u _cgoa_19_exp10
	u.f = n
	if *(*uint64)(unsafe.Pointer(&u))>>int32(52)&uint64(2047) < uint64(1027) {
		if !(y != 0) {
			return *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&p10_cgo18_exp10)))) + uintptr(int32(n)+int32(15))*8))
		}
		y = Exp2(3.3219280948873622 * y)
		return y * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&p10_cgo18_exp10)))) + uintptr(int32(n)+int32(15))*8))
	}
	return Pow(10, x)
}

var p10_cgo18_exp10 [31]float64 = [31]float64{1.0000000000000001e-15, 9.9999999999999999e-15, 1.0e-13, 9.9999999999999998e-13, 9.9999999999999993e-12, 1.0e-10, 1.0000000000000001e-9, 1.0e-8, 9.9999999999999995e-8, 9.9999999999999995e-7, 1.0000000000000001e-5, 1.0e-4, 0.001, 0.01, 0.10000000000000001, float64(int32(1)), 10, 100, 1000, 1.0e+4, 1.0e+5, 1.0e+6, 1.0e+7, 1.0e+8, 1.0e+9, 1.0e+10, 1.0e+11, 1.0e+12, 1.0e+13, 1.0e+14, 1.0e+15}
