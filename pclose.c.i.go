package libc

import unsafe "unsafe"

func Pclose(f *Struct__IO_FILE) int32 {
	var status int32
	var r int32
	var pid int32 = f.Pipe_pid
	Fclose(f)
	for func() (_cgo_ret int32) {
		_cgo_addr := &r
		*_cgo_addr = int32(__syscall4(int64(7), int64(pid), int64(uintptr(unsafe.Pointer(&status))), int64(0), int64(0)))
		return *_cgo_addr
	}() == -4 {
	}
	if r < int32(0) {
		return int32(__syscall_ret(uint64(r)))
	}
	return status
}
