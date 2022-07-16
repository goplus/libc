package libc

import unsafe "unsafe"

func strcasestr(h *int8, n *int8) *int8 {
	var l uint64 = Strlen(n)
	for ; *h != 0; *(*uintptr)(unsafe.Pointer(&h))++ {
		if !(strncasecmp(h, n, l) != 0) {
			return (*int8)(unsafe.Pointer(h))
		}
	}
	return (*int8)(nil)
}
