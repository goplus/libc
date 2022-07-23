package libc

import unsafe "unsafe"

func dirname(s *int8) *int8 {
	var i uint64
	if !(s != nil) || !(*s != 0) {
		return (*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'}))
	}
	i = Strlen(s) - uint64(1)
	for ; int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) == '/'; i-- {
		if !(i != 0) {
			return (*int8)(unsafe.Pointer(&[2]int8{'/', '\x00'}))
		}
	}
	for ; int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) != '/'; i-- {
		if !(i != 0) {
			return (*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'}))
		}
	}
	for ; int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) == '/'; i-- {
		if !(i != 0) {
			return (*int8)(unsafe.Pointer(&[2]int8{'/', '\x00'}))
		}
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i+uint64(1)))) = int8(0)
	return s
}
