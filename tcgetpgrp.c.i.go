package libc

func tcgetpgrp(fd int32) int32 {
	var pgrp int32
	if Ioctl(fd, int32(21519), &pgrp) < int32(0) {
		return -1
	}
	return pgrp
}
