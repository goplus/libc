package libc

func Flockfile(f *Struct__IO_FILE) {
	if !(Ftrylockfile(f) != 0) {
		return
	}
	__lockfile(f)
	__register_locked_file(f, __pthread_self())
}
