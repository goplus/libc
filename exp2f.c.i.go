package libc

import unsafe "unsafe"

func _cgos_top12_exp2f(x float32) uint32 {
	return *(*uint32)(unsafe.Pointer(&_cgoz_18_exp2f{x})) >> int32(20)
}

type _cgoz_18_exp2f struct {
	_f float32
}

func Exp2f(x float32) float32 {
	var abstop uint32
	var ki uint64
	var t uint64
	var kd float64
	var xd float64
	var z float64
	var r float64
	var r2 float64
	var y float64
	var s float64
	xd = float64(x)
	abstop = _cgos_top12_exp2f(x) & uint32(2047)
	if func() int64 {
		if abstop >= _cgos_top12_exp2f(128) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if *(*uint32)(unsafe.Pointer(&_cgoz_19_exp2f{x})) == *(*uint32)(unsafe.Pointer(&_cgoz_20_exp2f{-X__builtin_inff()})) {
			return float32(0)
		}
		if abstop >= _cgos_top12_exp2f(X__builtin_inff()) {
			return x + x
		}
		if x > 0 {
			return __math_oflowf(uint32(0))
		}
		if x <= -150 {
			return __math_uflowf(uint32(0))
		}
	}
	kd = eval_as_double(xd + __exp2f_data.shift_scaled)
	ki = *(*uint64)(unsafe.Pointer(&_cgoz_21_exp2f{kd}))
	kd -= __exp2f_data.shift_scaled
	r = xd - kd
	t = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&__exp2f_data.tab)))) + uintptr(ki%uint64(32))*8))
	t += ki << 47
	s = *(*float64)(unsafe.Pointer(&_cgoz_22_exp2f{t}))
	z = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp2f_data.poly)))) + uintptr(int32(0))*8))*r + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp2f_data.poly)))) + uintptr(int32(1))*8))
	r2 = r * r
	y = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp2f_data.poly)))) + uintptr(int32(2))*8))*r + float64(int32(1))
	y = z*r2 + y
	y = y * s
	return eval_as_float(float32(y))
}

type _cgoz_19_exp2f struct {
	_f float32
}
type _cgoz_20_exp2f struct {
	_f float32
}
type _cgoz_21_exp2f struct {
	_f float64
}
type _cgoz_22_exp2f struct {
	_i uint64
}
