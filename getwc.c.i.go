package libc

func getwc(f *Struct__IO_FILE) uint32 {
	return fgetwc(f)
}
