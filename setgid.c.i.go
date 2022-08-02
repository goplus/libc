package libc

func setgid(gid uint32) int32 {
	return __setxid(int32(181), int32(gid), int32(0), int32(0))
}
