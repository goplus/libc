package libc

import unsafe "unsafe"

func strnlen(s *int8, n uint32) uint32 {
	var p *int8 = (*int8)(memchr(unsafe.Pointer(s), 0, n))
	return uint32(func() int64 {
		if p != nil {
			return int64(uintptr(unsafe.Pointer(p)) - uintptr(unsafe.Pointer(s)))
		} else {
			return int64(n)
		}
	}())
}
