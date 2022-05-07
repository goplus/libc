package libc

import unsafe "unsafe"

func __strchrnul(s *int8, c int32) *int8 {
	c = int32(uint8(c))
	if !(c != 0) {
		return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(s)))) + uintptr(strlen(s))))
	}
	type word = uint64
	var w *uint64
	for ; uint64(uintptr(unsafe.Pointer(s)))%8 != 0; *(*uintptr)(unsafe.Pointer(&s))++ {
		if !(*s != 0) || int32(*(*uint8)(unsafe.Pointer(s))) == c {
			return (*int8)(unsafe.Pointer(s))
		}
	}
	var k uint64 = uint64(18446744073709551615) / uint64(255) * uint64(c)
	for w = (*uint64)(unsafe.Pointer(s)); !((*w-uint64(18446744073709551615)/uint64(255)) & ^*w & (uint64(18446744073709551615)/uint64(255)*uint64(255/2+1)) != 0) && !((*w^k-uint64(18446744073709551615)/uint64(255)) & ^(*w^k) & (uint64(18446744073709551615)/uint64(255)*uint64(255/2+1)) != 0); *(*uintptr)(unsafe.Pointer(&w)) += 8 {
	}
	s = (*int8)(unsafe.Pointer(w))
	for ; int32(*s) != 0 && int32(*(*uint8)(unsafe.Pointer(s))) != c; *(*uintptr)(unsafe.Pointer(&s))++ {
	}
	return (*int8)(unsafe.Pointer(s))
}
