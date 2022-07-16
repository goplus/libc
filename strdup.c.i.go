package libc

import unsafe "unsafe"

func strdup(s *int8) *int8 {
	var l uint64 = Strlen(s)
	var d *int8 = (*int8)(Malloc(l + uint64(1)))
	if !(d != nil) {
		return (*int8)(nil)
	}
	return (*int8)(Memcpy(unsafe.Pointer(d), unsafe.Pointer(s), l+uint64(1)))
}
