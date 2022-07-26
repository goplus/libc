package libc

func Getc(f *Struct__IO_FILE) int32 {
	return do_getc(f)
}
