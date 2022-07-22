package libc

func Vprintf(fmt *int8, ap []interface {
}) int32 {
	return Vfprintf(&__stdout_FILE, fmt, ap)
}
