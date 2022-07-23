package libc

import unsafe "unsafe"

func Ilogbf(x float32) int32 {
	type _cgoa_18_ilogbf struct {
		f float32
	}
	var u _cgoa_18_ilogbf
	u.f = x
	var i uint32 = *(*uint32)(unsafe.Pointer(&u))
	var e int32 = int32(i >> int32(23) & uint32(255))
	if !(e != 0) {
		i <<= int32(9)
		if i == uint32(0) {
			for {
				if 4 == 4 {
					fp_force_evalf(func() float32 {
						return float32(int32(0))
					}() / 0)
				} else if 4 == 8 {
					fp_force_eval(float64(func() float32 {
						return float32(int32(0))
					}() / 0))
				} else {
					fp_force_evall(float64(func() float32 {
						return float32(int32(0))
					}() / 0))
				}
				if true {
					break
				}
			}
			return -2147483648
		}
		for e = -127; i>>int32(31) == uint32(0); func() uint32 {
			e--
			return func() (_cgo_ret uint32) {
				_cgo_addr := &i
				*_cgo_addr <<= int32(1)
				return *_cgo_addr
			}()
		}() {
		}
		return e
	}
	if e == int32(255) {
		for {
			if 4 == 4 {
				fp_force_evalf(func() float32 {
					return float32(int32(0))
				}() / 0)
			} else if 4 == 8 {
				fp_force_eval(float64(func() float32 {
					return float32(int32(0))
				}() / 0))
			} else {
				fp_force_evall(float64(func() float32 {
					return float32(int32(0))
				}() / 0))
			}
			if true {
				break
			}
		}
		return func() int32 {
			if i<<int32(9) != 0 {
				return -2147483648
			} else {
				return int32(2147483647)
			}
		}()
	}
	return e - int32(127)
}
