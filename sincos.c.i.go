package libc

import unsafe "unsafe"

func Sincos(x float64, sin *float64, cos *float64) {
	var y [2]float64
	var s float64
	var c float64
	var ix uint32
	var n uint32
	for {
		ix = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_18_sincos{x})) >> int32(32))
		if true {
			break
		}
	}
	ix &= uint32(2147483647)
	if ix <= uint32(1072243195) {
		if ix < uint32(1044816030) {
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
			*sin = x
			*cos = float64(1)
			return
		}
		*sin = __sin(x, 0, int32(0))
		*cos = __cos(x, 0)
		return
	}
	if ix >= uint32(2146435072) {
		*sin = func() (_cgo_ret float64) {
			_cgo_addr := &*cos
			*_cgo_addr = x - x
			return *_cgo_addr
		}()
		return
	}
	n = uint32(__rem_pio2(x, (*float64)(unsafe.Pointer(&y))))
	s = __sin(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(0))*8)), *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(1))*8)), int32(1))
	c = __cos(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(0))*8)), *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&y)))) + uintptr(int32(1))*8)))
	switch n & uint32(3) {
	case uint32(0):
		*sin = s
		*cos = c
		break
	case uint32(1):
		*sin = c
		*cos = -s
		break
	case uint32(2):
		*sin = -s
		*cos = -c
		break
	case uint32(3):
		fallthrough
	default:
		*sin = -c
		*cos = s
		break
	}
}

type _cgoz_18_sincos struct {
	_f float64
}
