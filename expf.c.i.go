package libc

import unsafe "unsafe"

func top12_cgos__expf(x float32) uint32 {
	return *(*uint32)(unsafe.Pointer(&_cgoz_18_expf{x})) >> int32(20)
}

type _cgoz_18_expf struct {
	_f float32
}

func Expf(x float32) float32 {
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
	abstop = top12_cgos__expf(x) & uint32(2047)
	if func() int64 {
		if abstop >= top12_cgos__expf(88) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if *(*uint32)(unsafe.Pointer(&_cgoz_19_expf{x})) == *(*uint32)(unsafe.Pointer(&_cgoz_20_expf{-X__builtin_inff()})) {
			return float32(0)
		}
		if abstop >= top12_cgos__expf(X__builtin_inff()) {
			return x + x
		}
		if x > 88.7228317 {
			return __math_oflowf(uint32(0))
		}
		if x < -103.972076 {
			return __math_uflowf(uint32(0))
		}
	}
	z = __exp2f_data.invln2_scaled * xd
	kd = eval_as_double(z + __exp2f_data.shift)
	ki = *(*uint64)(unsafe.Pointer(&_cgoz_21_expf{kd}))
	kd -= __exp2f_data.shift
	r = z - kd
	t = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&__exp2f_data.tab)))) + uintptr(ki%uint64(32))*8))
	t += ki << 47
	s = *(*float64)(unsafe.Pointer(&_cgoz_22_expf{t}))
	z = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp2f_data.poly_scaled)))) + uintptr(int32(0))*8))*r + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp2f_data.poly_scaled)))) + uintptr(int32(1))*8))
	r2 = r * r
	y = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp2f_data.poly_scaled)))) + uintptr(int32(2))*8))*r + float64(int32(1))
	y = z*r2 + y
	y = y * s
	return eval_as_float(float32(y))
}

type _cgoz_19_expf struct {
	_f float32
}
type _cgoz_20_expf struct {
	_f float32
}
type _cgoz_21_expf struct {
	_f float64
}
type _cgoz_22_expf struct {
	_i uint64
}
