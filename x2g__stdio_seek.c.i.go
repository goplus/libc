package libc

func __stdio_seek(f *Struct__IO_FILE, off int64, whence int32) int64 {
	return X__lseek(f.Fd, off, whence)
}
