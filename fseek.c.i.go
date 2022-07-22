package libc

import unsafe "unsafe"

func __fseeko_unlocked(f *struct__IO_FILE, off int64, whence int32) int32 {
	if whence == int32(1) && f.rend != nil {
		off -= int64(uintptr(unsafe.Pointer(f.rend)) - uintptr(unsafe.Pointer(f.rpos)))
	}
	if uintptr(unsafe.Pointer(f.wpos)) != uintptr(unsafe.Pointer(f.wbase)) {
		f.write(f, nil, uint64(0))
		if !(f.wpos != nil) {
			return -1
		}
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
	if f.seek(f, off, whence) < int64(0) {
		return -1
	}
	f.rpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.rend
		*_cgo_addr = (*uint8)(nil)
		return *_cgo_addr
	}()
	f.flags &= uint32(4294967279)
	return int32(0)
}
func __fseeko(f *struct__IO_FILE, off int64, whence int32) int32 {
	var result int32
	var __need_unlock int32 = func() int32 {
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	result = __fseeko_unlocked(f, off, whence)
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return result
}
func Fseek(f *struct__IO_FILE, off int64, whence int32) int32 {
	return __fseeko(f, int64(off), whence)
}
