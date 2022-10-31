package libc

import unsafe "unsafe"

func Cos(x float64) float64 {
	var y [2]float64
	var ix uint32
	var n uint32
	for {
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_18_cos{x})) >> int32(32))
		if true {
			break
		}
	}
	ix &= uint32(2147483647)
	if ix <= uint32(1072243195) {
		if ix < uint32(1044816030) {
			for {
				if 8 == 4 {
					fp_force_evalf(float32(x + float64(1.329228e+36)))
				} else if 8 == 8 {
					fp_force_eval(x + float64(1.329228e+36))
				} else {
					fp_force_evall(float64(x + float64(1.329228e+36)))
				}
				if true {
					break
				}
			}
			return float64(1.0)
		}
		return __cos(x, float64(int32(0)))
	}
	if ix >= uint32(2146435072) {
		return x - x
	}
	n = uint32(__rem_pio2(x, (*float64)(unsafe.Pointer(&y))))
	switch n & uint32(3) {
	case uint32(0):
		return __cos(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(0))*8)), *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(1))*8)))
	case uint32(1):
		return -__sin(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(0))*8)), *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(1))*8)), int32(1))
	case uint32(2):
		return -__cos(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(0))*8)), *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(1))*8)))
	default:
		return __sin(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(0))*8)), *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(1))*8)), int32(1))
	}
	return 0
}

type _cgoz_18_cos struct {
	_f float64
}
