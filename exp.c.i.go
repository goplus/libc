package libc

import unsafe "unsafe"

func _cgos_specialcase__exp(tmp float64, sbits uint64, ki uint64) float64 {
	var scale float64
	var y float64
	if ki&uint64(2147483648) == uint64(0) {
		sbits -= 4544132024016830464
		scale = *(*float64)(unsafe.Pointer(&_cgoz_18_exp{sbits}))
		y = 5.4861240687936887e+303 * (scale + scale*tmp)
		return eval_as_double(y)
	}
	sbits += 4602678819172646912
	scale = *(*float64)(unsafe.Pointer(&_cgoz_19_exp{sbits}))
	y = scale + scale*tmp
	if y < 1 {
		var hi float64
		var lo float64
		lo = scale - y + scale*tmp
		hi = 1 + y
		lo = 1 - hi + y + lo
		y = eval_as_double(hi+lo) - 1
		if int32(1) != 0 && y == 0 {
			y = float64(0)
		}
		fp_force_eval(fp_barrier(2.2250738585072014e-308) * 2.2250738585072014e-308)
	}
	y = 2.2250738585072014e-308 * y
	return eval_as_double(y)
}

type _cgoz_18_exp struct {
	_i uint64
}
type _cgoz_19_exp struct {
	_i uint64
}

func _cgos_top12__exp(x float64) uint32 {
	return uint32(*(*uint64)(unsafe.Pointer(&_cgoz_20_exp{x})) >> int32(52))
}

type _cgoz_20_exp struct {
	_f float64
}

func Exp(x float64) float64 {
	var abstop uint32
	var ki uint64
	var idx uint64
	var top uint64
	var sbits uint64
	var kd float64
	var z float64
	var r float64
	var r2 float64
	var scale float64
	var tail float64
	var tmp float64
	abstop = _cgos_top12__exp(x) & uint32(2047)
	if func() int64 {
		if abstop-_cgos_top12__exp(5.5511151231257827e-17) >= _cgos_top12__exp(512)-_cgos_top12__exp(5.5511151231257827e-17) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if abstop-_cgos_top12__exp(5.5511151231257827e-17) >= uint32(2147483648) {
			return func() float64 {
				if int32(1) != 0 {
					return 1 + x
				} else {
					return 1
				}
			}()
		}
		if abstop >= _cgos_top12__exp(1024) {
			if *(*uint64)(unsafe.Pointer(&_cgoz_21_exp{x})) == *(*uint64)(unsafe.Pointer(&_cgoz_22_exp{float64(-X__builtin_inff())})) {
				return float64(0)
			}
			if abstop >= _cgos_top12__exp(float64(X__builtin_inff())) {
				return 1 + x
			}
			if *(*uint64)(unsafe.Pointer(&_cgoz_23_exp{x}))>>int32(63) != 0 {
				return __math_uflow(uint32(0))
			} else {
				return __math_oflow(uint32(0))
			}
		}
		abstop = uint32(0)
	}
	z = __exp_data.invln2N * x
	kd = eval_as_double(z + __exp_data.shift)
	ki = *(*uint64)(unsafe.Pointer(&_cgoz_24_exp{kd}))
	kd -= __exp_data.shift
	r = x + kd*__exp_data.negln2hiN + kd*__exp_data.negln2loN
	idx = uint64(2) * (ki % uint64(128))
	top = ki << 45
	tail = *(*float64)(unsafe.Pointer(&_cgoz_25_exp{*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&__exp_data.tab)))) + uintptr(idx)*8))}))
	sbits = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&__exp_data.tab)))) + uintptr(idx+uint64(1))*8)) + top
	r2 = r * r
	tmp = tail + r + r2*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp_data.poly)))) + uintptr(0)*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp_data.poly)))) + uintptr(1)*8))) + r2*r2*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp_data.poly)))) + uintptr(2)*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp_data.poly)))) + uintptr(3)*8)))
	if func() int64 {
		if abstop == uint32(0) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		return _cgos_specialcase__exp(tmp, sbits, ki)
	}
	scale = *(*float64)(unsafe.Pointer(&_cgoz_26_exp{sbits}))
	return eval_as_double(scale + scale*tmp)
}

type _cgoz_21_exp struct {
	_f float64
}
type _cgoz_22_exp struct {
	_f float64
}
type _cgoz_23_exp struct {
	_f float64
}
type _cgoz_24_exp struct {
	_f float64
}
type _cgoz_25_exp struct {
	_i uint64
}
type _cgoz_26_exp struct {
	_i uint64
}
