package libc

import unsafe "unsafe"

func __toread(f *Struct__IO_FILE) int32 {
	f.Mode |= f.Mode - int32(1)
	if uintptr(unsafe.Pointer(f.Wpos)) != uintptr(unsafe.Pointer(f.Wbase)) {
		f.Write(f, nil, uint64(0))
	}
	f.Wpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.Wbase
		*_cgo_addr = func() (_cgo_ret *uint8) {
			_cgo_addr := &f.Wend
			*_cgo_addr = (*uint8)(nil)
			return *_cgo_addr
		}()
		return *_cgo_addr
	}()
	if f.Flags&uint32(4) != 0 {
		f.Flags |= uint32(32)
		return -1
	}
	f.Rpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.Rend
		*_cgo_addr = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.Buf)) + uintptr(f.Buf_size)))
		return *_cgo_addr
	}()
	return func() int32 {
		if f.Flags&uint32(16) != 0 {
			return -1
		} else {
			return int32(0)
		}
	}()
}
func __toread_needs_stdio_exit() {
	__stdio_exit_needed()
}
