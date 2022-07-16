package libc

import unsafe "unsafe"

func Fmod(x float64, y float64) float64 {
	type _cgoa_15_fmod struct {
		f float64
	}
	var ux _cgoa_15_fmod
	ux.f = x
	var uy _cgoa_15_fmod
	uy.f = y
	var ex int32 = int32(*(*uint64)(unsafe.Pointer(&ux)) >> int32(52) & uint64(2047))
	var ey int32 = int32(*(*uint64)(unsafe.Pointer(&uy)) >> int32(52) & uint64(2047))
	var sx int32 = int32(*(*uint64)(unsafe.Pointer(&ux)) >> int32(63))
	var i uint64
	var uxi uint64 = *(*uint64)(unsafe.Pointer(&ux))
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
	if uxi<<int32(1) <= *(*uint64)(unsafe.Pointer(&uy))<<int32(1) {
		if uxi<<int32(1) == *(*uint64)(unsafe.Pointer(&uy))<<int32(1) {
			return float64(int32(0)) * x
		}
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
	for ; ex > ey; ex-- {
		i = uxi - *(*uint64)(unsafe.Pointer(&uy))
		if i>>int32(63) == uint64(0) {
			if i == uint64(0) {
				return float64(int32(0)) * x
			}
			uxi = i
		}
		uxi <<= int32(1)
	}
	i = uxi - *(*uint64)(unsafe.Pointer(&uy))
	if i>>int32(63) == uint64(0) {
		if i == uint64(0) {
			return float64(int32(0)) * x
		}
		uxi = i
	}
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
	if ex > int32(0) {
		uxi -= 4503599627370496
		uxi |= uint64(ex) << int32(52)
	} else {
		uxi >>= -ex + int32(1)
	}
	uxi |= uint64(sx) << int32(63)
	*(*uint64)(unsafe.Pointer(&ux)) = uxi
	return ux.f
}
