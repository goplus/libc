package libc

import unsafe "unsafe"

func wcschr(s *uint32, c uint32) *uint32 {
	if !(c != 0) {
		return (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(s)))) + uintptr(wcslen(s))*4))
	}
	for ; *s != 0 && *s != c; *(*uintptr)(unsafe.Pointer(&s)) += 4 {
	}
	return func() *uint32 {
		if *s != 0 {
			return (*uint32)(unsafe.Pointer(s))
		} else {
			return nil
		}
	}()
}
