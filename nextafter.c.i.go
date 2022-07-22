package libc

import unsafe "unsafe"

func Nextafter(x float64, y float64) float64 {
	type _cgoa_18_nextafter struct {
		f float64
	}
	var ux _cgoa_18_nextafter
	ux.f = x
	var uy _cgoa_18_nextafter
	uy.f = y
	var ax uint64
	var ay uint64
	var e int32
	if func() int32 {
		if 8 == 4 {
			return func() int32 {
				if X__FLOAT_BITS(float32(x))&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 8 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(x)&9223372036854775807 > 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if X__fpclassifyl(float64(x)) == int32(0) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 || func() int32 {
		if 8 == 4 {
			return func() int32 {
				if X__FLOAT_BITS(float32(y))&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 8 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(y)&9223372036854775807 > 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if X__fpclassifyl(float64(y)) == int32(0) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 {
		return x + y
	}
	if *(*uint64)(unsafe.Pointer(&ux)) == *(*uint64)(unsafe.Pointer(&uy)) {
		return y
	}
	ax = *(*uint64)(unsafe.Pointer(&ux)) & 9223372036854775807
	ay = *(*uint64)(unsafe.Pointer(&uy)) & 9223372036854775807
	if ax == uint64(0) {
		if ay == uint64(0) {
			return y
		}
		*(*uint64)(unsafe.Pointer(&ux)) = *(*uint64)(unsafe.Pointer(&uy))&9223372036854775808 | uint64(1)
	} else if ax > ay || (*(*uint64)(unsafe.Pointer(&ux))^*(*uint64)(unsafe.Pointer(&uy)))&9223372036854775808 != 0 {
		*(*uint64)(unsafe.Pointer(&ux))--
	} else {
		*(*uint64)(unsafe.Pointer(&ux))++
	}
	e = int32(*(*uint64)(unsafe.Pointer(&ux)) >> int32(52) & uint64(2047))
	if e == int32(2047) {
		for {
			if 8 == 4 {
				fp_force_evalf(float32(x + x))
			} else if 8 == 8 {
				fp_force_eval(x + x)
			} else {
				fp_force_evall(float64(x + x))
			}
			if true {
				break
			}
		}
	}
	if e == int32(0) {
		for {
			if 8 == 4 {
				fp_force_evalf(float32(x*x + ux.f*ux.f))
			} else if 8 == 8 {
				fp_force_eval(x*x + ux.f*ux.f)
			} else {
				fp_force_evall(float64(x*x + ux.f*ux.f))
			}
			if true {
				break
			}
		}
	}
	return ux.f
}
