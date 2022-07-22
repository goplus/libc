package libc

func Setlinebuf(f *struct__IO_FILE) {
	Setvbuf(f, nil, int32(1), uint64(0))
}
