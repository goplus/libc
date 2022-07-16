package libc

import unsafe "unsafe"

func bzero(s unsafe.Pointer, n uint64) {
	Memset(s, int32(0), n)
}
