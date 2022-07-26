package libc

import unsafe "unsafe"

func Getchar_unlocked() int32 {
	return func() int32 {
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
}
