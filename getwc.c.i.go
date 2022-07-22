package libc

func getwc(f *struct__IO_FILE) uint32 {
	return fgetwc(f)
}
