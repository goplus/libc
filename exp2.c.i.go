package libc

import unsafe "unsafe"

func specialcase_cgo18_exp2(tmp float64, sbits uint64, ki uint64) float64 {
	var scale float64
	var y float64
	if ki&uint64(2147483648) == uint64(0) {
		sbits -= 4503599627370496
		scale = *(*float64)(unsafe.Pointer(&_cgoz_19_exp2{sbits}))
		y = float64(int32(2)) * (scale + scale*tmp)
		return eval_as_double(y)
	}
	sbits += 4602678819172646912
	scale = *(*float64)(unsafe.Pointer(&_cgoz_20_exp2{sbits}))
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

type _cgoz_19_exp2 struct {
	_i uint64
}
type _cgoz_20_exp2 struct {
	_i uint64
}

func top12_cgo21_exp2(x float64) uint32 {
	return uint32(*(*uint64)(unsafe.Pointer(&_cgoz_22_exp2{x})) >> int32(52))
}

type _cgoz_22_exp2 struct {
	_f float64
}

func Exp2(x float64) float64 {
	var abstop uint32
	var ki uint64
	var idx uint64
	var top uint64
	var sbits uint64
	var kd float64
	var r float64
	var r2 float64
	var scale float64
	var tail float64
	var tmp float64
	abstop = top12_cgo21_exp2(x) & uint32(2047)
	if func() int64 {
		if abstop-top12_cgo21_exp2(5.5511151231257827e-17) >= top12_cgo21_exp2(512)-top12_cgo21_exp2(5.5511151231257827e-17) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if abstop-top12_cgo21_exp2(5.5511151231257827e-17) >= uint32(2147483648) {
			return func() float64 {
				if int32(1) != 0 {
					return 1 + x
				} else {
					return 1
				}
			}()
		}
		if abstop >= top12_cgo21_exp2(1024) {
			if *(*uint64)(unsafe.Pointer(&_cgoz_23_exp2{x})) == *(*uint64)(unsafe.Pointer(&_cgoz_24_exp2{float64(-X__builtin_inff())})) {
				return float64(0)
			}
			if abstop >= top12_cgo21_exp2(float64(X__builtin_inff())) {
				return 1 + x
			}
			if !(*(*uint64)(unsafe.Pointer(&_cgoz_25_exp2{x}))>>int32(63) != 0) {
				return __math_oflow(uint32(0))
			} else if *(*uint64)(unsafe.Pointer(&_cgoz_26_exp2{x})) >= *(*uint64)(unsafe.Pointer(&_cgoz_27_exp2{-1075})) {
				return __math_uflow(uint32(0))
			}
		}
		if uint64(2)**(*uint64)(unsafe.Pointer(&_cgoz_28_exp2{x})) > uint64(2)**(*uint64)(unsafe.Pointer(&_cgoz_29_exp2{928})) {
			abstop = uint32(0)
		}
	}
	kd = eval_as_double(x + __exp_data.exp2_shift)
	ki = *(*uint64)(unsafe.Pointer(&_cgoz_30_exp2{kd}))
	kd -= __exp_data.exp2_shift
	r = x - kd
	idx = uint64(2) * (ki % uint64(128))
	top = ki << 45
	tail = *(*float64)(unsafe.Pointer(&_cgoz_31_exp2{*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&__exp_data.tab)))) + uintptr(idx)*8))}))
	sbits = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&__exp_data.tab)))) + uintptr(idx+uint64(1))*8)) + top
	r2 = r * r
	tmp = tail + r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp_data.exp2_poly)))) + uintptr(int32(0))*8)) + r2*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp_data.exp2_poly)))) + uintptr(int32(1))*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp_data.exp2_poly)))) + uintptr(int32(2))*8))) + r2*r2*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp_data.exp2_poly)))) + uintptr(int32(3))*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp_data.exp2_poly)))) + uintptr(int32(4))*8)))
	if func() int64 {
		if abstop == uint32(0) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		return specialcase_cgo18_exp2(tmp, sbits, ki)
	}
	scale = *(*float64)(unsafe.Pointer(&_cgoz_32_exp2{sbits}))
	return eval_as_double(scale + scale*tmp)
}

type _cgoz_23_exp2 struct {
	_f float64
}
type _cgoz_24_exp2 struct {
	_f float64
}
type _cgoz_25_exp2 struct {
	_f float64
}
type _cgoz_26_exp2 struct {
	_f float64
}
type _cgoz_27_exp2 struct {
	_f float64
}
type _cgoz_28_exp2 struct {
	_f float64
}
type _cgoz_29_exp2 struct {
	_f float64
}
type _cgoz_30_exp2 struct {
	_f float64
}
type _cgoz_31_exp2 struct {
	_i uint64
}
type _cgoz_32_exp2 struct {
	_i uint64
}
