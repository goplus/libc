package libc

func getpgrp() int32 {
	return int32(__syscall1(int64(151), int64(0)))
}
