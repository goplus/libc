package libc

func wcscasecmp_l(l *uint32, r *uint32, locale *Struct___locale_struct) int32 {
	return wcscasecmp(l, r)
}
