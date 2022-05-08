package libc

import unsafe "unsafe"

func __libc_malloc(size uint64) unsafe.Pointer {
	buf := make([]byte, size)
	return unsafe.Pointer(&buf[0])
}
