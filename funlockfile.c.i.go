package libc

func Funlockfile(f *Struct__IO_FILE) {
	if f.Lockcount == int64(1) {
		__unlist_locked_file(f)
		f.Lockcount = int64(0)
		__unlockfile(f)
	} else {
		f.Lockcount--
	}
}
