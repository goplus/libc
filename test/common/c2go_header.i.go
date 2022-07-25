package common

import libc "github.com/goplus/libc"

type struct_d_d struct {
	file *int8
	line int32
	r    int32
	x    float64
	y    float64
	dy   float32
	e    int32
}
type Struct_d_d = struct_d_d
type struct_f_f struct {
	file *int8
	line int32
	r    int32
	x    float32
	y    float32
	dy   float32
	e    int32
}
type Struct_f_f = struct_f_f
type struct_l_l struct {
	file *int8
	line int32
	r    int32
	x    float64
	y    float64
	dy   float32
	e    int32
}
type Struct_l_l = struct_l_l
type struct_ff_f struct {
	file *int8
	line int32
	r    int32
	x    float32
	x2   float32
	y    float32
	dy   float32
	e    int32
}
type Struct_ff_f = struct_ff_f
type struct_dd_d struct {
	file *int8
	line int32
	r    int32
	x    float64
	x2   float64
	y    float64
	dy   float32
	e    int32
}
type Struct_dd_d = struct_dd_d
type struct_ll_l struct {
	file *int8
	line int32
	r    int32
	x    float64
	x2   float64
	y    float64
	dy   float32
	e    int32
}
type Struct_ll_l = struct_ll_l
type struct_d_di struct {
	file *int8
	line int32
	r    int32
	x    float64
	y    float64
	dy   float32
	i    int64
	e    int32
}
type Struct_d_di = struct_d_di
type struct_f_fi struct {
	file *int8
	line int32
	r    int32
	x    float32
	y    float32
	dy   float32
	i    int64
	e    int32
}
type Struct_f_fi = struct_f_fi
type struct_l_li struct {
	file *int8
	line int32
	r    int32
	x    float64
	y    float64
	dy   float32
	i    int64
	e    int32
}
type Struct_l_li = struct_l_li
type struct_di_d struct {
	file *int8
	line int32
	r    int32
	x    float64
	i    int64
	y    float64
	dy   float32
	e    int32
}
type Struct_di_d = struct_di_d
type struct_fi_f struct {
	file *int8
	line int32
	r    int32
	x    float32
	i    int64
	y    float32
	dy   float32
	e    int32
}
type Struct_fi_f = struct_fi_f
type struct_li_l struct {
	file *int8
	line int32
	r    int32
	x    float64
	i    int64
	y    float64
	dy   float32
	e    int32
}
type Struct_li_l = struct_li_l
type struct_d_i struct {
	file *int8
	line int32
	r    int32
	x    float64
	i    int64
	e    int32
}
type Struct_d_i = struct_d_i
type struct_f_i struct {
	file *int8
	line int32
	r    int32
	x    float32
	i    int64
	e    int32
}
type Struct_f_i = struct_f_i
type struct_l_i struct {
	file *int8
	line int32
	r    int32
	x    float64
	i    int64
	e    int32
}
type Struct_l_i = struct_l_i
type struct_d_dd struct {
	file *int8
	line int32
	r    int32
	x    float64
	y    float64
	dy   float32
	y2   float64
	dy2  float32
	e    int32
}
type Struct_d_dd = struct_d_dd
type struct_f_ff struct {
	file *int8
	line int32
	r    int32
	x    float32
	y    float32
	dy   float32
	y2   float32
	dy2  float32
	e    int32
}
type Struct_f_ff = struct_f_ff
type struct_l_ll struct {
	file *int8
	line int32
	r    int32
	x    float64
	y    float64
	dy   float32
	y2   float64
	dy2  float32
	e    int32
}
type Struct_l_ll = struct_l_ll
type struct_ff_fi struct {
	file *int8
	line int32
	r    int32
	x    float32
	x2   float32
	y    float32
	dy   float32
	i    int64
	e    int32
}
type Struct_ff_fi = struct_ff_fi
type struct_dd_di struct {
	file *int8
	line int32
	r    int32
	x    float64
	x2   float64
	y    float64
	dy   float32
	i    int64
	e    int32
}
type Struct_dd_di = struct_dd_di
type struct_ll_li struct {
	file *int8
	line int32
	r    int32
	x    float64
	x2   float64
	y    float64
	dy   float32
	i    int64
	e    int32
}
type Struct_ll_li = struct_ll_li
type struct_fff_f struct {
	file *int8
	line int32
	r    int32
	x    float32
	x2   float32
	x3   float32
	y    float32
	dy   float32
	e    int32
}
type Struct_fff_f = struct_fff_f
type struct_ddd_d struct {
	file *int8
	line int32
	r    int32
	x    float64
	x2   float64
	x3   float64
	y    float64
	dy   float32
	e    int32
}
type Struct_ddd_d = struct_ddd_d
type struct_lll_l struct {
	file *int8
	line int32
	r    int32
	x    float64
	x2   float64
	x3   float64
	y    float64
	dy   float32
	e    int32
}
type Struct_lll_l = struct_lll_l

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
