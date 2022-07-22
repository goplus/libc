package libc

import unsafe "unsafe"

func Ungetc(c int32, f *struct__IO_FILE) int32 {
	if c == -1 {
		return c
	}
	var __need_unlock int32 = func() int32 {
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if !(f.rpos != nil) {
		__toread(f)
	}
	if !(f.rpos != nil) || uintptr(unsafe.Pointer(f.rpos)) <= uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.buf))-uintptr(int32(8)))))) {
		for {
			if __need_unlock != 0 {
				__unlockfile(f)
			}
			if true {
				break
			}
		}
		return -1
	}
	*func() (_cgo_ret *uint8) {
		_cgo_addr := &f.rpos
		*(*uintptr)(unsafe.Pointer(_cgo_addr))--
		return *_cgo_addr
	}() = uint8(c)
	f.flags &= uint32(4294967279)
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return int32(uint8(c))
}
