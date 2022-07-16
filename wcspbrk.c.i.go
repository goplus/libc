package libc

import unsafe "unsafe"

func wcspbrk(s *uint32, b *uint32) *uint32 {
	*(*uintptr)(unsafe.Pointer(&s)) += uintptr(wcscspn(s, b)) * 4
	return func() *uint32 {
		if *s != 0 {
			return (*uint32)(unsafe.Pointer(s))
		} else {
			return nil
		}
	}()
}
