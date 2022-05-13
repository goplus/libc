package libc

import unsafe "unsafe"

type size_t = uint64
type Size_t = uint64
type time_t = int64
type Time_t = int64
type clockid_t = int32
type Clockid_t = int32
type struct_timespec struct {
	tv_sec  int64
	Xbf_0   int32
	tv_nsec int64
	Xbf_1   int32
}
type pthread_t = *struct___pthread
type Pthread_t = *struct___pthread
type pthread_once_t = int32
type Pthread_once_t = int32
type pthread_key_t = uint32
type Pthread_key_t = uint32
type pthread_spinlock_t = int32
type Pthread_spinlock_t = int32
type _cgoa_1 struct {
	__attr uint32
}
type _cgoa_2 struct {
	__attr uint32
}
type _cgoa_3 struct {
	__attr uint32
}
type _cgoa_4 struct {
	__attr [2]uint32
}
type struct___sigset_t struct {
	__bits [16]uint64
}
type sigset_t = struct___sigset_t
type Sigset_t = struct___sigset_t
type _cgoa_6 struct {
	__i [14]int32
}
type _cgoa_5 struct {
	__u _cgoa_6
}
type _cgoa_8 struct {
	__i [10]int32
}
type _cgoa_7 struct {
	__u _cgoa_8
}
type _cgoa_10 struct {
	__i [12]int32
}
type _cgoa_9 struct {
	__u _cgoa_10
}
type _cgoa_12 struct {
	__i [14]int32
}
type _cgoa_11 struct {
	__u _cgoa_12
}
type _cgoa_14 struct {
	__i [8]int32
}
type _cgoa_13 struct {
	__u _cgoa_14
}
type pid_t = int32
type Pid_t = int32
type struct_sched_param struct {
	sched_priority int32
	__reserved1    int32
	__reserved2    [4]int64
	__reserved3    int32
}
type timer_t = unsafe.Pointer
type Timer_t = unsafe.Pointer
type clock_t = int64
type Clock_t = int64
type locale_t = *struct___locale_struct
type Locale_t = *struct___locale_struct
type struct_tm struct {
	tm_sec    int32
	tm_min    int32
	tm_hour   int32
	tm_mday   int32
	tm_mon    int32
	tm_year   int32
	tm_wday   int32
	tm_yday   int32
	tm_isdst  int32
	tm_gmtoff int64
	tm_zone   *int8
}
type struct_itimerspec struct {
	it_interval struct_timespec
	it_value    struct_timespec
}
type struct___ptcb struct {
	__f    func(unsafe.Pointer)
	__x    unsafe.Pointer
	__next *struct___ptcb
}
type struct___pthread struct {
	tid       int32
	errno_val int32
	sys_r1    int64
	locale    *struct___locale_struct
}
type wchar_t = uint32
type Wchar_t = uint32
type _cgoa_15 struct {
	quot int32
	rem  int32
}
type _cgoa_16 struct {
	quot int64
	rem  int64
}
type _cgoa_17 struct {
	quot int64
	rem  int64
}
type ssize_t = int64
type Ssize_t = int64
type off_t = int64
type Off_t = int64
type FILE = struct__IO_FILE
type va_list = []interface {
}
type Va_list = []interface {
}
type __isoc_va_list = []interface {
}
type union__G_fpos64_t struct {
	__opaque [16]int8
}
type fpos_t = union__G_fpos64_t
type struct___locale_struct struct {
	cat [6]*struct___locale_map
}
type struct_tls_module struct {
	next   *struct_tls_module
	image  unsafe.Pointer
	len    uint64
	size   uint64
	align  uint64
	offset uint64
}
type struct___libc struct {
	can_do_threads  int8
	threaded        int8
	secure          int8
	need_locks      int8
	threads_minus_1 int32
	auxv            *uint64
	tls_head        *struct_tls_module
	tls_size        uint64
	tls_align       uint64
	tls_cnt         uint64
	page_size       uint64
	global_locale   struct___locale_struct
}




type _cgoa_23 struct {
	__i [14]int32
}

type _cgoa_25 struct {
	__i [10]int32
}

type _cgoa_27 struct {
	__i [12]int32
}

type _cgoa_29 struct {
	__i [14]int32
}

type _cgoa_31 struct {
	__i [8]int32
}

type _cgoa_32 struct {
	__ll int64
	__ld float64
}
type ptrdiff_t = int64
type Ptrdiff_t = int64
type struct_lconv struct {
	decimal_point      *int8
	thousands_sep      *int8
	grouping           *int8
	int_curr_symbol    *int8
	currency_symbol    *int8
	mon_decimal_point  *int8
	mon_thousands_sep  *int8
	mon_grouping       *int8
	positive_sign      *int8
	negative_sign      *int8
	int_frac_digits    int8
	frac_digits        int8
	p_cs_precedes      int8
	p_sep_by_space     int8
	n_cs_precedes      int8
	n_sep_by_space     int8
	p_sign_posn        int8
	n_sign_posn        int8
	int_p_cs_precedes  int8
	int_p_sep_by_space int8
	int_n_cs_precedes  int8
	int_n_sep_by_space int8
	int_p_sign_posn    int8
	int_n_sign_posn    int8
}



type struct___locale_map struct {
	map_     unsafe.Pointer
	map_size uint64
	name     [24]int8
	next     *struct___locale_map
}




type _cgoa_43 struct {
	__i [14]int32
}

type _cgoa_45 struct {
	__i [10]int32
}

type _cgoa_47 struct {
	__i [12]int32
}

type _cgoa_49 struct {
	__i [14]int32
}

type _cgoa_51 struct {
	__i [8]int32
}





type _cgoa_57 struct {
	__i [14]int32
}

type _cgoa_59 struct {
	__i [10]int32
}

type _cgoa_61 struct {
	__i [12]int32
}

type _cgoa_63 struct {
	__i [14]int32
}

type _cgoa_65 struct {
	__i [8]int32
}

type float_t = float32
type Float_t = float32
type double_t = float64
type Double_t = float64

func __FLOAT_BITS(__f float32) uint32 {
	type _cgoa_66 struct {
		__f float32
	}
	var __u _cgoa_66
	__u.__f = __f
	return *(*uint32)(unsafe.Pointer(&__u))
}
func __DOUBLE_BITS(__f float64) uint64 {
	type _cgoa_67 struct {
		__f float64
	}
	var __u _cgoa_67
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

type uintptr_t = uint64
type Uintptr_t = uint64
type intptr_t = int64
type Intptr_t = int64
type int8_t = int8
type Int8_t = int8
type int16_t = int16
type Int16_t = int16
type int32_t = int32
type Int32_t = int32
type int64_t = int64
type Int64_t = int64
type intmax_t = int64
type Intmax_t = int64
type uint8_t = uint8
type Uint8_t = uint8
type uint16_t = uint16
type Uint16_t = uint16
type uint32_t = uint32
type Uint32_t = uint32
type uint64_t = uint64
type Uint64_t = uint64
type uintmax_t = uint64
type Uintmax_t = uint64
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




type _cgoa_74 struct {
	__i [14]int32
}

type _cgoa_76 struct {
	__i [10]int32
}

type _cgoa_78 struct {
	__i [12]int32
}

type _cgoa_80 struct {
	__i [14]int32
}

type _cgoa_82 struct {
	__i [8]int32
}


func __bswap16(__x uint16) uint16 {
	return uint16(int32(__x)<<8 | int32(__x)>>8)
}
func __bswap32(__x uint32) uint32 {
	return __x>>24 | __x>>8&uint32(65280) | __x<<8&uint32(16711680) | __x<<24
}
func __bswap64(__x uint64) uint64 {
	return (uint64(__bswap32(uint32(__x)))+0)<<32 | uint64(__bswap32(uint32(__x>>32)))
}
func eval_as_float(x float32) float32 {
	var y float32 = x
	return y
}
func eval_as_double(x float64) float64 {
	var y float64 = x
	return y
}
func fp_barrierf(x float32) float32 {
	var y float32 = x
	_ = y
	return y
}
func fp_barrier(x float64) float64 {
	var y float64 = x
	_ = y
	return y
}
func fp_barrierl(x float64) float64 {
	var y float64 = x
	_ = y
	return y
}
func fp_force_evalf(x float32) {
	var y float32
	_ = y
	y = x
}
func fp_force_eval(x float64) {
	var y float64
	_ = y
	y = x
}
func fp_force_evall(x float64) {
	var y float64
	_ = y
	y = x
}





type _cgoa_88 struct {
	__i [14]int32
}

type _cgoa_90 struct {
	__i [10]int32
}

type _cgoa_92 struct {
	__i [12]int32
}

type _cgoa_94 struct {
	__i [14]int32
}

type _cgoa_96 struct {
	__i [8]int32
}





type _cgoa_103 struct {
	__i [14]int32
}

type _cgoa_105 struct {
	__i [10]int32
}

type _cgoa_107 struct {
	__i [12]int32
}

type _cgoa_109 struct {
	__i [14]int32
}

type _cgoa_111 struct {
	__i [8]int32
}





type _cgoa_117 struct {
	__i [14]int32
}

type _cgoa_119 struct {
	__i [10]int32
}

type _cgoa_121 struct {
	__i [12]int32
}

type _cgoa_123 struct {
	__i [14]int32
}

type _cgoa_125 struct {
	__i [8]int32
}





type _cgoa_132 struct {
	__i [14]int32
}

type _cgoa_134 struct {
	__i [10]int32
}

type _cgoa_136 struct {
	__i [12]int32
}

type _cgoa_138 struct {
	__i [14]int32
}

type _cgoa_140 struct {
	__i [8]int32
}





type _cgoa_146 struct {
	__i [14]int32
}

type _cgoa_148 struct {
	__i [10]int32
}

type _cgoa_150 struct {
	__i [12]int32
}

type _cgoa_152 struct {
	__i [14]int32
}

type _cgoa_154 struct {
	__i [8]int32
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





type _cgoa_160 struct {
	__i [14]int32
}

type _cgoa_162 struct {
	__i [10]int32
}

type _cgoa_164 struct {
	__i [12]int32
}

type _cgoa_166 struct {
	__i [14]int32
}

type _cgoa_168 struct {
	__i [8]int32
}





type _cgoa_174 struct {
	__i [14]int32
}

type _cgoa_176 struct {
	__i [10]int32
}

type _cgoa_178 struct {
	__i [12]int32
}

type _cgoa_180 struct {
	__i [14]int32
}

type _cgoa_182 struct {
	__i [8]int32
}





type _cgoa_188 struct {
	__i [14]int32
}

type _cgoa_190 struct {
	__i [10]int32
}

type _cgoa_192 struct {
	__i [12]int32
}

type _cgoa_194 struct {
	__i [14]int32
}

type _cgoa_196 struct {
	__i [8]int32
}





type _cgoa_202 struct {
	__i [14]int32
}

type _cgoa_204 struct {
	__i [10]int32
}

type _cgoa_206 struct {
	__i [12]int32
}

type _cgoa_208 struct {
	__i [14]int32
}

type _cgoa_210 struct {
	__i [8]int32
}





type _cgoa_216 struct {
	__i [14]int32
}

type _cgoa_218 struct {
	__i [10]int32
}

type _cgoa_220 struct {
	__i [12]int32
}

type _cgoa_222 struct {
	__i [14]int32
}

type _cgoa_224 struct {
	__i [8]int32
}





type _cgoa_230 struct {
	__i [14]int32
}

type _cgoa_232 struct {
	__i [10]int32
}

type _cgoa_234 struct {
	__i [12]int32
}

type _cgoa_236 struct {
	__i [14]int32
}

type _cgoa_238 struct {
	__i [8]int32
}





type _cgoa_244 struct {
	__i [14]int32
}

type _cgoa_246 struct {
	__i [10]int32
}

type _cgoa_248 struct {
	__i [12]int32
}

type _cgoa_250 struct {
	__i [14]int32
}

type _cgoa_252 struct {
	__i [8]int32
}





type _cgoa_258 struct {
	__i [14]int32
}

type _cgoa_260 struct {
	__i [10]int32
}

type _cgoa_262 struct {
	__i [12]int32
}

type _cgoa_264 struct {
	__i [14]int32
}

type _cgoa_266 struct {
	__i [8]int32
}





type _cgoa_272 struct {
	__i [14]int32
}

type _cgoa_274 struct {
	__i [10]int32
}

type _cgoa_276 struct {
	__i [12]int32
}

type _cgoa_278 struct {
	__i [14]int32
}

type _cgoa_280 struct {
	__i [8]int32
}





type _cgoa_286 struct {
	__i [14]int32
}

type _cgoa_288 struct {
	__i [10]int32
}

type _cgoa_290 struct {
	__i [12]int32
}

type _cgoa_292 struct {
	__i [14]int32
}

type _cgoa_294 struct {
	__i [8]int32
}





type _cgoa_300 struct {
	__i [14]int32
}

type _cgoa_302 struct {
	__i [10]int32
}

type _cgoa_304 struct {
	__i [12]int32
}

type _cgoa_306 struct {
	__i [14]int32
}

type _cgoa_308 struct {
	__i [8]int32
}

type uid_t = uint32
type Uid_t = uint32
type gid_t = uint32
type Gid_t = uint32
type useconds_t = uint32
type Useconds_t = uint32




type _cgoa_314 struct {
	__i [14]int32
}

type _cgoa_316 struct {
	__i [10]int32
}

type _cgoa_318 struct {
	__i [12]int32
}

type _cgoa_320 struct {
	__i [14]int32
}

type _cgoa_322 struct {
	__i [8]int32
}








type _cgoa_331 struct {
	__i [14]int32
}

type _cgoa_333 struct {
	__i [10]int32
}

type _cgoa_335 struct {
	__i [12]int32
}

type _cgoa_337 struct {
	__i [14]int32
}

type _cgoa_339 struct {
	__i [8]int32
}

type mode_t = uint32
type Mode_t = uint32
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





type _cgoa_346 struct {
	__i [14]int32
}

type _cgoa_348 struct {
	__i [10]int32
}

type _cgoa_350 struct {
	__i [12]int32
}

type _cgoa_352 struct {
	__i [14]int32
}

type _cgoa_354 struct {
	__i [8]int32
}





type _cgoa_361 struct {
	__i [14]int32
}

type _cgoa_363 struct {
	__i [10]int32
}

type _cgoa_365 struct {
	__i [12]int32
}

type _cgoa_367 struct {
	__i [14]int32
}

type _cgoa_369 struct {
	__i [8]int32
}

type struct_flock struct {
	l_type   int16
	l_whence int16
	l_start  int64
	l_len    int64
	l_pid    int32
}
type nlink_t = uint32
type Nlink_t = uint32
type ino_t = uint64
type Ino_t = uint64
type dev_t = uint64
type Dev_t = uint64
type blksize_t = int64
type Blksize_t = int64
type blkcnt_t = int64
type Blkcnt_t = int64
type _cgoa_370 struct {
	tv_sec  int64
	tv_nsec int64
}
type struct_stat struct {
	st_dev             uint64
	__st_dev_padding   int32
	__st_ino_truncated int64
	st_mode            uint32
	st_nlink           uint32
	st_uid             uint32
	st_gid             uint32
	st_rdev            uint64
	__st_rdev_padding  int32
	st_size            int64
	st_blksize         int64
	st_blocks          int64
	__st_atim32        _cgoa_370
	__st_mtim32        _cgoa_370
	__st_ctim32        _cgoa_370
	st_ino             uint64
	st_atim            struct_timespec
	st_mtim            struct_timespec
	st_ctim            struct_timespec
}
type struct_kstat struct {
	st_dev             uint64
	__st_dev_padding   int32
	__st_ino_truncated int64
	st_mode            uint32
	st_nlink           uint32
	st_uid             uint32
	st_gid             uint32
	st_rdev            uint64
	__st_rdev_padding  int32
	st_size            int64
	st_blksize         int64
	st_blocks          int64
	st_atime_sec       int64
	st_atime_nsec      int64
	st_mtime_sec       int64
	st_mtime_nsec      int64
	st_ctime_sec       int64
	st_ctime_nsec      int64
	st_ino             uint64
}




type _cgoa_376 struct {
	__i [14]int32
}

type _cgoa_378 struct {
	__i [10]int32
}

type _cgoa_380 struct {
	__i [12]int32
}

type _cgoa_382 struct {
	__i [14]int32
}

type _cgoa_384 struct {
	__i [8]int32
}





type _cgoa_391 struct {
	__i [14]int32
}

type _cgoa_393 struct {
	__i [10]int32
}

type _cgoa_395 struct {
	__i [12]int32
}

type _cgoa_397 struct {
	__i [14]int32
}

type _cgoa_399 struct {
	__i [8]int32
}








type _cgoa_409 struct {
	__i [14]int32
}

type _cgoa_411 struct {
	__i [10]int32
}

type _cgoa_413 struct {
	__i [12]int32
}

type _cgoa_415 struct {
	__i [14]int32
}

type _cgoa_417 struct {
	__i [8]int32
}








type _cgoa_427 struct {
	__i [14]int32
}

type _cgoa_429 struct {
	__i [10]int32
}

type _cgoa_431 struct {
	__i [12]int32
}

type _cgoa_433 struct {
	__i [14]int32
}

type _cgoa_435 struct {
	__i [8]int32
}








type _cgoa_445 struct {
	__i [14]int32
}

type _cgoa_447 struct {
	__i [10]int32
}

type _cgoa_449 struct {
	__i [12]int32
}

type _cgoa_451 struct {
	__i [14]int32
}

type _cgoa_453 struct {
	__i [8]int32
}





type _cgoa_459 struct {
	__i [14]int32
}

type _cgoa_461 struct {
	__i [10]int32
}

type _cgoa_463 struct {
	__i [12]int32
}

type _cgoa_465 struct {
	__i [14]int32
}

type _cgoa_467 struct {
	__i [8]int32
}





type _cgoa_473 struct {
	__i [14]int32
}

type _cgoa_475 struct {
	__i [10]int32
}

type _cgoa_477 struct {
	__i [12]int32
}

type _cgoa_479 struct {
	__i [14]int32
}

type _cgoa_481 struct {
	__i [8]int32
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
	read         func(*struct__IO_FILE, *uint8, uint64) uint64
	write        func(*struct__IO_FILE, *uint8, uint64) uint64
	seek         func(*struct__IO_FILE, int64, int32) int64
	buf          *uint8
	buf_size     uint64
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




type _cgoa_488 struct {
	__i [14]int32
}

type _cgoa_490 struct {
	__i [10]int32
}

type _cgoa_492 struct {
	__i [12]int32
}

type _cgoa_494 struct {
	__i [14]int32
}

type _cgoa_496 struct {
	__i [8]int32
}





type _cgoa_502 struct {
	__i [14]int32
}

type _cgoa_504 struct {
	__i [10]int32
}

type _cgoa_506 struct {
	__i [12]int32
}

type _cgoa_508 struct {
	__i [14]int32
}

type _cgoa_510 struct {
	__i [8]int32
}





type _cgoa_516 struct {
	__i [14]int32
}

type _cgoa_518 struct {
	__i [10]int32
}

type _cgoa_520 struct {
	__i [12]int32
}

type _cgoa_522 struct {
	__i [14]int32
}

type _cgoa_524 struct {
	__i [8]int32
}





type _cgoa_530 struct {
	__i [14]int32
}

type _cgoa_532 struct {
	__i [10]int32
}

type _cgoa_534 struct {
	__i [12]int32
}

type _cgoa_536 struct {
	__i [14]int32
}

type _cgoa_538 struct {
	__i [8]int32
}





type _cgoa_546 struct {
	__i [14]int32
}

type _cgoa_548 struct {
	__i [10]int32
}

type _cgoa_550 struct {
	__i [12]int32
}

type _cgoa_552 struct {
	__i [14]int32
}

type _cgoa_554 struct {
	__i [8]int32
}





type _cgoa_560 struct {
	__i [14]int32
}

type _cgoa_562 struct {
	__i [10]int32
}

type _cgoa_564 struct {
	__i [12]int32
}

type _cgoa_566 struct {
	__i [14]int32
}

type _cgoa_568 struct {
	__i [8]int32
}





type _cgoa_574 struct {
	__i [14]int32
}

type _cgoa_576 struct {
	__i [10]int32
}

type _cgoa_578 struct {
	__i [12]int32
}

type _cgoa_580 struct {
	__i [14]int32
}

type _cgoa_582 struct {
	__i [8]int32
}

type struct_iovec struct {
	iov_base unsafe.Pointer
	iov_len  uint64
}




type _cgoa_588 struct {
	__i [14]int32
}

type _cgoa_590 struct {
	__i [10]int32
}

type _cgoa_592 struct {
	__i [12]int32
}

type _cgoa_594 struct {
	__i [14]int32
}

type _cgoa_596 struct {
	__i [8]int32
}





type _cgoa_602 struct {
	__i [14]int32
}

type _cgoa_604 struct {
	__i [10]int32
}

type _cgoa_606 struct {
	__i [12]int32
}

type _cgoa_608 struct {
	__i [14]int32
}

type _cgoa_610 struct {
	__i [8]int32
}





type _cgoa_616 struct {
	__i [14]int32
}

type _cgoa_618 struct {
	__i [10]int32
}

type _cgoa_620 struct {
	__i [12]int32
}

type _cgoa_622 struct {
	__i [14]int32
}

type _cgoa_624 struct {
	__i [8]int32
}





type _cgoa_631 struct {
	__i [14]int32
}

type _cgoa_633 struct {
	__i [10]int32
}

type _cgoa_635 struct {
	__i [12]int32
}

type _cgoa_637 struct {
	__i [14]int32
}

type _cgoa_639 struct {
	__i [8]int32
}

type struct_winsize struct {
	ws_row    uint16
	ws_col    uint16
	ws_xpixel uint16
	ws_ypixel uint16
}




type _cgoa_645 struct {
	__i [14]int32
}

type _cgoa_647 struct {
	__i [10]int32
}

type _cgoa_649 struct {
	__i [12]int32
}

type _cgoa_651 struct {
	__i [14]int32
}

type _cgoa_653 struct {
	__i [8]int32
}





type _cgoa_661 struct {
	__i [14]int32
}

type _cgoa_663 struct {
	__i [10]int32
}

type _cgoa_665 struct {
	__i [12]int32
}

type _cgoa_667 struct {
	__i [14]int32
}

type _cgoa_669 struct {
	__i [8]int32
}





type _cgoa_676 struct {
	__i [14]int32
}

type _cgoa_678 struct {
	__i [10]int32
}

type _cgoa_680 struct {
	__i [12]int32
}

type _cgoa_682 struct {
	__i [14]int32
}

type _cgoa_684 struct {
	__i [8]int32
}





type _cgoa_691 struct {
	__i [14]int32
}

type _cgoa_693 struct {
	__i [10]int32
}

type _cgoa_695 struct {
	__i [12]int32
}

type _cgoa_697 struct {
	__i [14]int32
}

type _cgoa_699 struct {
	__i [8]int32
}





type _cgoa_706 struct {
	__i [14]int32
}

type _cgoa_708 struct {
	__i [10]int32
}

type _cgoa_710 struct {
	__i [12]int32
}

type _cgoa_712 struct {
	__i [14]int32
}

type _cgoa_714 struct {
	__i [8]int32
}





type _cgoa_720 struct {
	__i [14]int32
}

type _cgoa_722 struct {
	__i [10]int32
}

type _cgoa_724 struct {
	__i [12]int32
}

type _cgoa_726 struct {
	__i [14]int32
}

type _cgoa_728 struct {
	__i [8]int32
}





type wint_t = uint32
type Wint_t = uint32
type wctype_t = uint64
type Wctype_t = uint64
type struct___mbstate_t struct {
	__opaque1 uint32
	__opaque2 uint32
}
type mbstate_t = struct___mbstate_t
type Mbstate_t = struct___mbstate_t
type _cgoa_733 struct {
	quot int64
	rem  int64
}




type _cgoa_750 struct {
	__i [14]int32
}

type _cgoa_752 struct {
	__i [10]int32
}

type _cgoa_754 struct {
	__i [12]int32
}

type _cgoa_756 struct {
	__i [14]int32
}

type _cgoa_758 struct {
	__i [8]int32
}





type _cgoa_764 struct {
	__i [14]int32
}

type _cgoa_766 struct {
	__i [10]int32
}

type _cgoa_768 struct {
	__i [12]int32
}

type _cgoa_770 struct {
	__i [14]int32
}

type _cgoa_772 struct {
	__i [8]int32
}

