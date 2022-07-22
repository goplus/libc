package libc

func Rewind(f *struct__IO_FILE) {
	var __need_unlock int32 = func() int32 {
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	__fseeko_unlocked(f, int64(0), int32(0))
	f.flags &= uint32(4294967263)
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
}
