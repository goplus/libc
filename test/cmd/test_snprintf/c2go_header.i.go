package main

import unsafe "unsafe"

type size_t = uint64
type time_t = int64
type clockid_t = int32
type struct_timespec struct {
	tv_sec  int64
	Xbf_0   int32
	tv_nsec int64
	Xbf_1   int32
}
type pthread_t = *struct___pthread
type pthread_once_t = int32
type pthread_key_t = uint32
type pthread_spinlock_t = int32
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
type struct_sched_param struct {
	sched_priority int32
	__reserved1    int32
	__reserved2    [4]int64
	__reserved3    int32
}
type timer_t = unsafe.Pointer
type clock_t = int64
type locale_t = *struct___locale_struct
type struct_tm struct {
	tm_sec      int32
	tm_min      int32
	tm_hour     int32
	tm_mday     int32
	tm_mon      int32
	tm_year     int32
	tm_wday     int32
	tm_yday     int32
	tm_isdst    int32
	__tm_gmtoff int64
	__tm_zone   *int8
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
type ssize_t = int64
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
type float_t = float32
type double_t = float64

func __FLOAT_BITS_cgo15(__f float32) uint32 {
	type _cgoa_16 struct {
		__f float32
	}
	var __u _cgoa_16
	__u.__f = __f
	return *(*uint32)(unsafe.Pointer(&__u))
}
func __DOUBLE_BITS_cgo17(__f float64) uint64 {
	type _cgoa_18 struct {
		__f float64
	}
	var __u _cgoa_18
	__u.__f = __f
	return *(*uint64)(unsafe.Pointer(&__u))
}
func __islessf_cgo19(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(float64(__x))&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(float64(__y))&18446744073709551615 > 2047<<52 {
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
func __isless_cgo20(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(__x)&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(__y)&18446744073709551615 > 2047<<52 {
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
func __islessl_cgo21(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(float64(__x))&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(float64(__y))&18446744073709551615 > 2047<<52 {
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
func __islessequalf_cgo22(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(float64(__x))&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(float64(__y))&18446744073709551615 > 2047<<52 {
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
func __islessequal_cgo23(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(__x)&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(__y)&18446744073709551615 > 2047<<52 {
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
func __islessequall_cgo24(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(float64(__x))&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(float64(__y))&18446744073709551615 > 2047<<52 {
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
func __islessgreaterf_cgo25(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(float64(__x))&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(float64(__y))&18446744073709551615 > 2047<<52 {
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
func __islessgreater_cgo26(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(__x)&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(__y)&18446744073709551615 > 2047<<52 {
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
func __islessgreaterl_cgo27(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(float64(__x))&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(float64(__y))&18446744073709551615 > 2047<<52 {
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
func __isgreaterf_cgo28(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(float64(__x))&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(float64(__y))&18446744073709551615 > 2047<<52 {
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
func __isgreater_cgo29(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(__x)&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(__y)&18446744073709551615 > 2047<<52 {
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
func __isgreaterl_cgo30(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(float64(__x))&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(float64(__y))&18446744073709551615 > 2047<<52 {
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
func __isgreaterequalf_cgo31(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(float64(__x))&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(float64(__y))&18446744073709551615 > 2047<<52 {
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
func __isgreaterequal_cgo32(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(__x)&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(__y)&18446744073709551615 > 2047<<52 {
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
func __isgreaterequall_cgo33(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if __FLOAT_BITS_cgo15(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if __DOUBLE_BITS_cgo17(float64(__x))&18446744073709551615 > 2047<<52 {
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
							if __FLOAT_BITS_cgo15(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if __DOUBLE_BITS_cgo17(float64(__y))&18446744073709551615 > 2047<<52 {
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
type intptr_t = int64
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
type uid_t = uint32
type gid_t = uint32
type useconds_t = uint32




type _cgoa_41 struct {
	__i [14]int32
}

type _cgoa_43 struct {
	__i [10]int32
}

type _cgoa_45 struct {
	__i [12]int32
}

type _cgoa_47 struct {
	__i [14]int32
}

type _cgoa_49 struct {
	__i [8]int32
}

