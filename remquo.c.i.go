package libc

import unsafe "unsafe"

func Remquo(x float64, y float64, quo *int32) float64 {
	type _cgoa_18_remquo struct {
		f float64
	}
	var ux _cgoa_18_remquo
	ux.f = x
	var uy _cgoa_18_remquo
	uy.f = y
	var ex int32 = int32(*(*uint64)(unsafe.Pointer(&ux)) >> int32(52) & uint64(2047))
	var ey int32 = int32(*(*uint64)(unsafe.Pointer(&uy)) >> int32(52) & uint64(2047))
	var sx int32 = int32(*(*uint64)(unsafe.Pointer(&ux)) >> int32(63))
	var sy int32 = int32(*(*uint64)(unsafe.Pointer(&uy)) >> int32(63))
	var q uint32
	var i uint64
	var uxi uint64 = *(*uint64)(unsafe.Pointer(&ux))
	*quo = int32(0)
	if *(*uint64)(unsafe.Pointer(&uy))<<int32(1) == uint64(0) || func() int32 {
		if 8 == 4 {
			return func() int32 {
				if X__FLOAT_BITS(float32(y))&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 8 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(y)&9223372036854775807 > 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if X__fpclassifyl(float64(y)) == int32(0) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0 || ex == int32(2047) {
		return x * y / (x * y)
	}
	if *(*uint64)(unsafe.Pointer(&ux))<<int32(1) == uint64(0) {
		return x
	}
	if !(ex != 0) {
		for i = uxi << int32(12); i>>int32(63) == uint64(0); func() uint64 {
			ex--
			return func() (_cgo_ret uint64) {
				_cgo_addr := &i
				*_cgo_addr <<= int32(1)
				return *_cgo_addr
			}()
		}() {
		}
		uxi <<= -ex + int32(1)
	} else {
		uxi &= 4503599627370495
		uxi |= 4503599627370496
	}
	if !(ey != 0) {
		for i = *(*uint64)(unsafe.Pointer(&uy)) << int32(12); i>>int32(63) == uint64(0); func() uint64 {
			ey--
			return func() (_cgo_ret uint64) {
				_cgo_addr := &i
				*_cgo_addr <<= int32(1)
				return *_cgo_addr
			}()
		}() {
		}
		*(*uint64)(unsafe.Pointer(&uy)) <<= -ey + int32(1)
	} else {
		*(*uint64)(unsafe.Pointer(&uy)) &= 4503599627370495
		*(*uint64)(unsafe.Pointer(&uy)) |= 4503599627370496
	}
	q = uint32(0)
	if ex < ey {
		if ex+int32(1) == ey {
			goto end
		}
		return x
	}
	for ; ex > ey; ex-- {
		i = uxi - *(*uint64)(unsafe.Pointer(&uy))
		if i>>int32(63) == uint64(0) {
			uxi = i
			q++
		}
		uxi <<= int32(1)
		q <<= int32(1)
	}
	i = uxi - *(*uint64)(unsafe.Pointer(&uy))
	if i>>int32(63) == uint64(0) {
		uxi = i
		q++
	}
	if uxi == uint64(0) {
		ex = -60
	} else {
		for ; uxi>>int32(52) == uint64(0); func() int32 {
			uxi <<= int32(1)
			return func() (_cgo_ret int32) {
				_cgo_addr := &ex
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}()
		}() {
		}
	}
end:
	if ex > int32(0) {
		uxi -= 4503599627370496
		uxi |= uint64(ex) << int32(52)
	} else {
		uxi >>= -ex + int32(1)
	}
	*(*uint64)(unsafe.Pointer(&ux)) = uxi
	x = ux.f
	if sy != 0 {
		y = -y
	}
	if ex == ey || ex+int32(1) == ey && (float64(int32(2))*x > y || float64(int32(2))*x == y && q%uint32(2) != 0) {
		x -= y
		q++
	}
	q &= uint32(2147483647)
	*quo = func() int32 {
		if sx^sy != 0 {
			return -int32(q)
		} else {
			return int32(q)
		}
	}()
	return func() float64 {
		if sx != 0 {
			return -x
		} else {
			return x
		}
	}()
}
