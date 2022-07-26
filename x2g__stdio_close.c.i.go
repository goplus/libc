package libc

func _cgos_dummy___stdio_close(fd int32) int32 {
	return fd
}
func __stdio_close(f *Struct__IO_FILE) int32 {
	return int32(__syscall1_r1(int64(6), int64(__aio_close(f.Fd))))
}
