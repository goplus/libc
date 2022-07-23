package libc

import unsafe "unsafe"

var dummy_cgos__fflush *struct__IO_FILE = nil

func Fflush(f *struct__IO_FILE) int32 {
	if !(f != nil) {
		var r int32 = int32(0)
		if __stdout_used != nil {
			r |= Fflush(__stdout_used)
		}
		if __stderr_used != nil {
			r |= Fflush(__stderr_used)
		}
		for f = *__ofl_lock(); f != nil; f = f.next {
			var __need_unlock int32 = func() int32 {
				if f.lock >= int32(0) {
					return __lockfile(f)
				} else {
					return int32(0)
				}
			}()
			if uintptr(unsafe.Pointer(f.wpos)) != uintptr(unsafe.Pointer(f.wbase)) {
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
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if uintptr(unsafe.Pointer(f.wpos)) != uintptr(unsafe.Pointer(f.wbase)) {
		f.write(f, nil, uint64(0))
		if !(f.wpos != nil) {
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
	if uintptr(unsafe.Pointer(f.rpos)) != uintptr(unsafe.Pointer(f.rend)) {
		f.seek(f, int64(uintptr(unsafe.Pointer(f.rpos))-uintptr(unsafe.Pointer(f.rend))), int32(1))
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
	f.rpos = func() (_cgo_ret *uint8) {
		_cgo_addr := &f.rend
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
