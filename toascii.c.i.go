package libc

func Toascii(c int32) int32 {
	return c & int32(127)
}
