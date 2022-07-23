package libc

import unsafe "unsafe"

var T_cgos____tan [13]float64 = [13]float64{0.33333333333333409, 0.13333333333320124, 0.053968253976226052, 0.021869488294859542, 0.0088632398235993, 0.0035920791075913124, 0.0014562094543252903, 5.880412408202641e-4, 2.4646313481846991e-4, 7.8179444293955709e-5, 7.1407249138260819e-5, -1.8558637485527546e-5, 2.5907305186363371e-5}
var pio4_cgos____tan float64 = 0.78539816339744828
var pio4lo_cgos____tan float64 = 3.061616997868383e-17

func __tan(x float64, y float64, odd int32) float64 {
	var z float64
	var r float64
	var v float64
	var w float64
	var s float64
	var a float64
	var w0 float64
	var a0 float64
	var hx uint32
	var big int32
	var sign int32
	for {
		hx = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_18___tan{x})) >> int32(32))
		if true {
			break
		}
	}
	big = func() int32 {
		if hx&uint32(2147483647) >= uint32(1072010280) {
			return 1
		} else {
			return 0
		}
	}()
	if big != 0 {
		sign = int32(hx >> int32(31))
		if sign != 0 {
			x = -x
			y = -y
		}
		x = pio4_cgos____tan - x + (pio4lo_cgos____tan - y)
		y = float64(0)
	}
	z = x * x
	w = z * z
	r = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&T_cgos____tan)))) + uintptr(int32(1))*8)) + w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&T_cgos____tan)))) + uintptr(int32(3))*8))+w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&T_cgos____tan)))) + uintptr(int32(5))*8))+w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&T_cgos____tan)))) + uintptr(int32(7))*8))+w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&T_cgos____tan)))) + uintptr(int32(9))*8))+w**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&T_cgos____tan)))) + uintptr(int32(11))*8))))))
	v = z * (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&T_cgos____tan)))) + uintptr(int32(2))*8)) + w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&T_cgos____tan)))) + uintptr(int32(4))*8))+w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&T_cgos____tan)))) + uintptr(int32(6))*8))+w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&T_cgos____tan)))) + uintptr(int32(8))*8))+w*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&T_cgos____tan)))) + uintptr(int32(10))*8))+w**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&T_cgos____tan)))) + uintptr(int32(12))*8)))))))
	s = z * x
	r = y + z*(s*(r+v)+y) + s**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&T_cgos____tan)))) + uintptr(int32(0))*8))
	w = x + r
	if big != 0 {
		s = float64(int32(1) - int32(2)*odd)
		v = s - 2*(x+(r-w*w/(w+s)))
		return func() float64 {
			if sign != 0 {
				return -v
			} else {
				return v
			}
		}()
	}
	if !(odd != 0) {
		return w
	}
	w0 = w
	for {
		w0 = *(*float64)(unsafe.Pointer(&_cgoz_19___tan{uint64(*(*uint64)(unsafe.Pointer(&_cgoz_20___tan{w0}))>>int32(32))<<int32(32) | uint64(0)}))
		if true {
			break
		}
	}
	v = r - (w0 - x)
	a0 = func() (_cgo_ret float64) {
		_cgo_addr := &a
		*_cgo_addr = -1 / w
		return *_cgo_addr
	}()
	for {
		a0 = *(*float64)(unsafe.Pointer(&_cgoz_21___tan{uint64(*(*uint64)(unsafe.Pointer(&_cgoz_22___tan{a0}))>>int32(32))<<int32(32) | uint64(0)}))
		if true {
			break
		}
	}
	return a0 + a*(1+a0*w0+a0*v)
}

type _cgoz_18___tan struct {
	_f float64
}
type _cgoz_19___tan struct {
	_i uint64
}
type _cgoz_20___tan struct {
	_f float64
}
type _cgoz_21___tan struct {
	_i uint64
}
type _cgoz_22___tan struct {
	_f float64
}
