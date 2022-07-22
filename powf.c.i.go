package libc

import unsafe "unsafe"

func log2_inline_cgo19_powf(ix uint32) float64 {
	var z float64
	var r float64
	var r2 float64
	var r4 float64
	var p float64
	var q float64
	var y float64
	var y0 float64
	var invc float64
	var logc float64
	var iz uint32
	var top uint32
	var tmp uint32
	var k int32
	var i int32
	tmp = ix - uint32(1060306944)
	i = int32(tmp >> 19 % uint32(16))
	top = tmp & uint32(4286578688)
	iz = ix - top
	k = int32(top) >> 23
	invc = (*(*_cgoa_18_powf)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_18_powf)(unsafe.Pointer(&__powf_log2_data.tab)))) + uintptr(i)*16))).invc
	logc = (*(*_cgoa_18_powf)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_18_powf)(unsafe.Pointer(&__powf_log2_data.tab)))) + uintptr(i)*16))).logc
	z = float64(*(*float32)(unsafe.Pointer(&_cgoz_20_powf{iz})))
	r = z*invc - float64(int32(1))
	y0 = logc + float64(k)
	r2 = r * r
	y = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__powf_log2_data.poly)))) + uintptr(int32(0))*8))*r + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__powf_log2_data.poly)))) + uintptr(int32(1))*8))
	p = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__powf_log2_data.poly)))) + uintptr(int32(2))*8))*r + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__powf_log2_data.poly)))) + uintptr(int32(3))*8))
	r4 = r2 * r2
	q = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__powf_log2_data.poly)))) + uintptr(int32(4))*8))*r + y0
	q = p*r2 + q
	y = y*r4 + q
	return y
}

type _cgoz_20_powf struct {
	_i uint32
}

func exp2_inline_cgo21_powf(xd float64, sign_bias uint32) float32 {
	var ki uint64
	var ski uint64
	var t uint64
	var kd float64
	var z float64
	var r float64
	var r2 float64
	var y float64
	var s float64
	kd = eval_as_double(xd + __exp2f_data.shift_scaled)
	ki = *(*uint64)(unsafe.Pointer(&_cgoz_22_powf{kd}))
	kd -= __exp2f_data.shift_scaled
	r = xd - kd
	t = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&__exp2f_data.tab)))) + uintptr(ki%uint64(32))*8))
	ski = ki + uint64(sign_bias)
	t += ski << 47
	s = *(*float64)(unsafe.Pointer(&_cgoz_23_powf{t}))
	z = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp2f_data.poly)))) + uintptr(int32(0))*8))*r + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp2f_data.poly)))) + uintptr(int32(1))*8))
	r2 = r * r
	y = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__exp2f_data.poly)))) + uintptr(int32(2))*8))*r + float64(int32(1))
	y = z*r2 + y
	y = y * s
	return eval_as_float(float32(y))
}

type _cgoz_22_powf struct {
	_f float64
}
type _cgoz_23_powf struct {
	_i uint64
}

func checkint_cgo24_powf(iy uint32) int32 {
	var e int32 = int32(iy >> int32(23) & uint32(255))
	if e < int32(127) {
		return int32(0)
	}
	if e > 150 {
		return int32(2)
	}
	if iy&uint32(int32(1)<<(150-e)-int32(1)) != 0 {
		return int32(0)
	}
	if iy&uint32(int32(1)<<(150-e)) != 0 {
		return int32(1)
	}
	return int32(2)
}
func zeroinfnan_cgo25_powf(ix uint32) int32 {
	return func() int32 {
		if uint32(2)*ix-uint32(1) >= 4278190079 {
			return 1
		} else {
			return 0
		}
	}()
}
func Powf(x float32, y float32) float32 {
	var sign_bias uint32 = uint32(0)
	var ix uint32
	var iy uint32
	ix = *(*uint32)(unsafe.Pointer(&_cgoz_26_powf{x}))
	iy = *(*uint32)(unsafe.Pointer(&_cgoz_27_powf{y}))
	if func() int64 {
		if ix-uint32(8388608) >= uint32(2130706432) || zeroinfnan_cgo25_powf(iy) != 0 {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if int64(zeroinfnan_cgo25_powf(iy)) == int64(0) {
			if uint32(2)*iy == uint32(0) {
				return func() float32 {
					if int32(0) != 0 {
						return x + y
					} else {
						return 1
					}
				}()
			}
			if ix == uint32(1065353216) {
				return func() float32 {
					if int32(0) != 0 {
						return x + y
					} else {
						return 1
					}
				}()
			}
			if uint32(2)*ix > 4278190080 || uint32(2)*iy > 4278190080 {
				return x + y
			}
			if uint32(2)*ix == uint32(2130706432) {
				return float32(1)
			}
			if func() int32 {
				if uint32(2)*ix < uint32(2130706432) {
					return 1
				} else {
					return 0
				}
			}() == func() int32 {
				if !(iy&uint32(2147483648) != 0) {
					return 1
				} else {
					return 0
				}
			}() {
				return float32(0)
			}
			return y * y
		}
		if int64(zeroinfnan_cgo25_powf(ix)) == int64(0) {
			var x2 float32 = x * x
			if ix&uint32(2147483648) != 0 && checkint_cgo24_powf(iy) == int32(1) {
				x2 = -x2
			}
			return func() float32 {
				if iy&uint32(2147483648) != 0 {
					return fp_barrierf(float32(int32(1)) / x2)
				} else {
					return x2
				}
			}()
		}
		if ix&uint32(2147483648) != 0 {
			var yint int32 = checkint_cgo24_powf(iy)
			if yint == int32(0) {
				return __math_invalidf(x)
			}
			if yint == int32(1) {
				sign_bias = uint32(65536)
			}
			ix &= uint32(2147483647)
		}
		if ix < uint32(8388608) {
			ix = *(*uint32)(unsafe.Pointer(&_cgoz_28_powf{x * 8388608}))
			ix &= uint32(2147483647)
			ix -= uint32(192937984)
		}
	}
	var logx float64 = log2_inline_cgo19_powf(ix)
	var ylogx float64 = float64(y) * logx
	if func() int64 {
		if *(*uint64)(unsafe.Pointer(&_cgoz_29_powf{ylogx}))>>int32(47)&uint64(65535) >= *(*uint64)(unsafe.Pointer(&_cgoz_30_powf{126 * float64(1)}))>>int32(47) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if ylogx > 127.99999995700433*float64(1) {
			return __math_oflowf(sign_bias)
		}
		if ylogx <= -150*float64(1) {
			return __math_uflowf(sign_bias)
		}
	}
	return exp2_inline_cgo21_powf(ylogx, sign_bias)
}

type _cgoz_26_powf struct {
	_f float32
}
type _cgoz_27_powf struct {
	_f float32
}
type _cgoz_28_powf struct {
	_f float32
}
type _cgoz_29_powf struct {
	_f float64
}
type _cgoz_30_powf struct {
	_f float64
}
