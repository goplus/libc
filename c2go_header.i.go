package libc

import unsafe "unsafe"

type float_t = float32
type double_t = float64

func __FLOAT_BITS(__f float32) uint32 {
	type _cgoa_1 struct {
		__f float32
	}
	var __u _cgoa_1
	__u.__f = __f
	return *(*uint32)(unsafe.Pointer(&__u))
}
func __DOUBLE_BITS(__f float64) uint64 {
	type _cgoa_2 struct {
		__f float64
	}
	var __u _cgoa_2
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
type size_t = uint64
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
type _cgoa_7 struct {
	__attr uint32
}
type _cgoa_8 struct {
	__attr uint32
}
type _cgoa_9 struct {
	__attr uint32
}
type _cgoa_10 struct {
	__attr [2]uint32
}
type struct___sigset_t struct {
	__bits [16]uint64
}
type sigset_t = struct___sigset_t
type _cgoa_12 struct {
	__i [14]int32
}
type _cgoa_11 struct {
	__u _cgoa_12
}
type _cgoa_14 struct {
	__i [10]int32
}
type _cgoa_13 struct {
	__u _cgoa_14
}
type _cgoa_16 struct {
	__i [12]int32
}
type _cgoa_15 struct {
	__u _cgoa_16
}
type _cgoa_18 struct {
	__i [14]int32
}
type _cgoa_17 struct {
	__u _cgoa_18
}
type _cgoa_20 struct {
	__i [8]int32
}
type _cgoa_19 struct {
	__u _cgoa_20
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
type uid_t = uint32
type stack_t = struct_sigaltstack
type greg_t = int32
type gregset_t = [18]int32
type struct_sigcontext struct {
	trap_no       uint64
	error_code    uint64
	oldmask       uint64
	arm_r0        uint64
	arm_r1        uint64
	arm_r2        uint64
	arm_r3        uint64
	arm_r4        uint64
	arm_r5        uint64
	arm_r6        uint64
	arm_r7        uint64
	arm_r8        uint64
	arm_r9        uint64
	arm_r10       uint64
	arm_fp        uint64
	arm_ip        uint64
	arm_sp        uint64
	arm_lr        uint64
	arm_pc        uint64
	arm_cpsr      uint64
	fault_address uint64
}
type mcontext_t = struct_sigcontext
type struct_sigaltstack struct {
	ss_sp    unsafe.Pointer
	ss_flags int32
	ss_size  uint64
}
type struct___ucontext struct {
	uc_flags    uint64
	uc_link     *struct___ucontext
	uc_stack    struct_sigaltstack
	uc_mcontext struct_sigcontext
	uc_sigmask  struct___sigset_t
	uc_regspace [64]uint64
}
type ucontext_t = struct___ucontext
type union_sigval struct {
	sival_ptr unsafe.Pointer
}
type _cgoa_25 struct {
	si_pid int32
	si_uid uint32
}
type _cgoa_26 struct {
	si_timerid int32
	si_overrun int32
}
type _cgoa_24 struct {
	__piduid _cgoa_25
}
type _cgoa_28 struct {
	si_status int32
	si_utime  int64
	si_stime  int64
}
type _cgoa_27 struct {
	__sigchld _cgoa_28
}
type _cgoa_23 struct {
	__first  _cgoa_24
	__second _cgoa_27
}
type _cgoa_31 struct {
	si_lower unsafe.Pointer
	si_upper unsafe.Pointer
}
type _cgoa_30 struct {
	__addr_bnd _cgoa_31
}
type _cgoa_29 struct {
	si_addr     unsafe.Pointer
	si_addr_lsb int16
	__first     _cgoa_30
}
type _cgoa_32 struct {
	si_band int64
	si_fd   int32
}
type _cgoa_33 struct {
	si_call_addr unsafe.Pointer
	si_syscall   int32
	si_arch      uint32
}
type _cgoa_22 struct {
	__pad [112]int8
}
type _cgoa_21 struct {
	si_signo    int32
	si_errno    int32
	si_code     int32
	__si_fields _cgoa_22
}
type _cgoa_34 struct {
	sa_handler func(int32)
}
type struct_sigaction struct {
	__sa_handler _cgoa_34
	sa_mask      struct___sigset_t
	sa_flags     int32
	sa_restorer  func()
}
type _cgoa_36 struct {
	sigev_notify_function   func(union_sigval)
	sigev_notify_attributes *_cgoa_11
}
type _cgoa_35 struct {
	__pad [48]int8
}
type struct_sigevent struct {
	sigev_value  union_sigval
	sigev_signo  int32
	sigev_notify int32
	__sev_fields _cgoa_35
}
type sig_t = func(int32)
type sig_atomic_t = int32
type mode_t = uint32
type wchar_t = uint32
type _cgoa_37 struct {
	quot int32
	rem  int32
}
type _cgoa_38 struct {
	quot int64
	rem  int64
}
type _cgoa_39 struct {
	quot int64
	rem  int64
}
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

func a_swap(p *int32, v int32) int32 {
	var old int32
	for {
		old = *p
		if !(a_cas(p, old, v) != old) {
			break
		}
	}
	return old
}
func a_fetch_add(p *int32, v int32) int32 {
	var old int32
	for {
		old = *p
		if !(a_cas(p, old, int32(uint32(old)+uint32(v))) != old) {
			break
		}
	}
	return old
}
func a_fetch_and(p *int32, v int32) int32 {
	var old int32
	for {
		old = *p
		if !(a_cas(p, old, old&v) != old) {
			break
		}
	}
	return old
}
func a_fetch_or(p *int32, v int32) int32 {
	var old int32
	for {
		old = *p
		if !(a_cas(p, old, old|v) != old) {
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
	a_fetch_add(p, 1)
}
func a_dec(p *int32) {
	a_fetch_add(p, -1)
}
func a_store(p *int32, v int32) {
	a_swap(p, v)
}
func a_barrier() {
	var tmp int32 = 0
	_ = tmp
	a_cas(&tmp, 0, 0)
}
func a_and_64(p *uint64, v uint64) {
	type _cgoa_40 struct {
		v uint64
	}
	var u _cgoa_40
	u.v = v
	if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(0)*4))+uint32(1) != 0 {
		a_and((*int32)(unsafe.Pointer(p)), int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(0)*4))))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(1)*4))+uint32(1) != 0 {
		a_and((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(p))))+uintptr(1)*4)), int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(1)*4))))
	}
}
func a_or_64(p *uint64, v uint64) {
	type _cgoa_41 struct {
		v uint64
	}
	var u _cgoa_41
	u.v = v
	if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(0)*4)) != 0 {
		a_or((*int32)(unsafe.Pointer(p)), int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(0)*4))))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(1)*4)) != 0 {
		a_or((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(p))))+uintptr(1)*4)), int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(1)*4))))
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
	var debruijn32 [32]int8 = [32]int8{int8(0), int8(1), int8(23), int8(2), int8(29), int8(24), int8(19), int8(3), int8(30), int8(27), int8(25), int8(11), int8(20), int8(8), int8(4), int8(13), int8(31), int8(22), int8(28), int8(18), int8(26), int8(10), int8(7), int8(12), int8(21), int8(17), int8(9), int8(6), int8(16), int8(5), int8(15), int8(14)}
	return int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&debruijn32)))) + uintptr(x&-x*uint32(124511785)>>27))))
}
func a_ctz_64(x uint64) int32 {
	var debruijn64 [64]int8 = [64]int8{int8(0), int8(1), int8(2), int8(53), int8(3), int8(7), int8(54), int8(27), int8(4), int8(38), int8(41), int8(8), int8(34), int8(55), int8(48), int8(28), int8(62), int8(5), int8(39), int8(46), int8(44), int8(42), int8(22), int8(9), int8(24), int8(35), int8(59), int8(56), int8(49), int8(18), int8(29), int8(11), int8(63), int8(52), int8(6), int8(26), int8(37), int8(40), int8(33), int8(47), int8(61), int8(45), int8(43), int8(21), int8(23), int8(58), int8(17), int8(10), int8(51), int8(25), int8(36), int8(32), int8(60), int8(20), int8(57), int8(16), int8(50), int8(31), int8(19), int8(15), int8(30), int8(14), int8(13), int8(12)}
	if false {
		var y uint32 = uint32(x)
		if !(y != 0) {
			y = uint32(x >> 32)
			return 32 + a_ctz_32(y)
		}
		return a_ctz_32(y)
	}
	return int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&debruijn64)))) + uintptr(x&-x*157587932685088877>>58))))
}
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
	if x>>32 != 0 {
		func() int32 {
			y = uint32(x >> 32)
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
	if y>>16 != 0 {
		y >>= 16
	} else {
		r |= int32(16)
	}
	if y>>8 != 0 {
		y >>= 8
	} else {
		r |= int32(8)
	}
	if y>>4 != 0 {
		y >>= 4
	} else {
		r |= int32(4)
	}
	if y>>2 != 0 {
		y >>= 2
	} else {
		r |= int32(2)
	}
	return r | func() int32 {
		if !(y>>1 != 0) {
			return 1
		} else {
			return 0
		}
	}()
}
func a_clz_32(x uint32) int32 {
	x >>= 1
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x++
	return 31 - a_ctz_32(x)
}

type _cgoa_42 struct {
	head    unsafe.Pointer
	off     int64
	pending unsafe.Pointer
}
type struct___pthread struct {
	self          *struct___pthread
	prev          *struct___pthread
	next          *struct___pthread
	sysinfo       uint64
	tid           int32
	errno_val     int32
	detach_state  int32
	cancel        int32
	canceldisable uint8
	cancelasync   uint8
	Xbf_0         uint8
	map_base      *uint8
	map_size      uint64
	stack         unsafe.Pointer
	stack_size    uint64
	guard_size    uint64
	result        unsafe.Pointer
	cancelbuf     *struct___ptcb
	tsd           *unsafe.Pointer
	robust_list   _cgoa_42
	h_errno_val   int32
	timer_id      int32
	locale        *struct___locale_struct
	killlock      [1]int32
	dlerror_buf   *int8
	stdio_locks   unsafe.Pointer
	canary        uint64
	dtv           *uint64
}

const (
	DT_EXITED   int32 = 0
	DT_EXITING  int32 = 1
	DT_JOINABLE int32 = 2
	DT_DETACHED int32 = 3
)

func __wake(addr unsafe.Pointer, cnt int32, priv int32) {
	if priv != 0 {
		priv = int32(128)
	}
	if cnt < 0 {
		cnt = int32(2147483647)
	}
	if !(__syscall3(int64(240), int64(uintptr(addr)), int64(1|priv), int64(cnt)) != int64(-38)) {
		__syscall3(int64(240), int64(uintptr(addr)), int64(1), int64(cnt))
	}
}
func __futexwait(addr unsafe.Pointer, val int32, priv int32) {
	if priv != 0 {
		priv = int32(128)
	}
	if !(__syscall4(int64(240), int64(uintptr(addr)), int64(0|priv), int64(val), int64(0)) != int64(-38)) {
		__syscall4(int64(240), int64(uintptr(addr)), int64(0), int64(val), int64(0))
	}
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
type _cgoa_46 struct {
	__ll int64
	__ld float64
}
type ptrdiff_t = int64



type wint_t = uint32
type wctype_t = uint64
type struct___mbstate_t struct {
	__opaque1 uint32
	__opaque2 uint32
}
type mbstate_t = struct___mbstate_t
type _cgoa_50 struct {
	quot int64
	rem  int64
}
