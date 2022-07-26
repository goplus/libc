package libc

import unsafe "unsafe"

func __towrite(f *Struct__IO_FILE) int32 {
	f.Mode |= f.Mode - int32(1)
	if f.Flags&uint32(8) != 0 {
		f.Flags |= uint32(32)
		return -1
	}
	f.Rpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.Rend
		*_cgo_addr = (*uint8)(nil)
		return *_cgo_addr
	}()
	f.Wpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.Wbase
		*_cgo_addr = f.Buf
		return *_cgo_addr
	}()
	f.Wend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.Buf)) + uintptr(f.Buf_size)))
	return int32(0)
}
func __towrite_needs_stdio_exit() {
	__stdio_exit_needed()
}
