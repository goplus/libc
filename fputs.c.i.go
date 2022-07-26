package libc

import unsafe "unsafe"

func Fputs(s *int8, f *Struct__IO_FILE) int32 {
	var l uint64 = Strlen(s)
	return func() int32 {
		if Fwrite(unsafe.Pointer(s), uint64(1), l, f) == l {
			return 1
		} else {
			return 0
		}
	}() - int32(1)
}
