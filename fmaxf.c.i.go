package libc

func Fmaxf(x float32, y float32) float32 {
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
	}() != 0 {
		return y
	}
	if func() int32 {
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
		return x
	}
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
	}() != func() int32 {
		if 4 == 4 {
			return int32(X__FLOAT_BITS(y) >> int32(31))
		} else {
			return func() int32 {
				if 4 == 8 {
					return int32(X__DOUBLE_BITS(float64(y)) >> int32(63))
				} else {
					return X__signbitl(float64(y))
				}
			}()
		}
	}() {
		return func() float32 {
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
				return y
			} else {
				return x
			}
		}()
	}
	return func() float32 {
		if x < y {
			return y
		} else {
			return x
		}
	}()
}
