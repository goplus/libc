package libc

func Funlockfile(f *struct__IO_FILE) {
	if f.lockcount == int64(1) {
		__unlist_locked_file(f)
		f.lockcount = int64(0)
		__unlockfile(f)
	} else {
		f.lockcount--
	}
}
