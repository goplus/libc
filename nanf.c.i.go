package libc

import unsafe "unsafe"

func Nanf(s *int8) float32 {
	return X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))
}
