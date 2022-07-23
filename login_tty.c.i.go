package libc

func login_tty(fd int32) int32 {
	setsid()
	if ioctl(fd, int32(21518), (*int8)(nil)) != 0 {
		return -1
	}
	dup2(fd, int32(0))
	dup2(fd, int32(1))
	dup2(fd, int32(2))
	if fd > int32(2) {
		close(fd)
	}
	return int32(0)
}
