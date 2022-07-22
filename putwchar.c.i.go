package libc

func putwchar(c uint32) uint32 {
	return fputwc(c, &__stdout_FILE)
}
