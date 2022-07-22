package libc

import unsafe "unsafe"

func __toread(f *struct__IO_FILE) int32 {
	f.mode |= f.mode - int32(1)
	if uintptr(unsafe.Pointer(f.wpos)) != uintptr(unsafe.Pointer(f.wbase)) {
		f.write(f, nil, uint64(0))
	}
	f.wpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.wbase
		*_cgo_addr = func() (_cgo_ret *uint8) {
			_cgo_addr := &f.wend
			*_cgo_addr = (*uint8)(nil)
			return *_cgo_addr
		}()
		return *_cgo_addr
	}()
	if f.flags&uint32(4) != 0 {
		f.flags |= uint32(32)
		return -1
	}
	f.rpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.rend
		*_cgo_addr = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.buf)) + uintptr(f.buf_size)))
		return *_cgo_addr
	}()
	return func() int32 {
		if f.flags&uint32(16) != 0 {
			return -1
		} else {
			return int32(0)
		}
	}()
}
func __toread_needs_stdio_exit() {
	__stdio_exit_needed()
}
