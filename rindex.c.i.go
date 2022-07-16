package libc

func rindex(s *int8, c int32) *int8 {
	return Strrchr(s, c)
}
