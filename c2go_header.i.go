package libc

import unsafe "unsafe"

type size_t = uint32
type ssize_t = int32
type off_t = int64
type FILE = struct__IO_FILE
type va_list = []interface {
}
type __isoc_va_list = []interface {
}
type union__G_fpos64_t struct {
	__opaque [16]int8
}
type fpos_t = union__G_fpos64_t
type syscall_arg_t = int64

func __alt_socketcall(sys int32, sock int32, cp int32, a int64, b int64, c int64, d int64, e int64, f int64) int64 {
	var r int64
	if cp != 0 {
		r = __syscall_cp(int64(sys), int64(a), int64(b), int64(c), int64(d), int64(e), int64(f))
	} else {
		r = __syscall6(int64(sys), int64(a), int64(b), int64(c), int64(d), int64(e), int64(f))
	}
	if r != int64(-38) {
		return r
	}
	return r
}

type struct__IO_FILE struct {
	flags        uint32
	rpos         *uint8
	rend         *uint8
	close        func(*struct__IO_FILE) int32
	wend         *uint8
	wpos         *uint8
	mustbezero_1 *uint8
	wbase        *uint8
	read         func(*struct__IO_FILE, *uint8, uint32) uint32
	write        func(*struct__IO_FILE, *uint8, uint32) uint32
	seek         func(*struct__IO_FILE, int64, int32) int64
	buf          *uint8
	buf_size     uint32
	prev         *struct__IO_FILE
	next         *struct__IO_FILE
	fd           int32
	pipe_pid     int32
	lockcount    int64
	mode         int32
	lock         int32
	lbf          int32
	cookie       unsafe.Pointer
	off          int64
	getln_buf    *int8
	mustbezero_2 unsafe.Pointer
	shend        *uint8
	shlim        int64
	shcnt        int64
	prev_locked  *struct__IO_FILE
	next_locked  *struct__IO_FILE
	locale       *struct___locale_struct
}

func __isspace(_c int32) int32 {
	return func() int32 {
		if _c == ' ' || uint32(_c)-uint32('\t') < uint32(5) {
			return 1
		} else {
			return 0
		}
	}()
}

type locale_t = *struct___locale_struct
type wchar_t = uint32
type _cgoa_1 struct {
	__ll int64
	__ld float64
}
type ptrdiff_t = int32
type _cgoa_2 struct {
	quot int32
	rem  int32
}
type _cgoa_3 struct {
	quot int64
	rem  int64
}
type _cgoa_4 struct {
	quot int64
	rem  int64
}
type wint_t = uint32
type wctype_t = uint64
type struct___mbstate_t struct {
	__opaque1 uint32
	__opaque2 uint32
}
type mbstate_t = struct___mbstate_t
type uintptr_t = uint32
type intptr_t = int32
type int8_t = int8
type int16_t = int16
type int32_t = int32
type int64_t = int64
type intmax_t = int64
type uint8_t = uint8
type uint16_t = uint16
type uint32_t = uint32
type uint64_t = uint64
type uintmax_t = uint64
type int_fast8_t = int8
type int_fast64_t = int64
type int_least8_t = int8
type int_least16_t = int16
type int_least32_t = int32
type int_least64_t = int64
type uint_fast8_t = uint8
type uint_fast64_t = uint64
type uint_least8_t = uint8
type uint_least16_t = uint16
type uint_least32_t = uint32
type uint_least64_t = uint64
type int_fast16_t = int32
type int_fast32_t = int32
type uint_fast16_t = uint32
type uint_fast32_t = uint32
type _cgoa_5 struct {
	quot int64
	rem  int64
}
type float_t = float32
type double_t = float64

func __FLOAT_BITS(__f float32) uint32 {
	type _cgoa_6 struct {
		__f float32
	}
	var __u _cgoa_6
	__u.__f = __f
	return *(*uint32)(unsafe.Pointer(&__u))
}
func __DOUBLE_BITS(__f float64) uint64 {
	type _cgoa_7 struct {
		__f float64
	}
	var __u _cgoa_7
	__u.__f = __f
	return *(*uint64)(unsafe.Pointer(&__u))
}
func __islessf(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if __FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(float64(__x))&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(float64(__x)) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if __FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(float64(__y))&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(float64(__y)) == 0 {
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
		}() != 0) && __x < __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __isless(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(__x)&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(float64(__x)) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if __FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(__y)&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(float64(__y)) == 0 {
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
		}() != 0) && __x < __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __islessl(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(float64(__x))&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(__x) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if __FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(float64(__y))&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(__y) == 0 {
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
		}() != 0) && __x < __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __islessequalf(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if __FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(float64(__x))&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(float64(__x)) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if __FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(float64(__y))&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(float64(__y)) == 0 {
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
		}() != 0) && __x <= __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __islessequal(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(__x)&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(float64(__x)) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if __FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(__y)&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(float64(__y)) == 0 {
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
		}() != 0) && __x <= __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __islessequall(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(float64(__x))&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(__x) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if __FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(float64(__y))&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(__y) == 0 {
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
		}() != 0) && __x <= __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __islessgreaterf(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if __FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(float64(__x))&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(float64(__x)) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if __FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(float64(__y))&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(float64(__y)) == 0 {
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
		}() != 0) && __x != __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __islessgreater(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(__x)&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(float64(__x)) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if __FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(__y)&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(float64(__y)) == 0 {
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
		}() != 0) && __x != __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __islessgreaterl(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(float64(__x))&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(__x) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if __FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(float64(__y))&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(__y) == 0 {
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
		}() != 0) && __x != __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __isgreaterf(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if __FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(float64(__x))&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(float64(__x)) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if __FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(float64(__y))&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(float64(__y)) == 0 {
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
		}() != 0) && __x > __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __isgreater(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(__x)&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(float64(__x)) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if __FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(__y)&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(float64(__y)) == 0 {
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
		}() != 0) && __x > __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __isgreaterl(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(float64(__x))&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(__x) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if __FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(float64(__y))&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(__y) == 0 {
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
		}() != 0) && __x > __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __isgreaterequalf(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if __FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(float64(__x))&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(float64(__x)) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if __FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(float64(__y))&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(float64(__y)) == 0 {
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
		}() != 0) && __x >= __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __isgreaterequal(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(__x)&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(float64(__x)) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if __FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(__y)&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(float64(__y)) == 0 {
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
		}() != 0) && __x >= __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __isgreaterequall(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS(float64(__x))&18446744073709551615 > 2047<<52 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if __fpclassifyl(__x) == 0 {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return int32(func() int {
					func() int {
						_ = __y
						return 0
					}()
					return 1
				}())
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if __FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS(float64(__y))&18446744073709551615 > 2047<<52 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if __fpclassifyl(__y) == 0 {
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
		}() != 0) && __x >= __y {
			return 1
		} else {
			return 0
		}
	}()
}
