package libc

func wcsncasecmp_l(l *uint32, r *uint32, n uint64, locale *Struct___locale_struct) int32 {
	return wcsncasecmp(l, r, n)
}
