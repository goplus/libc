package libc

func vwprintf(fmt *uint32, ap []interface {
}) int32 {
	return vfwprintf(&__stdout_FILE, fmt, ap)
}
