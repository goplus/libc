package libc

import unsafe "unsafe"

func Fgetln(f *Struct__IO_FILE, plen *uint64) *int8 {
	var ret *int8 = nil
	var z *int8
	var l int64
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	Ungetc(func() int32 {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __uflow(f)
		}
	}(), f)
	if f.Rend != nil && func() (_cgo_ret *int8) {
		_cgo_addr := &z
		*_cgo_addr = (*int8)(Memchr(unsafe.Pointer(f.Rpos), '\n', uint64(uintptr(unsafe.Pointer(f.Rend))-uintptr(unsafe.Pointer(f.Rpos)))))
		return *_cgo_addr
	}() != nil {
		ret = (*int8)(unsafe.Pointer(f.Rpos))
		*plen = uint64(uintptr(unsafe.Pointer(func() (_cgo_ret *int8) {
			_cgo_addr := &z
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return *_cgo_addr
		}())) - uintptr(unsafe.Pointer(ret)))
		f.Rpos = (*uint8)(unsafe.Pointer(z))
	} else if func() (_cgo_ret int64) {
		_cgo_addr := &l
		*_cgo_addr = Getline(&f.Getln_buf, (*uint64)(unsafe.Pointer(&[1]uint64{uint64(0)})), f)
		return *_cgo_addr
	}() > int64(0) {
		*plen = uint64(l)
		ret = f.Getln_buf
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
