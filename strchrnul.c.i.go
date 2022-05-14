package libc

import unsafe "unsafe"

func __strchrnul(s *int8, c int32) *int8 {
	c = int32(uint8(c))
	if !(c != 0) {
		return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(s)))) + uintptr(Strlen(s))))
	}
	type word = uint64
	var w *uint64
	for ; uint64(uintptr(unsafe.Pointer(s)))%8 != 0; *(*uintptr)(unsafe.Pointer(&s))++ {
		if !(*s != 0) || int32(*(*uint8)(unsafe.Pointer(s))) == c {
			return (*int8)(unsafe.Pointer(s))
		}
	}
	var k uint64 = 72340172838076673 * uint64(c)
	for w = (*uint64)(unsafe.Pointer(s)); !((*w-72340172838076673) & ^*w & 9259542123273814144 != 0) && !((*w^k-72340172838076673) & ^(*w^k) & 9259542123273814144 != 0); *(*uintptr)(unsafe.Pointer(&w)) += 8 {
	}
	s = (*int8)(unsafe.Pointer(w))
	for ; int32(*s) != 0 && int32(*(*uint8)(unsafe.Pointer(s))) != c; *(*uintptr)(unsafe.Pointer(&s))++ {
	}
	return (*int8)(unsafe.Pointer(s))
}
