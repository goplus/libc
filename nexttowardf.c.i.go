package libc

import unsafe "unsafe"

func Nexttowardf(x float32, y float64) float32 {
	type _cgoa_18_nexttowardf struct {
		f float32
	}
	var ux _cgoa_18_nexttowardf
	ux.f = x
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
						if X__DOUBLE_BITS(float64(y))&9223372036854775807 > 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if X__fpclassifyl(y) == int32(0) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 {
		return float32(float64(x) + y)
	}
	if float64(x) == y {
		return float32(y)
	}
	if x == float32(int32(0)) {
		*(*uint32)(unsafe.Pointer(&ux)) = uint32(1)
		if func() int32 {
			if 8 == 4 {
				return int32(X__FLOAT_BITS(float32(y)) >> int32(31))
			} else {
				return func() int32 {
					if 8 == 8 {
						return int32(X__DOUBLE_BITS(float64(y)) >> int32(63))
					} else {
						return X__signbitl(y)
					}
				}()
			}
		}() != 0 {
			*(*uint32)(unsafe.Pointer(&ux)) |= uint32(2147483648)
		}
	} else if float64(x) < y {
		if func() int32 {
			if 4 == 4 {
				return int32(X__FLOAT_BITS(x) >> int32(31))
			} else {
				return func() int32 {
					if 4 == 8 {
						return int32(X__DOUBLE_BITS(float64(x)) >> int32(63))
					} else {
						return X__signbitl(float64(x))
					}
				}()
			}
		}() != 0 {
			*(*uint32)(unsafe.Pointer(&ux))--
		} else {
			*(*uint32)(unsafe.Pointer(&ux))++
		}
	} else if func() int32 {
		if 4 == 4 {
			return int32(X__FLOAT_BITS(x) >> int32(31))
		} else {
			return func() int32 {
				if 4 == 8 {
					return int32(X__DOUBLE_BITS(float64(x)) >> int32(63))
				} else {
					return X__signbitl(float64(x))
				}
			}()
		}
	}() != 0 {
		*(*uint32)(unsafe.Pointer(&ux))++
	} else {
		*(*uint32)(unsafe.Pointer(&ux))--
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
