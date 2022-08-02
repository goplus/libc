package libc

func setuid(uid uint32) int32 {
	return __setxid(int32(23), int32(uid), int32(0), int32(0))
}
