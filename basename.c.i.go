package libc

import unsafe "unsafe"

func Basename(s *int8) *int8 {
	var i uint64
	if !(s != nil) || !(*s != 0) {
		return (*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'}))
	}
	i = Strlen(s) - uint64(1)
	for ; i != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) == '/'; i-- {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i))) = int8(0)
	}
	for ; i != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i-uint64(1))))) != '/'; i-- {
	}
	return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))
}
