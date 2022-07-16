package libc

import unsafe "unsafe"

func wcscat(dest *uint32, src *uint32) *uint32 {
	wcscpy((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(dest))+uintptr(wcslen(dest))*4)), src)
	return dest
}
