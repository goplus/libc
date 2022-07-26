package libc

import unsafe "unsafe"

func Getc_unlocked(f *Struct__IO_FILE) int32 {
	return func() int32 {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __uflow(f)
		}
	}()
}
