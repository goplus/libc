package libc

import unsafe "unsafe"

func bcmp(s1 unsafe.Pointer, s2 unsafe.Pointer, n uint64) int32 {
	return Memcmp(s1, s2, n)
}
