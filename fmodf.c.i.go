package libc

import unsafe "unsafe"

func Fmodf(x float32, y float32) float32 {
	type _cgoa_15_fmodf struct {
		f float32
	}
	var ux _cgoa_15_fmodf
	ux.f = x
	var uy _cgoa_15_fmodf
	uy.f = y
	var ex int32 = int32(*(*uint32)(unsafe.Pointer(&ux)) >> int32(23) & uint32(255))
	var ey int32 = int32(*(*uint32)(unsafe.Pointer(&uy)) >> int32(23) & uint32(255))
	var sx uint32 = *(*uint32)(unsafe.Pointer(&ux)) & uint32(2147483648)
	var i uint32
	var uxi uint32 = *(*uint32)(unsafe.Pointer(&ux))
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
	if uxi<<int32(1) <= *(*uint32)(unsafe.Pointer(&uy))<<int32(1) {
		if uxi<<int32(1) == *(*uint32)(unsafe.Pointer(&uy))<<int32(1) {
			return float32(int32(0)) * x
		}
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
	for ; ex > ey; ex-- {
		i = uxi - *(*uint32)(unsafe.Pointer(&uy))
		if i>>int32(31) == uint32(0) {
			if i == uint32(0) {
				return float32(int32(0)) * x
			}
			uxi = i
		}
		uxi <<= int32(1)
	}
	i = uxi - *(*uint32)(unsafe.Pointer(&uy))
	if i>>int32(31) == uint32(0) {
		if i == uint32(0) {
			return float32(int32(0)) * x
		}
		uxi = i
	}
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
	if ex > int32(0) {
		uxi -= 8388608
		uxi |= uint32(ex) << int32(23)
	} else {
		uxi >>= -ex + int32(1)
	}
	uxi |= sx
	*(*uint32)(unsafe.Pointer(&ux)) = uxi
	return ux.f
}
