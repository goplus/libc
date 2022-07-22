package libc

import unsafe "unsafe"

func Fmaf(x float32, y float32, z float32) float32 {
	var xy float64
	var result float64
	type _cgoa_19_fmaf struct {
		f float64
	}
	var u _cgoa_19_fmaf
	var e int32
	xy = float64(x) * float64(y)
	result = xy + float64(z)
	u.f = result
	e = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(52) & uint64(2047))
	if *(*uint64)(unsafe.Pointer(&u))&uint64(536870911) != uint64(268435456) || e == int32(2047) || result-xy == float64(z) && result-float64(z) == xy || Fegetround() != int32(0) {
		z = float32(result)
		return z
	}
	var vxy float64 = xy
	_ = vxy
	var adjusted_result float64 = vxy + float64(z)
	Fesetround(int32(0))
	if result == adjusted_result {
		u.f = adjusted_result
		*(*uint64)(unsafe.Pointer(&u))++
		adjusted_result = u.f
	}
	z = float32(adjusted_result)
	return z
}
