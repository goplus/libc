package libc

import unsafe "unsafe"

func Nan(s *int8) float64 {
	return float64(X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))))
}
