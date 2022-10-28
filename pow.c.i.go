package libc

import unsafe "unsafe"

func _cgos_top12_pow(x float64) uint32 {
	return uint32(*(*uint64)(unsafe.Pointer(&_cgoz_19_pow{x})) >> int32(52))
}

type _cgoz_19_pow struct {
	_f float64
}

func _cgos_log_inline_pow(ix uint64, tail *float64) float64 {
	var z float64
	var r float64
	var y float64
	var invc float64
	var logc float64
	var logctail float64
	var kd float64
	var hi float64
	var t1 float64
	var t2 float64
	var lo float64
	var lo1 float64
	var lo2 float64
	var p float64
	var iz uint64
	var tmp uint64
	var k int32
	var i int32
	tmp = ix - uint64(4604531861337669632)
	i = int32(tmp >> 45 % uint64(128))
	k = int32(int64(tmp) >> int32(52))
	iz = ix - tmp&18442240474082181120
	z = *(*float64)(unsafe.Pointer(&_cgoz_20_pow{iz}))
	kd = float64(k)
	invc = (*(*_cgoa_18_pow)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_18_pow)(unsafe.Pointer(&__pow_log_data.tab)))) + uintptr(i)*32))).invc
	logc = (*(*_cgoa_18_pow)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_18_pow)(unsafe.Pointer(&__pow_log_data.tab)))) + uintptr(i)*32))).logc
	logctail = (*(*_cgoa_18_pow)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_18_pow)(unsafe.Pointer(&__pow_log_data.tab)))) + uintptr(i)*32))).logctail
	var zhi float64 = *(*float64)(unsafe.Pointer(&_cgoz_21_pow{(iz + 2147483648) & 18446744069414584320}))
	var zlo float64 = z - zhi
	var rhi float64 = zhi*invc - 1.0
	var rlo float64 = zlo * invc
	r = rhi + rlo
	t1 = kd*__pow_log_data.ln2hi + logc
	t2 = t1 + r
	lo1 = kd*__pow_log_data.ln2lo + logctail
	lo2 = t1 - t2 + r
	var ar float64
	var ar2 float64
	var ar3 float64
	var lo3 float64
	var lo4 float64
	ar = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__pow_log_data.poly)))) + uintptr(int32(0))*8)) * r
	ar2 = r * ar
	ar3 = r * ar2
	var arhi float64 = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__pow_log_data.poly)))) + uintptr(int32(0))*8)) * rhi
	var arhi2 float64 = rhi * arhi
	hi = t2 + arhi2
	lo3 = rlo * (ar + arhi)
	lo4 = t2 - hi + arhi2
	p = ar3 * (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__pow_log_data.poly)))) + uintptr(int32(1))*8)) + r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__pow_log_data.poly)))) + uintptr(int32(2))*8)) + ar2*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__pow_log_data.poly)))) + uintptr(int32(3))*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__pow_log_data.poly)))) + uintptr(int32(4))*8))+ar2*(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__pow_log_data.poly)))) + uintptr(int32(5))*8))+r**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&__pow_log_data.poly)))) + uintptr(int32(6))*8)))))
	lo = lo1 + lo2 + lo3 + lo4 + p
	y = hi + lo
	*tail = hi - y + lo
	return y
}

type _cgoz_20_pow struct {
	_i uint64
}
type _cgoz_21_pow struct {
	_i uint64
}

func _cgos_specialcase_pow(tmp float64, sbits uint64, ki uint64) float64 {
	var scale float64
	var y float64
	if ki&uint64(2147483648) == uint64(0) {
		sbits -= 4544132024016830464
		scale = *(*float64)(unsafe.Pointer(&_cgoz_22_pow{sbits}))
		y = 5.4861240687936887e+303 * (scale + scale*tmp)
		return eval_as_double(y)
	}
	sbits += 4602678819172646912
	scale = *(*float64)(unsafe.Pointer(&_cgoz_23_pow{sbits}))
	y = scale + scale*tmp
	if Fabs(y) < 1.0 {
		var hi float64
		var lo float64
		var one float64 = 1.0
		if y < 0.0 {
			one = float64(-1.0)
		}
		lo = scale - y + scale*tmp
		hi = one + y
		lo = one - hi + y + lo
		y = eval_as_double(hi+lo) - one
		if y == 0.0 {
			y = *(*float64)(unsafe.Pointer(&_cgoz_24_pow{sbits & uint64(9223372036854775808)}))
		}
		fp_force_eval(fp_barrier(2.2250738585072014e-308) * 2.2250738585072014e-308)
	}
	y = 2.2250738585072014e-308 * y
	return eval_as_double(y)
}

type _cgoz_22_pow struct {
	_i uint64
}
type _cgoz_23_pow struct {
	_i uint64
}
type _cgoz_24_pow struct {
	_i uint64
}

func _cgos_exp_inline_pow(x float64, xtail float64, sign_bias uint32) float64 {
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
	abstop = _cgos_top12_pow(x) & uint32(2047)
	if func() int64 {
		if abstop-_cgos_top12_pow(5.5511151231257827e-17) >= _cgos_top12_pow(512.0)-_cgos_top12_pow(5.5511151231257827e-17) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if abstop-_cgos_top12_pow(5.5511151231257827e-17) >= uint32(2147483648) {
			var one float64 = func() float64 {
				if int32(1) != 0 {
					return 1.0 + x
				} else {
					return 1.0
				}
			}()
			return func() float64 {
				if sign_bias != 0 {
					return -one
				} else {
					return one
				}
			}()
		}
		if abstop >= _cgos_top12_pow(1024.0) {
			if *(*uint64)(unsafe.Pointer(&_cgoz_25_pow{x}))>>int32(63) != 0 {
				return __math_uflow(sign_bias)
			} else {
				return __math_oflow(sign_bias)
			}
		}
		abstop = uint32(0)
	}
	z = __exp_data.invln2N * x
	kd = eval_as_double(z + __exp_data.shift)
	ki = *(*uint64)(unsafe.Pointer(&_cgoz_26_pow{kd}))
	kd -= __exp_data.shift
	r = x + kd*__exp_data.negln2hiN + kd*__exp_data.negln2loN
	r += xtail
	idx = uint64(2) * (ki % uint64(128))
	top = (ki + uint64(sign_bias)) << 45
	tail = *(*float64)(unsafe.Pointer(&_cgoz_27_pow{*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(&__exp_data.tab)))) + uintptr(idx)*8))}))
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
		return _cgos_specialcase_pow(tmp, sbits, ki)
	}
	scale = *(*float64)(unsafe.Pointer(&_cgoz_28_pow{sbits}))
	return eval_as_double(scale + scale*tmp)
}

type _cgoz_25_pow struct {
	_f float64
}
type _cgoz_26_pow struct {
	_f float64
}
type _cgoz_27_pow struct {
	_i uint64
}
type _cgoz_28_pow struct {
	_i uint64
}

func _cgos_checkint_pow(iy uint64) int32 {
	var e int32 = int32(iy >> int32(52) & uint64(2047))
	if e < int32(1023) {
		return int32(0)
	}
	if e > 1075 {
		return int32(2)
	}
	if iy&(uint64(1)<<(1075-e)-uint64(1)) != 0 {
		return int32(0)
	}
	if iy&(uint64(1)<<(1075-e)) != 0 {
		return int32(1)
	}
	return int32(2)
}
func _cgos_zeroinfnan_pow(i uint64) int32 {
	return func() int32 {
		if uint64(2)*i-uint64(1) >= uint64(2)**(*uint64)(unsafe.Pointer(&_cgoz_29_pow{float64(X__builtin_inff())}))-uint64(1) {
			return 1
		} else {
			return 0
		}
	}()
}

type _cgoz_29_pow struct {
	_f float64
}

func Pow(x float64, y float64) float64 {
	var sign_bias uint32 = uint32(0)
	var ix uint64
	var iy uint64
	var topx uint32
	var topy uint32
	ix = *(*uint64)(unsafe.Pointer(&_cgoz_30_pow{x}))
	iy = *(*uint64)(unsafe.Pointer(&_cgoz_31_pow{y}))
	topx = _cgos_top12_pow(x)
	topy = _cgos_top12_pow(y)
	if func() int64 {
		if topx-uint32(1) >= uint32(2046) || topy&uint32(2047)-uint32(958) >= uint32(128) {
			return 1
		} else {
			return 0
		}
	}() == int64(0) {
		if int64(_cgos_zeroinfnan_pow(iy)) == int64(0) {
			if uint64(2)*iy == uint64(0) {
				return func() float64 {
					if int32(0) != 0 {
						return x + y
					} else {
						return 1.0
					}
				}()
			}
			if ix == *(*uint64)(unsafe.Pointer(&_cgoz_32_pow{1.0})) {
				return func() float64 {
					if int32(0) != 0 {
						return x + y
					} else {
						return 1.0
					}
				}()
			}
			if uint64(2)*ix > uint64(2)**(*uint64)(unsafe.Pointer(&_cgoz_33_pow{float64(X__builtin_inff())})) || uint64(2)*iy > uint64(2)**(*uint64)(unsafe.Pointer(&_cgoz_34_pow{float64(X__builtin_inff())})) {
				return x + y
			}
			if uint64(2)*ix == uint64(2)**(*uint64)(unsafe.Pointer(&_cgoz_35_pow{1.0})) {
				return float64(1.0)
			}
			if func() int32 {
				if uint64(2)*ix < uint64(2)**(*uint64)(unsafe.Pointer(&_cgoz_36_pow{1.0})) {
					return 1
				} else {
					return 0
				}
			}() == func() int32 {
				if !(iy>>int32(63) != 0) {
					return 1
				} else {
					return 0
				}
			}() {
				return float64(0.0)
			}
			return y * y
		}
		if int64(_cgos_zeroinfnan_pow(ix)) == int64(0) {
			var x2 float64 = x * x
			if ix>>int32(63) != 0 && _cgos_checkint_pow(iy) == int32(1) {
				x2 = -x2
			}
			return func() float64 {
				if iy>>int32(63) != 0 {
					return fp_barrier(float64(int32(1)) / x2)
				} else {
					return x2
				}
			}()
		}
		if ix>>int32(63) != 0 {
			var yint int32 = _cgos_checkint_pow(iy)
			if yint == int32(0) {
				return __math_invalid(x)
			}
			if yint == int32(1) {
				sign_bias = uint32(262144)
			}
			ix &= uint64(9223372036854775807)
			topx &= uint32(2047)
		}
		if topy&uint32(2047)-uint32(958) >= uint32(128) {
			if ix == *(*uint64)(unsafe.Pointer(&_cgoz_37_pow{1.0})) {
				return float64(1.0)
			}
			if topy&uint32(2047) < uint32(958) {
				if int32(1) != 0 {
					return func() float64 {
						if ix > *(*uint64)(unsafe.Pointer(&_cgoz_38_pow{1.0})) {
							return 1.0 + y
						} else {
							return 1.0 - y
						}
					}()
				} else {
					return float64(1.0)
				}
			}
			return func() float64 {
				if func() int32 {
					if ix > *(*uint64)(unsafe.Pointer(&_cgoz_39_pow{1.0})) {
						return 1
					} else {
						return 0
					}
				}() == func() int32 {
					if topy < uint32(2048) {
						return 1
					} else {
						return 0
					}
				}() {
					return __math_oflow(uint32(0))
				} else {
					return __math_uflow(uint32(0))
				}
			}()
		}
		if topx == uint32(0) {
			ix = *(*uint64)(unsafe.Pointer(&_cgoz_40_pow{x * 4503599627370496.0}))
			ix &= uint64(9223372036854775807)
			ix -= 234187180623265792
		}
	}
	var lo float64
	var hi float64 = _cgos_log_inline_pow(ix, &lo)
	var ehi float64
	var elo float64
	var yhi float64 = *(*float64)(unsafe.Pointer(&_cgoz_41_pow{iy & 18446744073575333888}))
	var ylo float64 = y - yhi
	var lhi float64 = *(*float64)(unsafe.Pointer(&_cgoz_42_pow{*(*uint64)(unsafe.Pointer(&_cgoz_43_pow{hi})) & 18446744073575333888}))
	var llo float64 = hi - lhi + lo
	ehi = yhi * lhi
	elo = ylo*lhi + y*llo
	return _cgos_exp_inline_pow(ehi, elo, sign_bias)
}

type _cgoz_30_pow struct {
	_f float64
}
type _cgoz_31_pow struct {
	_f float64
}
type _cgoz_32_pow struct {
	_f float64
}
type _cgoz_33_pow struct {
	_f float64
}
type _cgoz_34_pow struct {
	_f float64
}
type _cgoz_35_pow struct {
	_f float64
}
type _cgoz_36_pow struct {
	_f float64
}
type _cgoz_37_pow struct {
	_f float64
}
type _cgoz_38_pow struct {
	_f float64
}
type _cgoz_39_pow struct {
	_f float64
}
type _cgoz_40_pow struct {
	_f float64
}
type _cgoz_41_pow struct {
	_i uint64
}
type _cgoz_42_pow struct {
	_i uint64
}
type _cgoz_43_pow struct {
	_f float64
}
