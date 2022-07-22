package libc

import unsafe "unsafe"

func Remquof(x float32, y float32, quo *int32) float32 {
	type _cgoa_18_remquof struct {
		f float32
	}
	var ux _cgoa_18_remquof
	ux.f = x
	var uy _cgoa_18_remquof
	uy.f = y
	var ex int32 = int32(*(*uint32)(unsafe.Pointer(&ux)) >> int32(23) & uint32(255))
	var ey int32 = int32(*(*uint32)(unsafe.Pointer(&uy)) >> int32(23) & uint32(255))
	var sx int32 = int32(*(*uint32)(unsafe.Pointer(&ux)) >> int32(31))
	var sy int32 = int32(*(*uint32)(unsafe.Pointer(&uy)) >> int32(31))
	var q uint32
	var i uint32
	var uxi uint32 = *(*uint32)(unsafe.Pointer(&ux))
	*quo = int32(0)
	if *(*uint32)(unsafe.Pointer(&uy))<<int32(1) == uint32(0) || func() int32 {
		if 4 == 4 {
			return func() int32 {
				if X__FLOAT_BITS(y)&uint32(2147483647) > uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 4 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(float64(y))&9223372036854775807 > 9218868437227405312 {
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
	}() != 0 || ex == int32(255) {
		return x * y / (x * y)
	}
	if *(*uint32)(unsafe.Pointer(&ux))<<int32(1) == uint32(0) {
		return x
	}
	if !(ex != 0) {
		for i = uxi << int32(9); i>>int32(31) == uint32(0); func() uint32 {
			ex--
			return func() (_cgo_ret uint32) {
				_cgo_addr := &i
				*_cgo_addr <<= int32(1)
				return *_cgo_addr
			}()
		}() {
		}
		uxi <<= -ex + int32(1)
	} else {
		uxi &= 8388607
		uxi |= 8388608
	}
	if !(ey != 0) {
		for i = *(*uint32)(unsafe.Pointer(&uy)) << int32(9); i>>int32(31) == uint32(0); func() uint32 {
			ey--
			return func() (_cgo_ret uint32) {
				_cgo_addr := &i
				*_cgo_addr <<= int32(1)
				return *_cgo_addr
			}()
		}() {
		}
		*(*uint32)(unsafe.Pointer(&uy)) <<= -ey + int32(1)
	} else {
		*(*uint32)(unsafe.Pointer(&uy)) &= 8388607
		*(*uint32)(unsafe.Pointer(&uy)) |= 8388608
	}
	q = uint32(0)
	if ex < ey {
		if ex+int32(1) == ey {
			goto end
		}
		return x
	}
	for ; ex > ey; ex-- {
		i = uxi - *(*uint32)(unsafe.Pointer(&uy))
		if i>>int32(31) == uint32(0) {
			uxi = i
			q++
		}
		uxi <<= int32(1)
		q <<= int32(1)
	}
	i = uxi - *(*uint32)(unsafe.Pointer(&uy))
	if i>>int32(31) == uint32(0) {
		uxi = i
		q++
	}
	if uxi == uint32(0) {
		ex = -30
	} else {
		for ; uxi>>int32(23) == uint32(0); func() int32 {
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
		uxi -= 8388608
		uxi |= uint32(ex) << int32(23)
	} else {
		uxi >>= -ex + int32(1)
	}
	*(*uint32)(unsafe.Pointer(&ux)) = uxi
	x = ux.f
	if sy != 0 {
		y = -y
	}
	if ex == ey || ex+int32(1) == ey && (float32(int32(2))*x > y || float32(int32(2))*x == y && q%uint32(2) != 0) {
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
	return func() float32 {
		if sx != 0 {
			return -x
		} else {
			return x
		}
	}()
}
