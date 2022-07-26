package libc

import unsafe "unsafe"

func wctrans(class *int8) *int32 {
	if !(Strcmp(class, (*int8)(unsafe.Pointer(&[8]int8{'t', 'o', 'u', 'p', 'p', 'e', 'r', '\x00'}))) != 0) {
		return (*int32)(unsafe.Pointer(uintptr(1)))
	}
	if !(Strcmp(class, (*int8)(unsafe.Pointer(&[8]int8{'t', 'o', 'l', 'o', 'w', 'e', 'r', '\x00'}))) != 0) {
		return (*int32)(unsafe.Pointer(uintptr(2)))
	}
	return (*int32)(nil)
}
func towctrans(wc uint32, trans *int32) uint32 {
	if uintptr(unsafe.Pointer(trans)) == uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(uintptr(1))))) {
		return towupper(wc)
	}
	if uintptr(unsafe.Pointer(trans)) == uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(uintptr(2))))) {
		return towlower(wc)
	}
	return wc
}
func __wctrans_l(s *int8, l *Struct___locale_struct) *int32 {
	return wctrans(s)
}
func __towctrans_l(c uint32, t *int32, l *Struct___locale_struct) uint32 {
	return towctrans(c, t)
}
