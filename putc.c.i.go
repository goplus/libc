package libc

func Putc(c int32, f *struct__IO_FILE) int32 {
	return do_putc(c, f)
}
