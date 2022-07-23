package libc

func getpriority(which int32, who uint32) int32 {
	var ret int32 = int32(__syscall2_r1(int64(100), int64(which), int64(who)))
	if ret < int32(0) {
		return ret
	}
	return int32(20) - ret
}
