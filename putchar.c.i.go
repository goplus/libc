package libc

func Putchar(c int32) int32 {
	return do_putc(c, &__stdout_FILE)
}
