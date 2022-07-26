package libc

func Setlinebuf(f *Struct__IO_FILE) {
	Setvbuf(f, nil, int32(1), uint64(0))
}
