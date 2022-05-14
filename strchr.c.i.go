package libc

import unsafe "unsafe"

func Strchr(s *int8, c int32) *int8 {
	var r *int8 = __strchrnul(s, c)
	return func() *int8 {
		if int32(*(*uint8)(unsafe.Pointer(r))) == int32(uint8(c)) {
			return r
		} else {
			return nil
		}
	}()
}
