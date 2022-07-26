package libc

func __fclose_ca(f *Struct__IO_FILE) int32 {
	return f.Close(f)
}
