package libc

import unsafe "unsafe"

func ungetwc(c uint32, f *struct__IO_FILE) uint32 {
	var mbc [4]uint8
	var l int32
	var ploc **struct___locale_struct = &__pthread_self().locale
	var loc *struct___locale_struct = *ploc
	var __need_unlock int32 = func() int32 {
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if f.mode <= int32(0) {
		fwide(f, int32(1))
	}
	*ploc = f.locale
	if !(f.rpos != nil) {
		__toread(f)
	}
	if !(f.rpos != nil) || c == uint32(4294967295) || func() (_cgo_ret int32) {
		_cgo_addr := &l
		*_cgo_addr = int32(wcrtomb((*int8)(unsafe.Pointer((*uint8)(unsafe.Pointer(&mbc)))), c, nil))
		return *_cgo_addr
	}() < int32(0) || uintptr(unsafe.Pointer(f.rpos)) < uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.buf))-uintptr(int32(8))))))+uintptr(l))))) {
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
			_cgo_addr := &f.rpos
			*(*uintptr)(unsafe.Pointer(_cgo_addr))--
			return *_cgo_addr
		}() = uint8(c)
	} else {
		Memcpy(unsafe.Pointer(func() (_cgo_ret *uint8) {
			_cgo_addr := &f.rpos
			*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) -= uintptr(l)
			return *_cgo_addr
		}()), unsafe.Pointer((*uint8)(unsafe.Pointer(&mbc))), uint64(l))
	}
	f.flags &= uint32(4294967279)
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
