package libc

func Setbuf(f *struct__IO_FILE, buf *int8) {
	Setvbuf(f, buf, func() int32 {
		if buf != nil {
			return int32(0)
		} else {
			return int32(2)
		}
	}(), uint64(1024))
}
