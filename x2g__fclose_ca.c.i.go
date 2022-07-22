package libc

func __fclose_ca(f *struct__IO_FILE) int32 {
	return f.close(f)
}
