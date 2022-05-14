package libc

import unsafe "unsafe"

func Frexp(x float64, e *int32) float64 {
	type _cgoa_126 struct {
		d float64
	}
	var y _cgoa_126
	y.d = x
	var ee int32 = int32(*(*uint64)(unsafe.Pointer(&y)) >> int32(52) & uint64(2047))
	if !(ee != 0) {
		if x != 0 {
			x = Frexp(x*1.8446744073709552e+19, e)
			*e -= int32(64)
		} else {
			*e = int32(0)
		}
		return x
	} else if ee == int32(2047) {
		return x
	}
	*e = ee - int32(1022)
	*(*uint64)(unsafe.Pointer(&y)) &= uint64(9227875636482146303)
	*(*uint64)(unsafe.Pointer(&y)) |= uint64(4602678819172646912)
	return y.d
}
