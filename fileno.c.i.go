package libc

func Fileno(f *Struct__IO_FILE) int32 {
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	var fd int32 = f.Fd
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	if fd < int32(0) {
		*X__errno_location() = int32(9)
		return -1
	}
	return fd
}
