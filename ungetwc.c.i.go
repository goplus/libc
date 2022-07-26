package libc

import unsafe "unsafe"

func ungetwc(c uint32, f *Struct__IO_FILE) uint32 {
	var mbc [4]uint8
	var l int32
	var ploc **Struct___locale_struct = &__pthread_self().Locale
	var loc *Struct___locale_struct = *ploc
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if f.Mode <= int32(0) {
		fwide(f, int32(1))
	}
	*ploc = f.Locale
	if !(f.Rpos != nil) {
		__toread(f)
	}
	if !(f.Rpos != nil) || c == uint32(4294967295) || func() (_cgo_ret int32) {
		_cgo_addr := &l
		*_cgo_addr = int32(wcrtomb((*int8)(unsafe.Pointer((*uint8)(unsafe.Pointer(&mbc)))), c, nil))
		return *_cgo_addr
	}() < int32(0) || uintptr(unsafe.Pointer(f.Rpos)) < uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.Buf))-uintptr(int32(8))))))+uintptr(l))))) {
		for {
			if __need_unlock != 0 {
				__unlockfile(f)
			}
			if true {
				break
			}
		}
		*ploc = loc
		return uint32(4294967295)
	}
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
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &f.Rpos
			*(*uintptr)(unsafe.Pointer(_cgo_addr))--
			return *_cgo_addr
		}() = uint8(c)
	} else {
		Memcpy(unsafe.Pointer(func() (_cgo_ret *uint8) {
			_cgo_addr := &f.Rpos
			*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) -= uintptr(l)
			return *_cgo_addr
		}()), unsafe.Pointer((*uint8)(unsafe.Pointer(&mbc))), uint64(l))
	}
	f.Flags &= uint32(4294967279)
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	*ploc = loc
	return c
}
