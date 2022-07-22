package libc

import unsafe "unsafe"

func Acosh(x float64) float64 {
	type _cgoa_18_acosh struct {
		f float64
	}
	var u _cgoa_18_acosh
	u.f = x
	var e uint32 = uint32(*(*uint64)(unsafe.Pointer(&u)) >> int32(52) & uint64(2047))
	if e < uint32(1024) {
		return Log1p(x - float64(int32(1)) + Sqrt((x-float64(int32(1)))*(x-float64(int32(1)))+float64(int32(2))*(x-float64(int32(1)))))
	}
	if e < uint32(1049) {
		return Log(float64(int32(2))*x - float64(int32(1))/(x+Sqrt(x*x-float64(int32(1)))))
	}
	return Log(x) + 0.69314718055994529
}
