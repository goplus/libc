package libc

func Getc(f *struct__IO_FILE) int32 {
	return do_getc(f)
}
