package libc

import unsafe "unsafe"

func Strlen(s *int8) uint64 {
	var a *int8 = s
	type word = uint64
	var w *uint64
	for ; uint64(uintptr(unsafe.Pointer(s)))%8 != 0; *(*uintptr)(unsafe.Pointer(&s))++ {
		if !(*s != 0) {
			return uint64(uintptr(unsafe.Pointer(s)) - uintptr(unsafe.Pointer(a)))
		}
	}
	for w = (*uint64)(unsafe.Pointer(s)); !((*w-72340172838076673) & ^*w & 9259542123273814144 != 0); *(*uintptr)(unsafe.Pointer(&w)) += 8 {
	}
	s = (*int8)(unsafe.Pointer(w))
	for ; *s != 0; *(*uintptr)(unsafe.Pointer(&s))++ {
	}
	return uint64(uintptr(unsafe.Pointer(s)) - uintptr(unsafe.Pointer(a)))
}
