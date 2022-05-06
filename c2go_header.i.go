package libc

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
	locale    *struct___locale_struct
}




type _cgoa_20 struct {
	__i [14]int32
}

type _cgoa_22 struct {
	__i [10]int32
}

type _cgoa_24 struct {
	__i [12]int32
}

type _cgoa_26 struct {
	__i [14]int32
}

type _cgoa_28 struct {
	__i [8]int32
}

type float_t = float32
type double_t = float64

func __FLOAT_BITS(__f float32) uint32 {
	type _cgoa_29 struct {
		__f float32
	}
	var __u _cgoa_29
	__u.__f = __f
	return *(*uint32)(unsafe.Pointer(&__u))
}
func __DOUBLE_BITS(__f float64) uint64 {
	type _cgoa_30 struct {
		__f float64
	}
	var __u _cgoa_30
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




type _cgoa_37 struct {
	__i [14]int32
}

type _cgoa_39 struct {
	__i [10]int32
}

type _cgoa_41 struct {
	__i [12]int32
}

type _cgoa_43 struct {
	__i [14]int32
}

type _cgoa_45 struct {
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





type _cgoa_51 struct {
	__i [14]int32
}

type _cgoa_53 struct {
	__i [10]int32
}

type _cgoa_55 struct {
	__i [12]int32
}

type _cgoa_57 struct {
	__i [14]int32
}

type _cgoa_59 struct {
	__i [8]int32
}





type _cgoa_66 struct {
	__i [14]int32
}

type _cgoa_68 struct {
	__i [10]int32
}

type _cgoa_70 struct {
	__i [12]int32
}

type _cgoa_72 struct {
	__i [14]int32
}

type _cgoa_74 struct {
	__i [8]int32
}





type _cgoa_80 struct {
	__i [14]int32
}

type _cgoa_82 struct {
	__i [10]int32
}

type _cgoa_84 struct {
	__i [12]int32
}

type _cgoa_86 struct {
	__i [14]int32
}

type _cgoa_88 struct {
	__i [8]int32
}





type _cgoa_95 struct {
	__i [14]int32
}

type _cgoa_97 struct {
	__i [10]int32
}

type _cgoa_99 struct {
	__i [12]int32
}

type _cgoa_101 struct {
	__i [14]int32
}

type _cgoa_103 struct {
	__i [8]int32
}





type _cgoa_109 struct {
	__i [14]int32
}

type _cgoa_111 struct {
	__i [10]int32
}

type _cgoa_113 struct {
	__i [12]int32
}

type _cgoa_115 struct {
	__i [14]int32
}

type _cgoa_117 struct {
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





type _cgoa_123 struct {
	__i [14]int32
}

type _cgoa_125 struct {
	__i [10]int32
}

type _cgoa_127 struct {
	__i [12]int32
}

type _cgoa_129 struct {
	__i [14]int32
}

type _cgoa_131 struct {
	__i [8]int32
}





type _cgoa_137 struct {
	__i [14]int32
}

type _cgoa_139 struct {
	__i [10]int32
}

type _cgoa_141 struct {
	__i [12]int32
}

type _cgoa_143 struct {
	__i [14]int32
}

type _cgoa_145 struct {
	__i [8]int32
}





type _cgoa_151 struct {
	__i [14]int32
}

type _cgoa_153 struct {
	__i [10]int32
}

type _cgoa_155 struct {
	__i [12]int32
}

type _cgoa_157 struct {
	__i [14]int32
}

type _cgoa_159 struct {
	__i [8]int32
}





type _cgoa_165 struct {
	__i [14]int32
}

type _cgoa_167 struct {
	__i [10]int32
}

type _cgoa_169 struct {
	__i [12]int32
}

type _cgoa_171 struct {
	__i [14]int32
}

type _cgoa_173 struct {
	__i [8]int32
}





type _cgoa_179 struct {
	__i [14]int32
}

type _cgoa_181 struct {
	__i [10]int32
}

type _cgoa_183 struct {
	__i [12]int32
}

type _cgoa_185 struct {
	__i [14]int32
}

type _cgoa_187 struct {
	__i [8]int32
}

type wchar_t = uint32
type _cgoa_188 struct {
	__ll int64
	__ld float64
}
type ptrdiff_t = int64
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
type _cgoa_189 struct {
	quot int32
	rem  int32
}
type _cgoa_190 struct {
	quot int64
	rem  int64
}
type _cgoa_191 struct {
	quot int64
	rem  int64
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
type struct___locale_map struct {
	map_     unsafe.Pointer
	map_size uint64
	name     [24]int8
	next     *struct___locale_map
}




type _cgoa_199 struct {
	__i [14]int32
}

type _cgoa_201 struct {
	__i [10]int32
}

type _cgoa_203 struct {
	__i [12]int32
}

type _cgoa_205 struct {
	__i [14]int32
}

type _cgoa_207 struct {
	__i [8]int32
}

type uid_t = uint32
type gid_t = uint32
type useconds_t = uint32
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





type _cgoa_213 struct {
	__i [14]int32
}

type _cgoa_215 struct {
	__i [10]int32
}

type _cgoa_217 struct {
	__i [12]int32
}

type _cgoa_219 struct {
	__i [14]int32
}

type _cgoa_221 struct {
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




type _cgoa_227 struct {
	__i [14]int32
}

type _cgoa_229 struct {
	__i [10]int32
}

type _cgoa_231 struct {
	__i [12]int32
}

type _cgoa_233 struct {
	__i [14]int32
}

type _cgoa_235 struct {
	__i [8]int32
}





type _cgoa_241 struct {
	__i [14]int32
}

type _cgoa_243 struct {
	__i [10]int32
}

type _cgoa_245 struct {
	__i [12]int32
}

type _cgoa_247 struct {
	__i [14]int32
}

type _cgoa_249 struct {
	__i [8]int32
}





type _cgoa_255 struct {
	__i [14]int32
}

type _cgoa_257 struct {
	__i [10]int32
}

type _cgoa_259 struct {
	__i [12]int32
}

type _cgoa_261 struct {
	__i [14]int32
}

type _cgoa_263 struct {
	__i [8]int32
}





type _cgoa_269 struct {
	__i [14]int32
}

type _cgoa_271 struct {
	__i [10]int32
}

type _cgoa_273 struct {
	__i [12]int32
}

type _cgoa_275 struct {
	__i [14]int32
}

type _cgoa_277 struct {
	__i [8]int32
}





type _cgoa_284 struct {
	__i [14]int32
}

type _cgoa_286 struct {
	__i [10]int32
}

type _cgoa_288 struct {
	__i [12]int32
}

type _cgoa_290 struct {
	__i [14]int32
}

type _cgoa_292 struct {
	__i [8]int32
}





type _cgoa_298 struct {
	__i [14]int32
}

type _cgoa_300 struct {
	__i [10]int32
}

type _cgoa_302 struct {
	__i [12]int32
}

type _cgoa_304 struct {
	__i [14]int32
}

type _cgoa_306 struct {
	__i [8]int32
}





type _cgoa_312 struct {
	__i [14]int32
}

type _cgoa_314 struct {
	__i [10]int32
}

type _cgoa_316 struct {
	__i [12]int32
}

type _cgoa_318 struct {
	__i [14]int32
}

type _cgoa_320 struct {
	__i [8]int32
}

type struct_iovec struct {
	iov_base unsafe.Pointer
	iov_len  uint64
}




type _cgoa_326 struct {
	__i [14]int32
}

type _cgoa_328 struct {
	__i [10]int32
}

type _cgoa_330 struct {
	__i [12]int32
}

type _cgoa_332 struct {
	__i [14]int32
}

type _cgoa_334 struct {
	__i [8]int32
}





type _cgoa_340 struct {
	__i [14]int32
}

type _cgoa_342 struct {
	__i [10]int32
}

type _cgoa_344 struct {
	__i [12]int32
}

type _cgoa_346 struct {
	__i [14]int32
}

type _cgoa_348 struct {
	__i [8]int32
}





type _cgoa_354 struct {
	__i [14]int32
}

type _cgoa_356 struct {
	__i [10]int32
}

type _cgoa_358 struct {
	__i [12]int32
}

type _cgoa_360 struct {
	__i [14]int32
}

type _cgoa_362 struct {
	__i [8]int32
}





type _cgoa_368 struct {
	__i [14]int32
}

type _cgoa_370 struct {
	__i [10]int32
}

type _cgoa_372 struct {
	__i [12]int32
}

type _cgoa_374 struct {
	__i [14]int32
}

type _cgoa_376 struct {
	__i [8]int32
}

type struct_winsize struct {
	ws_row    uint16
	ws_col    uint16
	ws_xpixel uint16
	ws_ypixel uint16
}




type _cgoa_382 struct {
	__i [14]int32
}

type _cgoa_384 struct {
	__i [10]int32
}

type _cgoa_386 struct {
	__i [12]int32
}

type _cgoa_388 struct {
	__i [14]int32
}

type _cgoa_390 struct {
	__i [8]int32
}





type _cgoa_398 struct {
	__i [14]int32
}

type _cgoa_400 struct {
	__i [10]int32
}

type _cgoa_402 struct {
	__i [12]int32
}

type _cgoa_404 struct {
	__i [14]int32
}

type _cgoa_406 struct {
	__i [8]int32
}





type _cgoa_413 struct {
	__i [14]int32
}

type _cgoa_415 struct {
	__i [10]int32
}

type _cgoa_417 struct {
	__i [12]int32
}

type _cgoa_419 struct {
	__i [14]int32
}

type _cgoa_421 struct {
	__i [8]int32
}





type _cgoa_428 struct {
	__i [14]int32
}

type _cgoa_430 struct {
	__i [10]int32
}

type _cgoa_432 struct {
	__i [12]int32
}

type _cgoa_434 struct {
	__i [14]int32
}

type _cgoa_436 struct {
	__i [8]int32
}





type _cgoa_443 struct {
	__i [14]int32
}

type _cgoa_445 struct {
	__i [10]int32
}

type _cgoa_447 struct {
	__i [12]int32
}

type _cgoa_449 struct {
	__i [14]int32
}

type _cgoa_451 struct {
	__i [8]int32
}





type _cgoa_457 struct {
	__i [14]int32
}

type _cgoa_459 struct {
	__i [10]int32
}

type _cgoa_461 struct {
	__i [12]int32
}

type _cgoa_463 struct {
	__i [14]int32
}

type _cgoa_465 struct {
	__i [8]int32
}





type wint_t = uint32
type wctype_t = uint64
type struct___mbstate_t struct {
	__opaque1 uint32
	__opaque2 uint32
}
type mbstate_t = struct___mbstate_t
type _cgoa_470 struct {
	quot int64
	rem  int64
}




type _cgoa_478 struct {
	__i [14]int32
}

type _cgoa_480 struct {
	__i [10]int32
}

type _cgoa_482 struct {
	__i [12]int32
}

type _cgoa_484 struct {
	__i [14]int32
}

type _cgoa_486 struct {
	__i [8]int32
}

