package libc

import unsafe "unsafe"

func top16_cgos__log2(x float64) uint32 {
	return uint32(*(*uint64)(unsafe.Pointer(&_cgoz_20_log2{x})) >> int32(48))
}

type _cgoz_20_log2 struct {
	_f float64
}

func Log2(x float64) float64 {
	var z float64
	var r float64
	var r2 float64
	var r4 float64
	var y float64
	var invc float64
	var logc float64
	var kd float64
	var hi float64
	var lo float64
	var t1 float64
	var t2 float64
	var t3 float64
	var p float64
	var ix uint64
	var iz uint64
	var tmp uint64
	var top uint32
	var k int32
	var i int32
	ix = *(*uint64)(unsafe.Pointer(&_cgoz_21_log2{x}))
	top = top16_cgos__log2(x)
	if func() int64 {
		if ix-*(*uint64)(unsafe.Pointer(&_cgoz_22_log2{1 - 0.042397022247314453})) < *(*uint64)(unsafe.Pointer(&_cgoz_23_log2{1 + 0.044274330139160156}))-*(*uint64)(unsafe.Pointer(&_cgoz_24_log2{1 - 0.042397022247314453})) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if int32(1) != 0 && func() int64 {
			if ix == *(*uint64)(unsafe.Pointer(&_cgoz_25_log2{1})) {
				return 1
			} else {
				return 0
			}
		}() == int64(0) {
			return float64(int32(0))
		}
		r = x - 1
		var rhi float64
		var rlo float64
		rhi = *(*float64)(unsafe.Pointer(&_cgoz_26_log2{*(*uint64)(unsafe.Pointer(&_cgoz_27_log2{r})) & 18446744069414584320}))
		rlo = r - rhi
		hi = rhi * __log2_data.invln2hi
		lo = rlo*__log2_data.invln2hi + r*__log2_data.invln2lo
		r2 = r * r
		r4 = r2 * r2
		p = r2 * (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly1)))) + uintptr(int32(0))*8)) + r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly1)))) + uintptr(int32(1))*8)))
		y = hi + p
		lo += hi - y + p
		lo += r4 * (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly1)))) + uintptr(int32(2))*8)) + r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly1)))) + uintptr(int32(3))*8)) + r2*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly1)))) + uintptr(int32(4))*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly1)))) + uintptr(int32(5))*8))) + r4*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly1)))) + uintptr(int32(6))*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly1)))) + uintptr(int32(7))*8))+r2*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly1)))) + uintptr(int32(8))*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly1)))) + uintptr(int32(9))*8)))))
		y += lo
		return eval_as_double(y)
	}
	if func() int64 {
		if top-uint32(16) >= uint32(32736) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if ix*uint64(2) == uint64(0) {
			return __math_divzero(uint32(1))
		}
		if ix == *(*uint64)(unsafe.Pointer(&_cgoz_28_log2{float64(X__builtin_inff())})) {
			return x
		}
		if top&uint32(32768) != 0 || top&uint32(32752) == uint32(32752) {
			return __math_invalid(x)
		}
		ix = *(*uint64)(unsafe.Pointer(&_cgoz_29_log2{x * 4503599627370496}))
		ix -= 234187180623265792
	}
	tmp = ix - uint64(4604367669032910848)
	i = int32(tmp >> 46 % uint64(64))
	k = int32(int64(tmp) >> int32(52))
	iz = ix - tmp&18442240474082181120
	invc = (*(*_cgoa_18_log2)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_18_log2)(unsafe.Pointer(&__log2_data.tab)))) + uintptr(i)*16))).invc
	logc = (*(*_cgoa_18_log2)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_18_log2)(unsafe.Pointer(&__log2_data.tab)))) + uintptr(i)*16))).logc
	z = *(*float64)(unsafe.Pointer(&_cgoz_30_log2{iz}))
	kd = float64(k)
	var rhi float64
	var rlo float64
	r = (z - (*(*_cgoa_19_log2)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_19_log2)(unsafe.Pointer(&__log2_data.tab2)))) + uintptr(i)*16))).chi - (*(*_cgoa_19_log2)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_19_log2)(unsafe.Pointer(&__log2_data.tab2)))) + uintptr(i)*16))).clo) * invc
	rhi = *(*float64)(unsafe.Pointer(&_cgoz_31_log2{*(*uint64)(unsafe.Pointer(&_cgoz_32_log2{r})) & 18446744069414584320}))
	rlo = r - rhi
	t1 = rhi * __log2_data.invln2hi
	t2 = rlo*__log2_data.invln2hi + r*__log2_data.invln2lo
	t3 = kd + logc
	hi = t3 + t1
	lo = t3 - hi + t1 + t2
	r2 = r * r
	r4 = r2 * r2
	p = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly)))) + uintptr(int32(0))*8)) + r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly)))) + uintptr(int32(1))*8)) + r2*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly)))) + uintptr(int32(2))*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly)))) + uintptr(int32(3))*8))) + r4*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly)))) + uintptr(int32(4))*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log2_data.poly)))) + uintptr(int32(5))*8)))
	y = lo + r2*p + hi
	return eval_as_double(y)
}

type _cgoz_21_log2 struct {
	_f float64
}
type _cgoz_22_log2 struct {
	_f float64
}
type _cgoz_23_log2 struct {
	_f float64
}
type _cgoz_24_log2 struct {
	_f float64
}
type _cgoz_25_log2 struct {
	_f float64
}
type _cgoz_26_log2 struct {
	_i uint64
}
type _cgoz_27_log2 struct {
	_f float64
}
type _cgoz_28_log2 struct {
	_f float64
}
type _cgoz_29_log2 struct {
	_f float64
}
type _cgoz_30_log2 struct {
	_i uint64
}
type _cgoz_31_log2 struct {
	_i uint64
}
type _cgoz_32_log2 struct {
	_f float64
}
