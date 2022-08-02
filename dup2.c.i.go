package libc

func Dup2(old int32, new int32) int32 {
	var r int32
	for func() (_cgo_ret int32) {
		_cgo_addr := &r
		*_cgo_addr = int32(__syscall2(int64(90), int64(old), int64(new)))
		return *_cgo_addr
	}() == -16 {
	}
	return int32(__syscall_ret(uint64(r)))
}
