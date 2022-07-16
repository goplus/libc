package libc

import unsafe "unsafe"

func mempcpy(dest unsafe.Pointer, src unsafe.Pointer, n uint64) unsafe.Pointer {
	return unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(Memcpy(dest, src, n)))) + uintptr(n))))
}
