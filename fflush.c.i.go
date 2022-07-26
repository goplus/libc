package libc

import unsafe "unsafe"

var _cgos_dummy_fflush *Struct__IO_FILE = nil

func Fflush(f *Struct__IO_FILE) int32 {
	if !(f != nil) {
		var r int32 = int32(0)
		if __stdout_used != nil {
			r |= Fflush(__stdout_used)
		}
		if __stderr_used != nil {
			r |= Fflush(__stderr_used)
		}
		for f = *__ofl_lock(); f != nil; f = f.Next {
			var __need_unlock int32 = func() int32 {
				if f.Lock >= int32(0) {
					return __lockfile(f)
				} else {
					return int32(0)
				}
			}()
			if uintptr(unsafe.Pointer(f.Wpos)) != uintptr(unsafe.Pointer(f.Wbase)) {
				r |= Fflush(f)
			}
			for {
				if __need_unlock != 0 {
					__unlockfile(f)
				}
				if true {
					break
				}
			}
		}
		__ofl_unlock()
		return r
	}
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if uintptr(unsafe.Pointer(f.Wpos)) != uintptr(unsafe.Pointer(f.Wbase)) {
		f.Write(f, nil, uint64(0))
		if !(f.Wpos != nil) {
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
	}
	if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) {
		f.Seek(f, int64(uintptr(unsafe.Pointer(f.Rpos))-uintptr(unsafe.Pointer(f.Rend))), int32(1))
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
	f.Rpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.Rend
		*_cgo_addr = (*uint8)(nil)
		return *_cgo_addr
	}()
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return int32(0)
}
