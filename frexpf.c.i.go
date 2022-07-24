package libc

import unsafe "unsafe"

func Frexpf(x float32, e *int32) float32 {
	type _cgoa_18_frexpf struct {
		f float32
	}
	var y _cgoa_18_frexpf
	y.f = x
	var ee int32 = int32(*(*uint32)(unsafe.Pointer(&y)) >> int32(23) & uint32(255))
	if !(ee != 0) {
		if x != 0 {
			x = Frexpf(float32(float64(x)*1.8446744073709552e+19), e)
			*e -= int32(64)
		} else {
			*e = int32(0)
		}
		return x
	} else if ee == int32(255) {
		return x
	}
	*e = ee - int32(126)
	*(*uint32)(unsafe.Pointer(&y)) &= uint32(2155872255)
	*(*uint32)(unsafe.Pointer(&y)) |= uint32(1056964608)
	return y.f
}
