package libc

import unsafe "unsafe"

func _cgos_top16__log(x float64) uint32 {
	return uint32(*(*uint64)(unsafe.Pointer(&_cgoz_20_log{x})) >> int32(48))
}

type _cgoz_20_log struct {
	_f float64
}

func Log(x float64) float64 {
	var w float64
	var z float64
	var r float64
	var r2 float64
	var r3 float64
	var y float64
	var invc float64
	var logc float64
	var kd float64
	var hi float64
	var lo float64
	var ix uint64
	var iz uint64
	var tmp uint64
	var top uint32
	var k int32
	var i int32
	ix = *(*uint64)(unsafe.Pointer(&_cgoz_21_log{x}))
	top = _cgos_top16__log(x)
	if func() int64 {
		if ix-*(*uint64)(unsafe.Pointer(&_cgoz_22_log{1 - 0.0625})) < *(*uint64)(unsafe.Pointer(&_cgoz_23_log{1 + 0.064697265625}))-*(*uint64)(unsafe.Pointer(&_cgoz_24_log{1 - 0.0625})) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if int32(1) != 0 && func() int64 {
			if ix == *(*uint64)(unsafe.Pointer(&_cgoz_25_log{1})) {
				return 1
			} else {
				return 0
			}
		}() == int64(0) {
			return float64(int32(0))
		}
		r = x - 1
		r2 = r * r
		r3 = r * r2
		y = r3 * (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly1)))) + uintptr(int32(1))*8)) + r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly1)))) + uintptr(int32(2))*8)) + r2**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly1)))) + uintptr(int32(3))*8)) + r3*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly1)))) + uintptr(int32(4))*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly1)))) + uintptr(int32(5))*8))+r2**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly1)))) + uintptr(int32(6))*8))+r3*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly1)))) + uintptr(int32(7))*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly1)))) + uintptr(int32(8))*8))+r2**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly1)))) + uintptr(int32(9))*8))+r3**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly1)))) + uintptr(int32(10))*8)))))
		w = r * 134217728
		var rhi float64 = r + w - w
		var rlo float64 = r - rhi
		w = rhi * rhi * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly1)))) + uintptr(int32(0))*8))
		hi = r + w
		lo = r - hi + w
		lo += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly1)))) + uintptr(int32(0))*8)) * rlo * (rhi + r)
		y += lo
		y += hi
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
		if ix == *(*uint64)(unsafe.Pointer(&_cgoz_26_log{float64(X__builtin_inff())})) {
			return x
		}
		if top&uint32(32768) != 0 || top&uint32(32752) == uint32(32752) {
			return __math_invalid(x)
		}
		ix = *(*uint64)(unsafe.Pointer(&_cgoz_27_log{x * 4503599627370496}))
		ix -= 234187180623265792
	}
	tmp = ix - uint64(4604367669032910848)
	i = int32(tmp >> 45 % uint64(128))
	k = int32(int64(tmp) >> int32(52))
	iz = ix - tmp&18442240474082181120
	invc = (*(*_cgoa_18_log)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_18_log)(unsafe.Pointer(&__log_data.tab)))) + uintptr(i)*16))).invc
	logc = (*(*_cgoa_18_log)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_18_log)(unsafe.Pointer(&__log_data.tab)))) + uintptr(i)*16))).logc
	z = *(*float64)(unsafe.Pointer(&_cgoz_28_log{iz}))
	r = (z - (*(*_cgoa_19_log)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_19_log)(unsafe.Pointer(&__log_data.tab2)))) + uintptr(i)*16))).chi - (*(*_cgoa_19_log)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_19_log)(unsafe.Pointer(&__log_data.tab2)))) + uintptr(i)*16))).clo) * invc
	kd = float64(k)
	w = kd*__log_data.ln2hi + logc
	hi = w + r
	lo = w - hi + r + kd*__log_data.ln2lo
	r2 = r * r
	y = lo + r2**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly)))) + uintptr(int32(0))*8)) + r*r2*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly)))) + uintptr(int32(1))*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly)))) + uintptr(int32(2))*8))+r2*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly)))) + uintptr(int32(3))*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__log_data.poly)))) + uintptr(int32(4))*8)))) + hi
	return eval_as_double(y)
}

type _cgoz_21_log struct {
	_f float64
}
type _cgoz_22_log struct {
	_f float64
}
type _cgoz_23_log struct {
	_f float64
}
type _cgoz_24_log struct {
	_f float64
}
type _cgoz_25_log struct {
	_f float64
}
type _cgoz_26_log struct {
	_f float64
}
type _cgoz_27_log struct {
	_f float64
}
type _cgoz_28_log struct {
	_i uint64
}
