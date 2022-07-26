package common

import libc "github.com/goplus/libc"

type Struct_d_d struct {
	File *int8
	Line int32
	R    int32
	X    float64
	Y    float64
	Dy   float32
	E    int32
}
type Struct_f_f struct {
	File *int8
	Line int32
	R    int32
	X    float32
	Y    float32
	Dy   float32
	E    int32
}
type Struct_l_l struct {
	File *int8
	Line int32
	R    int32
	X    float64
	Y    float64
	Dy   float32
	E    int32
}
type Struct_ff_f struct {
	File *int8
	Line int32
	R    int32
	X    float32
	X2   float32
	Y    float32
	Dy   float32
	E    int32
}
type Struct_dd_d struct {
	File *int8
	Line int32
	R    int32
	X    float64
	X2   float64
	Y    float64
	Dy   float32
	E    int32
}
type Struct_ll_l struct {
	File *int8
	Line int32
	R    int32
	X    float64
	X2   float64
	Y    float64
	Dy   float32
	E    int32
}
type Struct_d_di struct {
	File *int8
	Line int32
	R    int32
	X    float64
	Y    float64
	Dy   float32
	I    int64
	E    int32
}
type Struct_f_fi struct {
	File *int8
	Line int32
	R    int32
	X    float32
	Y    float32
	Dy   float32
	I    int64
	E    int32
}
type Struct_l_li struct {
	File *int8
	Line int32
	R    int32
	X    float64
	Y    float64
	Dy   float32
	I    int64
	E    int32
}
type Struct_di_d struct {
	File *int8
	Line int32
	R    int32
	X    float64
	I    int64
	Y    float64
	Dy   float32
	E    int32
}
type Struct_fi_f struct {
	File *int8
	Line int32
	R    int32
	X    float32
	I    int64
	Y    float32
	Dy   float32
	E    int32
}
type Struct_li_l struct {
	File *int8
	Line int32
	R    int32
	X    float64
	I    int64
	Y    float64
	Dy   float32
	E    int32
}
type Struct_d_i struct {
	File *int8
	Line int32
	R    int32
	X    float64
	I    int64
	E    int32
}
type Struct_f_i struct {
	File *int8
	Line int32
	R    int32
	X    float32
	I    int64
	E    int32
}
type Struct_l_i struct {
	File *int8
	Line int32
	R    int32
	X    float64
	I    int64
	E    int32
}
type Struct_d_dd struct {
	File *int8
	Line int32
	R    int32
	X    float64
	Y    float64
	Dy   float32
	Y2   float64
	Dy2  float32
	E    int32
}
type Struct_f_ff struct {
	File *int8
	Line int32
	R    int32
	X    float32
	Y    float32
	Dy   float32
	Y2   float32
	Dy2  float32
	E    int32
}
type Struct_l_ll struct {
	File *int8
	Line int32
	R    int32
	X    float64
	Y    float64
	Dy   float32
	Y2   float64
	Dy2  float32
	E    int32
}
type Struct_ff_fi struct {
	File *int8
	Line int32
	R    int32
	X    float32
	X2   float32
	Y    float32
	Dy   float32
	I    int64
	E    int32
}
type Struct_dd_di struct {
	File *int8
	Line int32
	R    int32
	X    float64
	X2   float64
	Y    float64
	Dy   float32
	I    int64
	E    int32
}
type Struct_ll_li struct {
	File *int8
	Line int32
	R    int32
	X    float64
	X2   float64
	Y    float64
	Dy   float32
	I    int64
	E    int32
}
type Struct_fff_f struct {
	File *int8
	Line int32
	R    int32
	X    float32
	X2   float32
	X3   float32
	Y    float32
	Dy   float32
	E    int32
}
type Struct_ddd_d struct {
	File *int8
	Line int32
	R    int32
	X    float64
	X2   float64
	X3   float64
	Y    float64
	Dy   float32
	E    int32
}
type Struct_lll_l struct {
	File *int8
	Line int32
	R    int32
	X    float64
	X2   float64
	X3   float64
	Y    float64
	Dy   float32
	E    int32
}

func Checkexcept(got int32, want int32, r int32) int32 {
	if r == int32(0) {
		return func() int32 {
			if got|int32(0) == want|int32(0) {
				return 1
			} else {
				return 0
			}
		}()
	}
	return func() int32 {
		if got|int32(0)|int32(0) == want|int32(0)|int32(0) {
			return 1
		} else {
			return 0
		}
	}()
}
func Checkexceptall(got int32, want int32, r int32) int32 {
	return func() int32 {
		if got == want {
			return 1
		} else {
			return 0
		}
	}()
}
func Checkulp(d float32, r int32) int32 {
	if r == int32(0) {
		return func() int32 {
			if float64(libc.Fabsf(d)) < 1.5 {
				return 1
			} else {
				return 0
			}
		}()
	}
	return func() int32 {
		if float64(libc.Fabsf(d)) < 3 {
			return 1
		} else {
			return 0
		}
	}()
}
func Checkcr(y float64, ywant float64, r int32) int32 {
	if func() int32 {
		if 8 == 4 {
			return func() int32 {
				if libc.X__FLOAT_BITS(float32(ywant))&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 8 == 8 {
					return func() int32 {
						if libc.X__DOUBLE_BITS(float64(ywant))&9223372036854775807 > 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if libc.X__fpclassifyl(ywant) == int32(0) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 {
		return func() int32 {
			if 8 == 4 {
				return func() int32 {
					if libc.X__FLOAT_BITS(float32(y))&uint32(2147483647) > uint32(2139095040) {
						return 1
					} else {
						return 0
					}
				}()
			} else {
				return func() int32 {
					if 8 == 8 {
						return func() int32 {
							if libc.X__DOUBLE_BITS(float64(y))&9223372036854775807 > 9218868437227405312 {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if libc.X__fpclassifyl(y) == int32(0) {
								return 1
							} else {
								return 0
							}
						}()
					}
				}()
			}
		}()
	}
	return func() int32 {
		if y == ywant && func() int32 {
			if 8 == 4 {
				return int32(libc.X__FLOAT_BITS(float32(y)) >> int32(31))
			} else {
				return func() int32 {
					if 8 == 8 {
						return int32(libc.X__DOUBLE_BITS(float64(y)) >> int32(63))
					} else {
						return libc.X__signbitl(y)
					}
				}()
			}
		}() == func() int32 {
			if 8 == 4 {
				return int32(libc.X__FLOAT_BITS(float32(ywant)) >> int32(31))
			} else {
				return func() int32 {
					if 8 == 8 {
						return int32(libc.X__DOUBLE_BITS(float64(ywant)) >> int32(63))
					} else {
						return libc.X__signbitl(ywant)
					}
				}()
			}
		}() {
			return 1
		} else {
			return 0
		}
	}()
}
