package libc

import unsafe "unsafe"

func Malloc(size uint64) unsafe.Pointer {
	buf := make([]byte, size)
	return unsafe.Pointer(&buf[0])
}

func Calloc(size, n uint64) unsafe.Pointer {
	return Malloc(size * n)
}

func Free(p unsafe.Pointer) {
}

func __libc_malloc(size uint64) unsafe.Pointer {
	return Malloc(size)
}
