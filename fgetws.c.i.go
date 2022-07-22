package libc

import unsafe "unsafe"

func fgetws(s *uint32, n int32, f *struct__IO_FILE) *uint32 {
	var p *uint32 = s
	if !(func() (_cgo_ret int32) {
		_cgo_addr := &n
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0) {
		return s
	}
	var __need_unlock int32 = func() int32 {
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	*__errno_location() = int32(11)
	for ; n != 0; n-- {
		var c uint32 = __fgetwc_unlocked(f)
		if c == uint32(4294967295) {
			break
		}
		*func() (_cgo_ret *uint32) {
			_cgo_addr := &p
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}() = c
		if c == uint32('\n') {
			break
		}
	}
	*p = uint32(0)
	if f.flags&uint32(32) != 0 || *__errno_location() == int32(84) {
		p = s
	}
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return func() *uint32 {
		if uintptr(unsafe.Pointer(p)) == uintptr(unsafe.Pointer(s)) {
			return nil
		} else {
			return s
		}
	}()
}
