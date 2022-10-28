package libc

import unsafe "unsafe"

func Ilogb(x float64) int32 {
	type _cgoa_18_ilogb struct {
		f float64
	}
	var u _cgoa_18_ilogb
	u.f = x
	var i uint64 = *(*uint64)(unsafe.Pointer(&u))
	var e int32 = int32(i >> int32(52) & uint64(2047))
	if !(e != 0) {
		i <<= int32(12)
		if i == uint64(0) {
			for {
				if 4 == 4 {
					fp_force_evalf(func() float32 {
						return float32(int32(0))
					}() / 0.0)
				} else if 4 == 8 {
					fp_force_eval(float64(func() float32 {
						return float32(int32(0))
					}() / 0.0))
				} else {
					fp_force_evall(float64(func() float32 {
						return float32(int32(0))
					}() / 0.0))
				}
				if true {
					break
				}
			}
			return -2147483648
		}
		for e = -1023; i>>int32(63) == uint64(0); func() uint64 {
			e--
			return func() (_cgo_ret uint64) {
				_cgo_addr := &i
				*_cgo_addr <<= int32(1)
				return *_cgo_addr
			}()
		}() {
		}
		return e
	}
	if e == int32(2047) {
		for {
			if 4 == 4 {
				fp_force_evalf(func() float32 {
					return float32(int32(0))
				}() / 0.0)
			} else if 4 == 8 {
				fp_force_eval(float64(func() float32 {
					return float32(int32(0))
				}() / 0.0))
			} else {
				fp_force_evall(float64(func() float32 {
					return float32(int32(0))
				}() / 0.0))
			}
			if true {
				break
			}
		}
		return func() int32 {
			if i<<int32(12) != 0 {
				return -2147483648
			} else {
				return int32(2147483647)
			}
		}()
	}
	return e - int32(1023)
}
