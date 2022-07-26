package libc

func Getline(s **int8, n *uint64, f *Struct__IO_FILE) int64 {
	return Getdelim(s, n, '\n', f)
}
