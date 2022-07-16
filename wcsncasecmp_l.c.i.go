package libc

func wcsncasecmp_l(l *uint32, r *uint32, n uint64, locale *struct___locale_struct) int32 {
	return wcsncasecmp(l, r, n)
}
