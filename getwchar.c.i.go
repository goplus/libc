package libc

func getwchar() uint32 {
	return fgetwc(&__stdin_FILE)
}
