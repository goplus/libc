package libc

func Strncpy(d *int8, s *int8, n uint64) *int8 {
	__stpncpy(d, s, n)
	return d
}
