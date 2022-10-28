package libc

import unsafe "unsafe"

func wcslen(s *uint32) uint64 {
	var a *uint32
	for a = s; *s != 0; *(*uintptr)(unsafe.Pointer(&s)) += 4 {
	}
	return uint64((uintptr(unsafe.Pointer(s)) - uintptr(unsafe.Pointer(a))) / 4)
}
