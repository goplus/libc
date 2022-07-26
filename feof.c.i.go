package libc

func Feof(f *Struct__IO_FILE) int32 {
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	var ret int32 = func() int32 {
		if !!(f.Flags&uint32(16) != 0) {
			return 1
		} else {
			return 0
		}
	}()
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return ret
}
