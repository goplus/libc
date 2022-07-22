package libc

import unsafe "unsafe"

func Fgetln(f *struct__IO_FILE, plen *uint64) *int8 {
	var ret *int8 = nil
	var z *int8
	var l int64
	var __need_unlock int32 = func() int32 {
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	Ungetc(func() int32 {
		if uintptr(unsafe.Pointer(f.rpos)) != uintptr(unsafe.Pointer(f.rend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __uflow(f)
		}
	}(), f)
	if f.rend != nil && func() (_cgo_ret *int8) {
		_cgo_addr := &z
		*_cgo_addr = (*int8)(Memchr(unsafe.Pointer(f.rpos), '\n', uint64(uintptr(unsafe.Pointer(f.rend))-uintptr(unsafe.Pointer(f.rpos)))))
		return *_cgo_addr
	}() != nil {
		ret = (*int8)(unsafe.Pointer(f.rpos))
		*plen = uint64(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
			_cgo_addr := &z
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return *_cgo_addr
		}())) - uintptr(unsafe.Pointer(ret)))
		f.rpos = (*uint8)(unsafe.Pointer(z))
	} else if func() (_cgo_ret int64) {
		_cgo_addr := &l
		*_cgo_addr = Getline(&f.getln_buf, (*uint64)(unsafe.Pointer(&[1]uint64{uint64(0)})), f)
		return *_cgo_addr
	}() > int64(0) {
		*plen = uint64(l)
		ret = f.getln_buf
	}
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return ret
}
