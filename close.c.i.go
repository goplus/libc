package libc

func _cgos_dummy_close(fd int32) int32 {
	return fd
}
func Close(fd int32) int32 {
	fd = __aio_close(fd)
	var r int32 = int32(__syscall_cp(int64(6), int64(fd), int64(0), int64(0), int64(0), int64(0), int64(0)))
	if r == -4 {
		r = int32(0)
	}
	return int32(__syscall_ret(uint64(r)))
}
