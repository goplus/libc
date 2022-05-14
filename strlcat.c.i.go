package libc

import unsafe "unsafe"

func Strlcat(d *int8, s *int8, n uint64) uint64 {
	var l uint64 = Strnlen(d, n)
	if l == n {
		return l + Strlen(s)
	}
	return l + Strlcpy((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(d))+uintptr(l))), s, n-l)
}
