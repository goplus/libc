package libc

import unsafe "unsafe"

func explicit_bzero(d unsafe.Pointer, n uint64) {
	d = Memset(d, int32(0), n)
}
