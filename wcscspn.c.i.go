package libc

import unsafe "unsafe"

func wcscspn(s *uint32, c *uint32) uint64 {
	var a *uint32
	if !(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(int32(0))*4)) != 0) {
		return wcslen(s)
	}
	if !(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(int32(1))*4)) != 0) {
		return func() uint64 {
			if func() (_cgo_ret *uint32) {
				_cgo_addr := &s
				*_cgo_addr = wcschr(func() (_cgo_ret *uint32) {
					_cgo_addr := &a
					*_cgo_addr = s
					return *_cgo_addr
				}(), *c)
				return *_cgo_addr
			}() != nil {
				return uint64((uintptr(unsafe.Pointer(s)) - uintptr(unsafe.Pointer(a))) / 4)
			} else {
				return wcslen(a)
			}
		}()
	}
	for a = s; *s != 0 && !(wcschr(c, *s) != nil); *(*uintptr)(unsafe.Pointer(&s)) += 4 {
	}
	return uint64((uintptr(unsafe.Pointer(s)) - uintptr(unsafe.Pointer(a))) / 4)
}
