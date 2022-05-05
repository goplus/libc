package libc

import unsafe "unsafe"

func strnlen(s *int8, n uint64) uint64 {
	var p *int8 = (*int8)(memchr(unsafe.Pointer(s), 0, n))
	return func() uint64 {
		if p != nil {
			return uint64(uintptr(unsafe.Pointer(p)) - uintptr(unsafe.Pointer(s)))
		} else {
			return n
		}
	}()
}
