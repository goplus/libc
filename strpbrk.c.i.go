package libc

import unsafe "unsafe"

func Strpbrk(s *int8, b *int8) *int8 {
	*(*uintptr)(unsafe.Pointer(&s)) += uintptr(Strcspn(s, b))
	return func() *int8 {
		if int32(*s) != 0 {
			return (*int8)(unsafe.Pointer(s))
		} else {
			return nil
		}
	}()
}
