package common

import (
	unsafe "unsafe"
	libc "github.com/goplus/libc"
)

func eulpf(x float32) int32 {
	type _cgoa_1_mtest struct {
		f float32
	}
	var u _cgoa_1_mtest
	u.f = x
	var e int32 = int32(*(*uint32)(unsafe.Pointer(&u)) >> int32(23) & uint32(255))
	if !(e != 0) {
		e++
	}
	return e - int32(127) - int32(23)
}
func eulp(x float64) int32 {
	type _cgoa_2_mtest struct {
		f float64
	}
	var u _cgoa_2_mtest
	u.f = x
	var e int32 = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(52) & uint64(2047))
	if !(e != 0) {
		e++
	}
	return e - int32(1023) - int32(52)
}
func eulpl(x float64) int32 {
	return eulp(float64(x))
}
func Ulperrf(got float32, want float32, dwant float32) float32 {
	if func() int32 {
		if 4 == 4 {
			return func() int32 {
				if libc.X__FLOAT_BITS(got)&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 4 == 8 {
					return func() int32 {
						if libc.X__DOUBLE_BITS(float64(got))&9223372036854775807 > 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if libc.X__fpclassifyl(float64(got)) == int32(0) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 && func() int32 {
		if 4 == 4 {
			return func() int32 {
				if libc.X__FLOAT_BITS(want)&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 4 == 8 {
					return func() int32 {
						if libc.X__DOUBLE_BITS(float64(want))&9223372036854775807 > 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if libc.X__fpclassifyl(float64(want)) == int32(0) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 {
		return float32(int32(0))
	}
	if got == want {
		if func() int32 {
			if 4 == 4 {
				return int32(libc.X__FLOAT_BITS(got) >> int32(31))
			} else {
				return func() int32 {
					if 4 == 8 {
						return int32(libc.X__DOUBLE_BITS(float64(got)) >> int32(63))
					} else {
						return libc.X__signbitl(float64(got))
					}
				}()
			}
		}() == func() int32 {
			if 4 == 4 {
				return int32(libc.X__FLOAT_BITS(want) >> int32(31))
			} else {
				return func() int32 {
					if 4 == 8 {
						return int32(libc.X__DOUBLE_BITS(float64(want)) >> int32(63))
					} else {
						return libc.X__signbitl(float64(want))
					}
				}()
			}
		}() {
			return dwant
		}
		return libc.X__builtin_inff()
	}
	if func() int32 {
		if 4 == 4 {
			return func() int32 {
				if libc.X__FLOAT_BITS(got)&uint32(2147483647) == uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 4 == 8 {
					return func() int32 {
						if libc.X__DOUBLE_BITS(float64(got))&9223372036854775807 == 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if libc.X__fpclassifyl(float64(got)) == int32(1) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 {
		got = libc.Copysignf(float32(1.7014118346046923e+38), got)
		want *= float32(0.5)
	}
	return float32(libc.Scalbn(float64(got-want), -eulpf(want)) + float64(dwant))
}
func Ulperr(got float64, want float64, dwant float32) float32 {
	if func() int32 {
		if 8 == 4 {
			return func() int32 {
				if libc.X__FLOAT_BITS(float32(got))&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 8 == 8 {
					return func() int32 {
						if libc.X__DOUBLE_BITS(got)&9223372036854775807 > 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if libc.X__fpclassifyl(float64(got)) == int32(0) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 && func() int32 {
		if 8 == 4 {
			return func() int32 {
				if libc.X__FLOAT_BITS(float32(want))&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 8 == 8 {
					return func() int32 {
						if libc.X__DOUBLE_BITS(want)&9223372036854775807 > 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if libc.X__fpclassifyl(float64(want)) == int32(0) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 {
		return float32(int32(0))
	}
	if got == want {
		if func() int32 {
			if 8 == 4 {
				return int32(libc.X__FLOAT_BITS(float32(got)) >> int32(31))
			} else {
				return func() int32 {
					if 8 == 8 {
						return int32(libc.X__DOUBLE_BITS(got) >> int32(63))
					} else {
						return libc.X__signbitl(float64(got))
					}
				}()
			}
		}() == func() int32 {
			if 8 == 4 {
				return int32(libc.X__FLOAT_BITS(float32(want)) >> int32(31))
			} else {
				return func() int32 {
					if 8 == 8 {
						return int32(libc.X__DOUBLE_BITS(want) >> int32(63))
					} else {
						return libc.X__signbitl(float64(want))
					}
				}()
			}
		}() {
			return dwant
		}
		return libc.X__builtin_inff()
	}
	if func() int32 {
		if 8 == 4 {
			return func() int32 {
				if libc.X__FLOAT_BITS(float32(got))&uint32(2147483647) == uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 8 == 8 {
					return func() int32 {
						if libc.X__DOUBLE_BITS(got)&9223372036854775807 == 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if libc.X__fpclassifyl(float64(got)) == int32(1) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 {
		got = libc.Copysign(8.9884656743115795e+307, got)
		want *= float64(0.5)
	}
	return float32(libc.Scalbn(got-want, -eulp(want)) + float64(dwant))
}
func Ulperrl(got float64, want float64, dwant float32) float32 {
	return Ulperr(float64(got), float64(want), dwant)
}

type _cgoa_3_mtest struct {
	flag int32
	s    *int8
}

var eflags [5]_cgoa_3_mtest = [5]_cgoa_3_mtest{_cgoa_3_mtest{int32(0), (*int8)(unsafe.Pointer(&[8]int8{'I', 'N', 'E', 'X', 'A', 'C', 'T', '\x00'}))}, _cgoa_3_mtest{int32(0), (*int8)(unsafe.Pointer(&[8]int8{'I', 'N', 'V', 'A', 'L', 'I', 'D', '\x00'}))}, _cgoa_3_mtest{int32(0), (*int8)(unsafe.Pointer(&[10]int8{'D', 'I', 'V', 'B', 'Y', 'Z', 'E', 'R', 'O', '\x00'}))}, _cgoa_3_mtest{int32(0), (*int8)(unsafe.Pointer(&[10]int8{'U', 'N', 'D', 'E', 'R', 'F', 'L', 'O', 'W', '\x00'}))}, _cgoa_3_mtest{int32(0), (*int8)(unsafe.Pointer(&[9]int8{'O', 'V', 'E', 'R', 'F', 'L', 'O', 'W', '\x00'}))}}

func Estr(f int32) *int8 {
	var p *int8 = (*int8)(unsafe.Pointer(&buf_cgo4_mtest))
	var i int32
	var all int32 = int32(0)
	for i = int32(0); uint64(i) < 5; i++ {
		if f&(*(*_cgoa_3_mtest)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_3_mtest)(unsafe.Pointer(&eflags)))) + uintptr(i)*16))).flag != 0 {
			*(*uintptr)(unsafe.Pointer(&p)) += uintptr(libc.Sprintf(p, (*int8)(unsafe.Pointer(&[5]int8{'%', 's', '%', 's', '\x00'})), func() *int8 {
				if all != 0 {
					return (*int8)(unsafe.Pointer(&[2]int8{'|', '\x00'}))
				} else {
					return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
				}
			}(), (*(*_cgoa_3_mtest)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_3_mtest)(unsafe.Pointer(&eflags)))) + uintptr(i)*16))).s))
			all |= (*(*_cgoa_3_mtest)(unsafe.Pointer(uintptr(unsafe.Pointer((*_cgoa_3_mtest)(unsafe.Pointer(&eflags)))) + uintptr(i)*16))).flag
		}
	}
	if all != f {
		*(*uintptr)(unsafe.Pointer(&p)) += uintptr(libc.Sprintf(p, (*int8)(unsafe.Pointer(&[5]int8{'%', 's', '%', 'd', '\x00'})), func() *int8 {
			if all != 0 {
				return (*int8)(unsafe.Pointer(&[2]int8{'|', '\x00'}))
			} else {
				return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
			}
		}(), f & ^all))
		all = f
	}
	*(*uintptr)(unsafe.Pointer(&p)) += uintptr(libc.Sprintf(p, (*int8)(unsafe.Pointer(&[3]int8{'%', 's', '\x00'})), func() *int8 {
		if all != 0 {
			return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
		} else {
			return (*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'}))
		}
	}()))
	return (*int8)(unsafe.Pointer(&buf_cgo4_mtest))
}

var buf_cgo4_mtest [256]int8

func Rstr(r int32) *int8 {
	switch r {
	case int32(0):
		return (*int8)(unsafe.Pointer(&[3]int8{'R', 'N', '\x00'}))
	}
	return (*int8)(unsafe.Pointer(&[3]int8{'R', '?', '\x00'}))
}
