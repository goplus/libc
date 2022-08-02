package libc

import unsafe "unsafe"

func Ctermid(s *int8) *int8 {
	return func() *int8 {
		if s != nil {
			return Strcpy(s, (*int8)(unsafe.Pointer(&[9]int8{'/', 'd', 'e', 'v', '/', 't', 't', 'y', '\x00'})))
		} else {
			return (*int8)(unsafe.Pointer(&[9]int8{'/', 'd', 'e', 'v', '/', 't', 't', 'y', '\x00'}))
		}
	}()
}
