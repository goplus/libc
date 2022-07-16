package libc

func toascii(c int32) int32 {
	return c & int32(127)
}
