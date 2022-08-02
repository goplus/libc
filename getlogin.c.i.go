package libc

import unsafe "unsafe"

func Getlogin() *int8 {
	return Getenv((*int8)(unsafe.Pointer(&[8]int8{'L', 'O', 'G', 'N', 'A', 'M', 'E', '\x00'})))
}
