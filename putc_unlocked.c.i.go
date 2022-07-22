package libc

import unsafe "unsafe"

func Putc_unlocked(c int32, f *struct__IO_FILE) int32 {
	return func() int32 {
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
	}()
}
