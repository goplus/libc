package libc

func __uflow(f *struct__IO_FILE) int32 {
	var c uint8
	if !(__toread(f) != 0) && f.read(f, &c, uint64(1)) == uint64(1) {
		return int32(c)
	}
	return -1
}
