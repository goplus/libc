package libc

import unsafe "unsafe"

func wcsnlen(s *uint32, n uint64) uint64 {
	var z *uint32 = wmemchr(s, uint32(0), n)
	if z != nil {
		n = uint64((uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(s))) / 4)
	}
	return n
}
