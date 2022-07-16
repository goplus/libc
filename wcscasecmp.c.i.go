package libc

func wcscasecmp(l *uint32, r *uint32) int32 {
	return wcsncasecmp(l, r, uint64(18446744073709551615))
}
