package libc

import unsafe "unsafe"

func Log2f(x float32) float32 {
	var z float64
	var r float64
	var r2 float64
	var p float64
	var y float64
	var y0 float64
	var invc float64
	var logc float64
	var ix uint32
	var iz uint32
	var top uint32
	var tmp uint32
	var k int32
	var i int32
	ix = *(*uint32)(unsafe.Pointer(&_cgoz_19_log2f{x}))
	if int32(1) != 0 && func() int64 {
		if ix == uint32(1065353216) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		return float32(int32(0))
	}
	if func() int64 {
		if ix-uint32(8388608) >= uint32(2130706432) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if ix*uint32(2) == uint32(0) {
			return __math_divzerof(uint32(1))
		}
		if ix == uint32(2139095040) {
			return x
		}
		if ix&uint32(2147483648) != 0 || ix*uint32(2) >= uint32(4278190080) {
			return __math_invalidf(x)
		}
		ix = *(*uint32)(unsafe.Pointer(&_cgoz_20_log2f{x * 8388608.0}))
		ix -= uint32(192937984)
	}
	tmp = ix - uint32(1060306944)
	i = int32(tmp >> 19 % uint32(16))
	top = tmp & uint32(4286578688)
	iz = ix - top
	k = int32(tmp) >> int32(23)
	invc = (*(*_cgoa_18_log2f)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_18_log2f)(unsafe.Pointer(&__log2f_data.tab)))) + uintptr(i)*16))).invc
	logc = (*(*_cgoa_18_log2f)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_18_log2f)(unsafe.Pointer(&__log2f_data.tab)))) + uintptr(i)*16))).logc
	z = float64(*(*float32)(unsafe.Pointer(&_cgoz_21_log2f{iz})))
	r = z*invc - float64(int32(1))
	y0 = logc + float64(k)
	r2 = r * r
	y = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2f_data.poly)))) + uintptr(int32(1))*8))*r + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2f_data.poly)))) + uintptr(int32(2))*8))
	y = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2f_data.poly)))) + uintptr(int32(0))*8))*r2 + y
	p = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2f_data.poly)))) + uintptr(int32(3))*8))*r + y0
	y = y*r2 + p
	return eval_as_float(float32(y))
}

type _cgoz_19_log2f struct {
	_f float32
}
type _cgoz_20_log2f struct {
	_f float32
}
type _cgoz_21_log2f struct {
	_i uint32
}
