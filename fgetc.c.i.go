package libc

func Fgetc(f *Struct__IO_FILE) int32 {
	return do_getc(f)
}
