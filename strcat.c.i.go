package libc

import unsafe "unsafe"

func Strcat(dest *int8, src *int8) *int8 {
	Strcpy((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dest))+uintptr(Strlen(dest)))), src)
	return dest
}
