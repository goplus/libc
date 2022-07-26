package libc

import unsafe "unsafe"

func Putchar_unlocked(c int32) int32 {
	return func() int32 {
		if int32(uint8(c)) != (&__stdout_FILE).Lbf && uintptr(unsafe.Pointer((&__stdout_FILE).Wpos)) != uintptr(unsafe.Pointer((&__stdout_FILE).Wend)) {
			return int32(func() (_cgo_ret uint8) {
				_cgo_addr := &*func() (_cgo_ret *uint8) {
					_cgo_addr := &(&__stdout_FILE).Wpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
				*_cgo_addr = uint8(c)
				return *_cgo_addr
			}())
		} else {
			return __overflow(&__stdout_FILE, int32(uint8(c)))
		}
	}()
}
