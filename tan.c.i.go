package libc

import unsafe "unsafe"

func Tan(x float64) float64 {
	var y [2]float64
	var ix uint32
	var n uint32
	for {
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_18_tan{x})) >> int32(32))
		if true {
			break
		}
	}
	ix &= uint32(2147483647)
	if ix <= uint32(1072243195) {
		if ix < uint32(1044381696) {
			for {
				if 8 == 4 {
					fp_force_evalf(float32(func() float64 {
						if ix < uint32(1048576) {
							return x / float64(1.329228e+36)
						} else {
							return x + float64(1.329228e+36)
						}
					}()))
				} else if 8 == 8 {
					fp_force_eval(func() float64 {
						if ix < uint32(1048576) {
							return x / float64(1.329228e+36)
						} else {
							return x + float64(1.329228e+36)
						}
					}())
				} else {
					fp_force_evall(float64(func() float64 {
						if ix < uint32(1048576) {
							return x / float64(1.329228e+36)
						} else {
							return x + float64(1.329228e+36)
						}
					}()))
				}
				if true {
					break
				}
			}
			return x
		}
		return __tan(x, 0, int32(0))
	}
	if ix >= uint32(2146435072) {
		return x - x
	}
	n = uint32(__rem_pio2(x, (*float64)(unsafe.Pointer(&y))))
	return __tan(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(0))*8)), *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(1))*8)), int32(n&uint32(1)))
}

type _cgoz_18_tan struct {
	_f float64
}
