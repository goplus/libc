package libc

func wcscasecmp_l(l *uint32, r *uint32, locale *struct___locale_struct) int32 {
	return wcscasecmp(l, r)
}
