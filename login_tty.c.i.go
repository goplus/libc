package libc

func login_tty(fd int32) int32 {
	Setsid()
	if Ioctl(fd, int32(21518), (*int8)(nil)) != 0 {
		return -1
	}
	Dup2(fd, int32(0))
	Dup2(fd, int32(1))
	Dup2(fd, int32(2))
	if fd > int32(2) {
		Close(fd)
	}
	return int32(0)
}
