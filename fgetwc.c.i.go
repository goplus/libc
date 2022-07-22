package libc

import unsafe "unsafe"

func __fgetwc_unlocked_internal_cgo18_fgetwc(f *struct__IO_FILE) uint32 {
	var wc uint32
	var c int32
	var l uint64
	if uintptr(unsafe.Pointer(f.rpos)) != uintptr(unsafe.Pointer(f.rend)) {
		l = uint64(Mbtowc(&wc, (*int8)(unsafe.Pointer(f.rpos)), uint64(uintptr(unsafe.Pointer(f.rend))-uintptr(unsafe.Pointer(f.rpos)))))
		if l+uint64(1) >= uint64(1) {
			*(*uintptr)(unsafe.Pointer(&f.rpos)) += uintptr(l + func() uint64 {
				if !(l != 0) {
					return 1
				} else {
					return 0
				}
			}())
			return wc
		}
	}
	var st struct___mbstate_t = struct___mbstate_t{uint32(0), 0}
	var b uint8
	var first int32 = int32(1)
	for {
		b = uint8(func() (_cgo_ret int32) {
			_cgo_addr := &c
			*_cgo_addr = func() int32 {
				if uintptr(unsafe.Pointer(f.rpos)) != uintptr(unsafe.Pointer(f.rend)) {
					return int32(*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}())
				} else {
					return __uflow(f)
				}
			}()
			return *_cgo_addr
		}())
		if c < int32(0) {
			if !(first != 0) {
				f.flags |= uint32(32)
				*__errno_location() = int32(84)
			}
			return uint32(4294967295)
		}
		l = mbrtowc(&wc, (*int8)(unsafe.Pointer(&b)), uint64(1), &st)
		if l == uint64(18446744073709551615) {
			if !(first != 0) {
				f.flags |= uint32(32)
				Ungetc(int32(b), f)
			}
			return uint32(4294967295)
		}
		first = int32(0)
		if !(l == uint64(18446744073709551614)) {
			break
		}
	}
	return wc
}
func __fgetwc_unlocked(f *struct__IO_FILE) uint32 {
	var ploc **struct___locale_struct = &__pthread_self().locale
	var loc *struct___locale_struct = *ploc
	if f.mode <= int32(0) {
		fwide(f, int32(1))
	}
	*ploc = f.locale
	var wc uint32 = __fgetwc_unlocked_internal_cgo18_fgetwc(f)
	*ploc = loc
	return wc
}
func fgetwc(f *struct__IO_FILE) uint32 {
	var c uint32
	var __need_unlock int32 = func() int32 {
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	c = __fgetwc_unlocked(f)
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
