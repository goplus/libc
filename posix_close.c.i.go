package libc

func posix_close(fd int32, flags int32) int32 {
	return Close(fd)
}
