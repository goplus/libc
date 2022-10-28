package libc

func Scalb(x float64, fn float64) float64 {
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
				if X__FLOAT_BITS(float32(fn))&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 8 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(fn)&9223372036854775807 > 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if X__fpclassifyl(float64(fn)) == int32(0) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 {
		return x * fn
	}
	if !(func() int32 {
		if 8 == 4 {
			return func() int32 {
				if X__FLOAT_BITS(float32(fn))&uint32(2147483647) < uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 8 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(fn)&9223372036854775807 < 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if X__fpclassifyl(float64(fn)) > int32(1) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0) {
		if fn > 0.0 {
			return x * fn
		} else {
			return x / -fn
		}
	}
	if Rint(fn) != fn {
		return (fn - fn) / (fn - fn)
	}
	if fn > 65000.0 {
		return Scalbn(x, int32(65000))
	}
	if -fn > 65000.0 {
		return Scalbn(x, -65000)
	}
	return Scalbn(x, int32(fn))
}
