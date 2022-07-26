package libc

import unsafe "unsafe"

func __fseeko_unlocked(f *Struct__IO_FILE, off int64, whence int32) int32 {
	if whence == int32(1) && f.Rend != nil {
		off -= int64(uintptr(unsafe.Pointer(f.Rend)) - uintptr(unsafe.Pointer(f.Rpos)))
	}
	if uintptr(unsafe.Pointer(f.Wpos)) != uintptr(unsafe.Pointer(f.Wbase)) {
		f.Write(f, nil, uint64(0))
		if !(f.Wpos != nil) {
			return -1
		}
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
	if f.Seek(f, off, whence) < int64(0) {
		return -1
	}
	f.Rpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.Rend
		*_cgo_addr = (*uint8)(nil)
		return *_cgo_addr
	}()
	f.Flags &= uint32(4294967279)
	return int32(0)
}
func __fseeko(f *Struct__IO_FILE, off int64, whence int32) int32 {
	var result int32
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
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
func Fseek(f *Struct__IO_FILE, off int64, whence int32) int32 {
	return __fseeko(f, int64(off), whence)
}
