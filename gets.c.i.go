package libc

import unsafe "unsafe"

func gets(s *int8) *int8 {
	var i uint64 = uint64(0)
	var c int32
	var __need_unlock int32 = func() int32 {
		if (&__stdin_FILE).Lock >= int32(0) {
			return __lockfile(&__stdin_FILE)
		} else {
			return int32(0)
		}
	}()
	for func() (_cgo_ret int32) {
		_cgo_addr := &c
		*_cgo_addr = func() int32 {
			if uintptr(unsafe.Pointer((&__stdin_FILE).Rpos)) != uintptr(unsafe.Pointer((&__stdin_FILE).Rend)) {
				return int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &(&__stdin_FILE).Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return __uflow(&__stdin_FILE)
			}
		}()
		return *_cgo_addr
	}() != -1 && c != '\n' {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(func() (_cgo_ret uint64) {
			_cgo_addr := &i
			_cgo_ret = *_cgo_addr
			*_cgo_addr++
			return
		}()))) = int8(c)
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i))) = int8(0)
	if c != '\n' && (!((&__stdin_FILE).Flags&uint32(16) != 0) || !(i != 0)) {
		s = (*int8)(nil)
	}
	for {
		if __need_unlock != 0 {
			__unlockfile(&__stdin_FILE)
		}
		if true {
			break
		}
	}
	return s
}
