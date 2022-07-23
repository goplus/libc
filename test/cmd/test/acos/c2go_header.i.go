package acos

import libc "github.com/goplus/libc"

type __int8_t = int8
type __uint8_t = uint8
type __int16_t = int16
type __uint16_t = uint16
type __int32_t = int32
type __uint32_t = uint32
type __int64_t = int64
type __uint64_t = uint64
type __darwin_intptr_t = int64
type __darwin_natural_t = uint32
type __darwin_ct_rune_t = int32
type _cgoa_1_acos struct {
	__mbstate8 [128]int8
}
type __mbstate_t = _cgoa_1_acos
type __darwin_mbstate_t = _cgoa_1_acos
type __darwin_ptrdiff_t = int64
type __darwin_size_t = uint64
type __darwin_va_list = []interface {
}
type __darwin_wchar_t = int32
type __darwin_rune_t = int32
type __darwin_wint_t = int32
type __darwin_clock_t = uint64
type __darwin_socklen_t = uint32
type __darwin_ssize_t = int64
type __darwin_time_t = int64
type natural_t = uint32
type integer_t = int32
type vm_offset_t = uint64
type vm_size_t = uint64
type mach_vm_address_t = uint64
type mach_vm_offset_t = uint64
type mach_vm_size_t = uint64
type vm_map_offset_t = uint64
type vm_map_address_t = uint64
type vm_map_size_t = uint64
type mach_port_context_t = uint64
type boolean_t = uint32
type kern_return_t = int32
type struct_qelem struct {
	q_forw *struct_qelem
	q_back *struct_qelem
	q_data *int8
}
type struct_d_d struct {
	file *int8
	line int32
	r    int32
	x    float64
	y    float64
	dy   float32
	e    int32
}
type struct_f_f struct {
	file *int8
	line int32
	r    int32
	x    float32
	y    float32
	dy   float32
	e    int32
}
type struct_l_l struct {
	file *int8
	line int32
	r    int32
	x    float64
	y    float64
	dy   float32
	e    int32
}
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
type struct_d_i struct {
	file *int8
	line int32
	r    int32
	x    float64
	i    int64
	e    int32
}
type struct_f_i struct {
	file *int8
	line int32
	r    int32
	x    float32
	i    int64
	e    int32
}
type struct_l_i struct {
	file *int8
	line int32
	r    int32
	x    float64
	i    int64
	e    int32
}
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

func checkexcept(got int32, want int32, r int32) int32 {
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
func checkexceptall(got int32, want int32, r int32) int32 {
	return func() int32 {
		if got == want {
			return 1
		} else {
			return 0
		}
	}()
}
func checkulp(d float32, r int32) int32 {
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
func checkcr(y float64, ywant float64, r int32) int32 {
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
