package libc

import unsafe "unsafe"

func Getdelim(s **int8, n *uint64, delim int32, f *Struct__IO_FILE) int64 {
	var tmp *int8
	var z *uint8
	var k uint64
	var i uint64 = uint64(0)
	var c int32
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if !(n != nil) || !(s != nil) {
		f.Mode |= f.Mode - int32(1)
		f.Flags |= uint32(32)
		for {
			if __need_unlock != 0 {
				__unlockfile(f)
			}
			if true {
				break
			}
		}
		*X__errno_location() = int32(22)
		return int64(-1)
	}
	if !(*s != nil) {
		*n = uint64(0)
	}
	for {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) {
			z = (*uint8)(Memchr(unsafe.Pointer(f.Rpos), delim, uint64(uintptr(unsafe.Pointer(f.Rend))-uintptr(unsafe.Pointer(f.Rpos)))))
			k = uint64(func() int64 {
				if z != nil {
					return int64(uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(f.Rpos)) + uintptr(int64(1)))
				} else {
					return int64(uintptr(unsafe.Pointer(f.Rend)) - uintptr(unsafe.Pointer(f.Rpos)))
				}
			}())
		} else {
			z = (*uint8)(nil)
			k = uint64(0)
		}
		if i+k >= *n {
			var m uint64 = i + k + uint64(2)
			if !(z != nil) && m < uint64(1073741823) {
				m += m / uint64(2)
			}
			tmp = (*int8)(Realloc(unsafe.Pointer(*s), m))
			if !(tmp != nil) {
				m = i + k + uint64(2)
				tmp = (*int8)(Realloc(unsafe.Pointer(*s), m))
				if !(tmp != nil) {
					k = *n - i
					Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*s))+uintptr(i)))), unsafe.Pointer(f.Rpos), k)
					*(*uintptr)(unsafe.Pointer(&f.Rpos)) += uintptr(k)
					f.Mode |= f.Mode - int32(1)
					f.Flags |= uint32(32)
					for {
						if __need_unlock != 0 {
							__unlockfile(f)
						}
						if true {
							break
						}
					}
					*X__errno_location() = int32(12)
					return int64(-1)
				}
			}
			*s = tmp
			*n = m
		}
		Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*s))+uintptr(i)))), unsafe.Pointer(f.Rpos), k)
		*(*uintptr)(unsafe.Pointer(&f.Rpos)) += uintptr(k)
		i += k
		if z != nil {
			break
		}
		if func() (_cgo_ret int32) {
			_cgo_addr := &c
			*_cgo_addr = func() int32 {
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
			}()
			return *_cgo_addr
		}() == -1 {
			if !(i != 0) || !(f.Flags&uint32(16) != 0) {
				for {
					if __need_unlock != 0 {
						__unlockfile(f)
					}
					if true {
						break
					}
				}
				return int64(-1)
			}
			break
		}
		if i+uint64(1) >= *n {
			*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return *_cgo_addr
			}() = uint8(c)
		} else if int32(func() (_cgo_ret int8) {
			_cgo_addr := &*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*s)) + uintptr(func() (_cgo_ret uint64) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}())))
			*_cgo_addr = int8(c)
			return *_cgo_addr
		}()) == delim {
			break
		}
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*s)) + uintptr(i))) = int8(0)
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return int64(i)
}
