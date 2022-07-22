package libc

func fwprintf(f *struct__IO_FILE, fmt *uint32, __cgo_args ...interface {
}) int32 {
	var ret int32
	var ap []interface {
	}
	ap = __cgo_args
	ret = vfwprintf(f, fmt, ap)
	return ret
}
