package libc

import unsafe "unsafe"

var _cgos_T___tandf [6]float64 = [6]float64{0.3333313950307914, 0.13339200271297674, 0.053381237844567039, 0.024528318116654728, 0.002974357433599673, 0.0094656478494367316}

func __tandf(x float64, odd int32) float32 {
	var z float64
	var r float64
	var w float64
	var s float64
	var t float64
	var u float64
	z = x * x
	r = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_T___tandf)))) + uintptr(int32(4))*8)) + z**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_T___tandf)))) + uintptr(int32(5))*8))
	t = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_T___tandf)))) + uintptr(int32(2))*8)) + z**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_T___tandf)))) + uintptr(int32(3))*8))
	w = z * z
	s = z * x
	u = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_T___tandf)))) + uintptr(int32(0))*8)) + z**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_T___tandf)))) + uintptr(int32(1))*8))
	r = x + s*u + s*w*(t+w*r)
	return float32(func() float64 {
		if odd != 0 {
			return -1.0 / r
		} else {
			return r
		}
	}())
}
