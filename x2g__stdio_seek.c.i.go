package libc

func __stdio_seek(f *struct__IO_FILE, off int64, whence int32) int64 {
	return __lseek(f.fd, off, whence)
}
