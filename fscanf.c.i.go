package libc

func Fscanf(f *struct__IO_FILE, fmt *int8, __cgo_args ...interface {
}) int32 {
	var ret int32
	var ap []interface {
	}
	ap = __cgo_args
	ret = Vfscanf(f, fmt, ap)
	return ret
}
