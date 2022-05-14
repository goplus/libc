package libc

import unsafe "unsafe"

func __towrite(f *struct__IO_FILE) int32 {
	f.mode |= f.mode - int32(1)
	if f.flags&uint32(8) != 0 {
		f.flags |= uint32(32)
		return -1
	}
	f.rpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.rend
		*_cgo_addr = (*uint8)(nil)
		return *_cgo_addr
	}()
	f.wpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.wbase
		*_cgo_addr = f.buf
		return *_cgo_addr
	}()
	f.wend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.buf)) + uintptr(f.buf_size)))
	return int32(0)
}
func __towrite_needs_stdio_exit() {
	__stdio_exit_needed()
}
