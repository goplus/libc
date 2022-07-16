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
type pthread_mutexattr_t = _cgoa_1
type Pthread_mutexattr_t = _cgoa_1
type _cgoa_2 struct {
	__attr uint32
}
type pthread_condattr_t = _cgoa_2
type Pthread_condattr_t = _cgoa_2
type _cgoa_3 struct {
	__attr uint32
}
type pthread_barrierattr_t = _cgoa_3
type Pthread_barrierattr_t = _cgoa_3
type _cgoa_4 struct {
	__attr [2]uint32
}
type pthread_rwlockattr_t = _cgoa_4
type Pthread_rwlockattr_t = _cgoa_4
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
type pthread_attr_t = _cgoa_5
type Pthread_attr_t = _cgoa_5
type _cgoa_8 struct {
	__i [10]int32
}
type _cgoa_7 struct {
	__u _cgoa_8
}
type pthread_mutex_t = _cgoa_7
type Pthread_mutex_t = _cgoa_7
type _cgoa_10 struct {
	__i [12]int32
}
type _cgoa_9 struct {
	__u _cgoa_10
}
type pthread_cond_t = _cgoa_9
type Pthread_cond_t = _cgoa_9
type _cgoa_12 struct {
	__i [14]int32
}
type _cgoa_11 struct {
	__u _cgoa_12
}
type pthread_rwlock_t = _cgoa_11
type Pthread_rwlock_t = _cgoa_11
type _cgoa_14 struct {
	__i [8]int32
}
type _cgoa_13 struct {
	__u _cgoa_14
}
type pthread_barrier_t = _cgoa_13
type Pthread_barrier_t = _cgoa_13
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
type float_t = float32
type Float_t = float32
type double_t = float64
type Double_t = float64

func X__FLOAT_BITS(__f float32) uint32 {
	type _cgoa_15 struct {
		__f float32
	}
	var __u _cgoa_15
	__u.__f = __f
	return *(*uint32)(unsafe.Pointer(&__u))
}
func X__DOUBLE_BITS(__f float64) uint64 {
	type _cgoa_16 struct {
		__f float64
	}
	var __u _cgoa_16
	__u.__f = __f
	return *(*uint64)(unsafe.Pointer(&__u))
}
func __islessf(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
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
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(__x)&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(__y)&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
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
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(__x) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(__y) == int32(0) {
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
						if X__FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
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
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(__x)&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(__y)&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
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
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(__x) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(__y) == int32(0) {
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
						if X__FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
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
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(__x)&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(__y)&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
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
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(__x) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(__y) == int32(0) {
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
						if X__FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
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
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(__x)&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(__y)&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
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
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(__x) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(__y) == int32(0) {
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
						if X__FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
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
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(__x)&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(__y)&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
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
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(__x) == int32(0) {
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
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(__y) == int32(0) {
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
func __isspace(_c int32) int32 {
	return func() int32 {
		if _c == ' ' || uint32(_c)-uint32('\t') < uint32(5) {
			return 1
		} else {
			return 0
		}
	}()
}

type wchar_t = uint32
type Wchar_t = uint32
type _cgoa_50 struct {
	quot int32
	rem  int32
}
type div_t = _cgoa_50
type _cgoa_51 struct {
	quot int64
	rem  int64
}
type ldiv_t = _cgoa_51
type _cgoa_52 struct {
	quot int64
	rem  int64
}
type lldiv_t = _cgoa_52
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
type _cgoa_84 struct {
	__ll int64
	__ld float64
}
type max_align_t = _cgoa_84
type Max_align_t = _cgoa_84
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
type fexcept_t = uint64
type _cgoa_104 struct {
	__cw uint64
}
type fenv_t = _cgoa_104

func __bswap16(__x uint16) uint16 {
	return uint16(int32(__x)<<int32(8) | int32(__x)>>int32(8))
}
func __bswap32(__x uint32) uint32 {
	return __x>>int32(24) | __x>>int32(8)&uint32(65280) | __x<<int32(8)&uint32(16711680) | __x<<int32(24)
}
func __bswap64(__x uint64) uint64 {
	return (uint64(__bswap32(uint32(__x)))+uint64(0))<<int32(32) | uint64(__bswap32(uint32(__x>>int32(32))))
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
func a_fetch_add(p *int32, v int32) int32 {
	var old int32
	for {
		old = a_ll(p)
		if !!(a_sc(p, int32(uint32(old)+uint32(v))) != 0) {
			break
		}
	}
	return old
}
func a_fetch_and(p *int32, v int32) int32 {
	var old int32
	for {
		old = a_ll(p)
		if !!(a_sc(p, old&v) != 0) {
			break
		}
	}
	return old
}
func a_fetch_or(p *int32, v int32) int32 {
	var old int32
	for {
		old = a_ll(p)
		if !!(a_sc(p, old|v) != 0) {
			break
		}
	}
	return old
}
func a_and(p *int32, v int32) {
	a_fetch_and(p, v)
}
func a_or(p *int32, v int32) {
	a_fetch_or(p, v)
}
func a_inc(p *int32) {
	a_fetch_add(p, int32(1))
}
func a_dec(p *int32) {
	a_fetch_add(p, -1)
}
func a_store(p *int32, v int32) {
	a_swap(p, v)
}
func a_barrier() {
	var tmp int32 = int32(0)
	_ = tmp
	a_cas(&tmp, int32(0), int32(0))
}
func a_and_64(p *uint64, v uint64) {
	type _cgoa_784 struct {
		v uint64
	}
	var u _cgoa_784
	u.v = v
	if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(0))*4))+uint32(1) != 0 {
		a_and((*int32)(unsafe.Pointer(p)), int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(0))*4))))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(1))*4))+uint32(1) != 0 {
		a_and((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(p))))+uintptr(int32(1))*4)), int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(1))*4))))
	}
}
func a_or_64(p *uint64, v uint64) {
	type _cgoa_785 struct {
		v uint64
	}
	var u _cgoa_785
	u.v = v
	if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(0))*4)) != 0 {
		a_or((*int32)(unsafe.Pointer(p)), int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(0))*4))))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(1))*4)) != 0 {
		a_or((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(p))))+uintptr(int32(1))*4)), int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(1))*4))))
	}
}
func a_or_l(p unsafe.Pointer, v int64) {
	if 8 == 4 {
		a_or((*int32)(p), int32(v))
	} else {
		a_or_64((*uint64)(p), uint64(v))
	}
}
func a_crash() {
	*(*int8)(nil) = int8(0)
}
func a_ctz_32(x uint32) int32 {
	return int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&debruijn32_cgo786)))) + uintptr(x&-x*uint32(124511785)>>int32(27)))))
}

var debruijn32_cgo786 [32]int8 = [32]int8{int8(0), int8(1), int8(23), int8(2), int8(29), int8(24), int8(19), int8(3), int8(30), int8(27), int8(25), int8(11), int8(20), int8(8), int8(4), int8(13), int8(31), int8(22), int8(28), int8(18), int8(26), int8(10), int8(7), int8(12), int8(21), int8(17), int8(9), int8(6), int8(16), int8(5), int8(15), int8(14)}

func a_ctz_64(x uint64) int32 {
	if false {
		var y uint32 = uint32(x)
		if !(y != 0) {
			y = uint32(x >> int32(32))
			return int32(32) + a_ctz_32(y)
		}
		return a_ctz_32(y)
	}
	return int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&debruijn64_cgo787)))) + uintptr(x&-x*uint64(157587932685088877)>>int32(58)))))
}

var debruijn64_cgo787 [64]int8 = [64]int8{int8(0), int8(1), int8(2), int8(53), int8(3), int8(7), int8(54), int8(27), int8(4), int8(38), int8(41), int8(8), int8(34), int8(55), int8(48), int8(28), int8(62), int8(5), int8(39), int8(46), int8(44), int8(42), int8(22), int8(9), int8(24), int8(35), int8(59), int8(56), int8(49), int8(18), int8(29), int8(11), int8(63), int8(52), int8(6), int8(26), int8(37), int8(40), int8(33), int8(47), int8(61), int8(45), int8(43), int8(21), int8(23), int8(58), int8(17), int8(10), int8(51), int8(25), int8(36), int8(32), int8(60), int8(20), int8(57), int8(16), int8(50), int8(31), int8(19), int8(15), int8(30), int8(14), int8(13), int8(12)}

func a_ctz_l(x uint64) int32 {
	return func() int32 {
		if false {
			return a_ctz_32(uint32(x))
		} else {
			return a_ctz_64(uint64(x))
		}
	}()
}
func a_clz_64(x uint64) int32 {
	var y uint32
	var r int32
	if x>>int32(32) != 0 {
		func() int32 {
			y = uint32(x >> int32(32))
			return func() (_cgo_ret int32) {
				_cgo_addr := &r
				*_cgo_addr = int32(0)
				return *_cgo_addr
			}()
		}()
	} else {
		func() int32 {
			y = uint32(x)
			return func() (_cgo_ret int32) {
				_cgo_addr := &r
				*_cgo_addr = int32(32)
				return *_cgo_addr
			}()
		}()
	}
	if y>>int32(16) != 0 {
		y >>= int32(16)
	} else {
		r |= int32(16)
	}
	if y>>int32(8) != 0 {
		y >>= int32(8)
	} else {
		r |= int32(8)
	}
	if y>>int32(4) != 0 {
		y >>= int32(4)
	} else {
		r |= int32(4)
	}
	if y>>int32(2) != 0 {
		y >>= int32(2)
	} else {
		r |= int32(2)
	}
	return r | func() int32 {
		if !(y>>int32(1) != 0) {
			return 1
		} else {
			return 0
		}
	}()
}
func a_clz_32(x uint32) int32 {
	x >>= int32(1)
	x |= x >> int32(1)
	x |= x >> int32(2)
	x |= x >> int32(4)
	x |= x >> int32(8)
	x |= x >> int32(16)
	x++
	return int32(31) - a_ctz_32(x)
}

type uid_t = uint32
type Uid_t = uint32
type gid_t = uint32
type Gid_t = uint32
type useconds_t = uint32
type Useconds_t = uint32
type mode_t = uint32
type Mode_t = uint32
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
type _cgoa_869 struct {
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
	__st_atim32        _cgoa_869
	__st_mtim32        _cgoa_869
	__st_ctim32        _cgoa_869
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
type struct_iovec struct {
	iov_base unsafe.Pointer
	iov_len  uint64
}
type struct_winsize struct {
	ws_row    uint16
	ws_col    uint16
	ws_xpixel uint16
	ws_ypixel uint16
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
type _cgoa_1233 struct {
	quot int64
	rem  int64
}
type imaxdiv_t = _cgoa_1233
