package libc

func Vscanf(fmt *int8, ap []interface {
}) int32 {
	return Vfscanf(&__stdin_FILE, fmt, ap)
}
