package libc

func tcsetpgrp(fd int32, pgrp int32) int32 {
	var pgrp_int int32 = pgrp
	return Ioctl(fd, int32(21520), &pgrp_int)
}
