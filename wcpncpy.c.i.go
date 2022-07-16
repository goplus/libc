package libc

import unsafe "unsafe"

func wcpncpy(d *uint32, s *uint32, n uint64) *uint32 {
	return (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(wcsncpy(d, s, n))) + uintptr(wcsnlen(s, n))*4))
}
