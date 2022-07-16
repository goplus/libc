package libc

import unsafe "unsafe"

func wcpcpy(d *uint32, s *uint32) *uint32 {
	return (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(wcscpy(d, s))) + uintptr(wcslen(s))*4))
}
