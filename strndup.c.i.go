package libc

import unsafe "unsafe"

func Strndup(s *int8, n uint64) *int8 {
	var l uint64 = Strnlen(s, n)
	var d *int8 = (*int8)(Malloc(l + uint64(1)))
	if !(d != nil) {
		return (*int8)(nil)
	}
	Memcpy(unsafe.Pointer(d), unsafe.Pointer(s), l)
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(l))) = int8(0)
	return d
}
