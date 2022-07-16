package libc

import unsafe "unsafe"

func wcsrchr(s *uint32, c uint32) *uint32 {
	var p *uint32
	for p = (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(wcslen(s))*4)); uintptr(unsafe.Pointer(p)) >= uintptr(unsafe.Pointer(s)) && *p != c; *(*uintptr)(unsafe.Pointer(&p)) -= 4 {
	}
	return func() *uint32 {
		if uintptr(unsafe.Pointer(p)) >= uintptr(unsafe.Pointer(s)) {
			return (*uint32)(unsafe.Pointer(p))
		} else {
			return nil
		}
	}()
}
