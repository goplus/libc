package libc

func Printf(fmt *int8, __cgo_args ...interface {
}) int32 {
	var ret int32
	var ap []interface {
	}
	ap = __cgo_args
	ret = vfprintf(&__stdout_FILE, fmt, ap)
	return ret
}
