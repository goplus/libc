package libc

func Fputc(c int32, f *Struct__IO_FILE) int32 {
	return do_putc(c, f)
}
