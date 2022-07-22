package libc

import unsafe "unsafe"

func Fgets(s *int8, n int32, f *struct__IO_FILE) *int8 {
	var p *int8 = s
	var z *uint8
	var k uint64
	var c int32
	var __need_unlock int32 = func() int32 {
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if func() (_cgo_ret int32) {
		_cgo_addr := &n
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() <= int32(1) {
		f.mode |= f.mode - int32(1)
		for {
			if __need_unlock != 0 {
				__unlockfile(f)
			}
			if true {
				break
			}
		}
		if n != 0 {
			return (*int8)(nil)
		}
		*s = int8(0)
		return s
	}
	for n != 0 {
		if uintptr(unsafe.Pointer(f.rpos)) != uintptr(unsafe.Pointer(f.rend)) {
			z = (*uint8)(Memchr(unsafe.Pointer(f.rpos), '\n', uint64(uintptr(unsafe.Pointer(f.rend))-uintptr(unsafe.Pointer(f.rpos)))))
			k = uint64(func() int64 {
				if z != nil {
					return int64(uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(f.rpos)) + uintptr(int64(1)))
				} else {
					return int64(uintptr(unsafe.Pointer(f.rend)) - uintptr(unsafe.Pointer(f.rpos)))
				}
			}())
			k = func() uint64 {
				if k < uint64(n) {
					return k
				} else {
					return uint64(n)
				}
			}()
			Memcpy(unsafe.Pointer(p), unsafe.Pointer(f.rpos), k)
			*(*uintptr)(unsafe.Pointer(&f.rpos)) += uintptr(k)
			*(*uintptr)(unsafe.Pointer(&p)) += uintptr(k)
			n -= int32(k)
			if z != nil || !(n != 0) {
				break
			}
		}
		if func() (_cgo_ret int32) {
			_cgo_addr := &c
			*_cgo_addr = func() int32 {
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
			}()
			return *_cgo_addr
		}() < int32(0) {
			if uintptr(unsafe.Pointer(p)) == uintptr(unsafe.Pointer(s)) || !(f.flags&uint32(16) != 0) {
				s = (*int8)(nil)
			}
			break
		}
		n--
		if int32(func() (_cgo_ret int8) {
			_cgo_addr := &*func() (_cgo_ret *int8) {
				_cgo_addr := &p
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
			*_cgo_addr = int8(c)
			return *_cgo_addr
		}()) == '\n' {
			break
		}
	}
	if s != nil {
		*p = int8(0)
	}
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return s
}
