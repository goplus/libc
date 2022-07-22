package libc

import unsafe "unsafe"

func Nextafterf(x float32, y float32) float32 {
	type _cgoa_18_nextafterf struct {
		f float32
	}
	var ux _cgoa_18_nextafterf
	ux.f = x
	var uy _cgoa_18_nextafterf
	uy.f = y
	var ax uint32
	var ay uint32
	var e uint32
	if func() int32 {
		if 4 == 4 {
			return func() int32 {
				if X__FLOAT_BITS(x)&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 4 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(float64(x))&9223372036854775807 > 9218868437227405312 {
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
		if 4 == 4 {
			return func() int32 {
				if X__FLOAT_BITS(y)&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 4 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(float64(y))&9223372036854775807 > 9218868437227405312 {
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
	if *(*uint32)(unsafe.Pointer(&ux)) == *(*uint32)(unsafe.Pointer(&uy)) {
		return y
	}
	ax = *(*uint32)(unsafe.Pointer(&ux)) & uint32(2147483647)
	ay = *(*uint32)(unsafe.Pointer(&uy)) & uint32(2147483647)
	if ax == uint32(0) {
		if ay == uint32(0) {
			return y
		}
		*(*uint32)(unsafe.Pointer(&ux)) = *(*uint32)(unsafe.Pointer(&uy))&uint32(2147483648) | uint32(1)
	} else if ax > ay || (*(*uint32)(unsafe.Pointer(&ux))^*(*uint32)(unsafe.Pointer(&uy)))&uint32(2147483648) != 0 {
		*(*uint32)(unsafe.Pointer(&ux))--
	} else {
		*(*uint32)(unsafe.Pointer(&ux))++
	}
	e = *(*uint32)(unsafe.Pointer(&ux)) & uint32(2139095040)
	if e == uint32(2139095040) {
		for {
			if 4 == 4 {
				fp_force_evalf(x + x)
			} else if 4 == 8 {
				fp_force_eval(float64(x + x))
			} else {
				fp_force_evall(float64(x + x))
			}
			if true {
				break
			}
		}
	}
	if e == uint32(0) {
		for {
			if 4 == 4 {
				fp_force_evalf(x*x + ux.f*ux.f)
			} else if 4 == 8 {
				fp_force_eval(float64(x*x + ux.f*ux.f))
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
