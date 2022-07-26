package libc

import unsafe "unsafe"

func Ungetc(c int32, f *Struct__IO_FILE) int32 {
	if c == -1 {
		return c
	}
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if !(f.Rpos != nil) {
		__toread(f)
	}
	if !(f.Rpos != nil) || uintptr(unsafe.Pointer(f.Rpos)) <= uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.Buf))-uintptr(int32(8)))))) {
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
		_cgo_addr := &f.Rpos
		*(*uintptr)(unsafe.Pointer(_cgo_addr))--
		return *_cgo_addr
	}() = uint8(c)
	f.Flags &= uint32(4294967279)
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
