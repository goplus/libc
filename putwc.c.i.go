package libc

func putwc(c uint32, f *Struct__IO_FILE) uint32 {
	return fputwc(c, f)
}
