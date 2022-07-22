package libc

import unsafe "unsafe"

func __fputwc_unlocked(c uint32, f *struct__IO_FILE) uint32 {
	var mbc [4]int8
	var l int32
	var ploc **struct___locale_struct = &__pthread_self().locale
	var loc *struct___locale_struct = *ploc
	if f.mode <= int32(0) {
		fwide(f, int32(1))
	}
	*ploc = f.locale
	if func() int32 {
		if int32(0) != 0 {
			return Isascii(int32(c))
		} else {
			return func() int32 {
				if uint32(c) < uint32(128) {
					return 1
				} else {
					return 0
				}
			}()
		}
	}() != 0 {
		c = uint32(func() int32 {
			if int32(uint8(c)) != f.lbf && uintptr(unsafe.Pointer(f.wpos)) != uintptr(unsafe.Pointer(f.wend)) {
				return int32(func() (_cgo_ret uint8) {
					_cgo_addr := &*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.wpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}()
					*_cgo_addr = uint8(c)
					return *_cgo_addr
				}())
			} else {
				return __overflow(f, int32(uint8(c)))
			}
		}())
	} else if uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.wpos))+uintptr(int32(4)))))) < uintptr(unsafe.Pointer(f.wend)) {
		l = Wctomb((*int8)(unsafe.Pointer(f.wpos)), c)
		if l < int32(0) {
			c = uint32(4294967295)
		} else {
			*(*uintptr)(unsafe.Pointer(&f.wpos)) += uintptr(l)
		}
	} else {
		l = Wctomb((*int8)(unsafe.Pointer(&mbc)), c)
		if l < int32(0) || __fwritex((*uint8)(unsafe.Pointer((*int8)(unsafe.Pointer(&mbc)))), uint64(l), f) < uint64(l) {
			c = uint32(4294967295)
		}
	}
	if c == uint32(4294967295) {
		f.flags |= uint32(32)
	}
	*ploc = loc
	return c
}
func fputwc(c uint32, f *struct__IO_FILE) uint32 {
	var __need_unlock int32 = func() int32 {
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	c = __fputwc_unlocked(c, f)
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return c
}
