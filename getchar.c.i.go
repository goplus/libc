package libc

func Getchar() int32 {
	return do_getc(&__stdin_FILE)
}
