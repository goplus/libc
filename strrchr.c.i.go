package libc

import unsafe "unsafe"

func Strrchr(s *int8, c int32) *int8 {
	return (*int8)(__memrchr(unsafe.Pointer(s), c, Strlen(s)+uint64(1)))
}
