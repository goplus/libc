package libc

func wcswcs(haystack *uint32, needle *uint32) *uint32 {
	return wcsstr(haystack, needle)
}
