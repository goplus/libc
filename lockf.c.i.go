package libc

func lockf(fd int32, op int32, size int64) int32 {
	var l struct_flock = struct_flock{int16(1), int16(1), 0, size, 0}
	switch op {
	case int32(3):
		l.l_type = int16(0)
		if fcntl(fd, int32(12), &l) < int32(0) {
			return -1
		}
		if int32(l.l_type) == int32(2) || l.l_pid == Getpid() {
			return int32(0)
		}
		*X__errno_location() = int32(13)
		return -1
	case int32(0):
		l.l_type = int16(2)
	case int32(2):
		return fcntl(fd, int32(13), &l)
	case int32(1):
		return fcntl(fd, int32(14), &l)
	}
	*X__errno_location() = int32(22)
	return -1
}
