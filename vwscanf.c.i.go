package libc

func vwscanf(fmt *uint32, ap []interface {
}) int32 {
	return vfwscanf(&__stdin_FILE, fmt, ap)
}
