package libc

func Fgetc(f *struct__IO_FILE) int32 {
	return do_getc(f)
}
