package libc

func index(s *int8, c int32) *int8 {
	return Strchr(s, c)
}
