package libc

import unsafe "unsafe"

func wcsspn(s *uint32, c *uint32) uint64 {
	var a *uint32
	for a = s; *s != 0 && wcschr(c, *s) != nil; *(*uintptr)(unsafe.Pointer(&s)) += 4 {
	}
	return uint64((uintptr(unsafe.Pointer(s)) - uintptr(unsafe.Pointer(a))) / 4)
}
