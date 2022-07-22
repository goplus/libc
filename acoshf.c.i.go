package libc

import unsafe "unsafe"

func Acoshf(x float32) float32 {
	type _cgoa_18_acoshf struct {
		f float32
	}
	var u _cgoa_18_acoshf
	u.f = x
	var a uint32 = *(*uint32)(unsafe.Pointer(&u)) & uint32(2147483647)
	if a < uint32(1073741824) {
		return Log1pf(x - float32(int32(1)) + Sqrtf((x-float32(int32(1)))*(x-float32(int32(1)))+float32(int32(2))*(x-float32(int32(1)))))
	}
	if a < uint32(1166016512) {
		return Logf(float32(int32(2))*x - float32(int32(1))/(x+Sqrtf(x*x-float32(int32(1)))))
	}
	return Logf(x) + 0.693147182
}
