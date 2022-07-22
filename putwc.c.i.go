package libc

func putwc(c uint32, f *struct__IO_FILE) uint32 {
	return fputwc(c, f)
}
