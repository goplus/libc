package libc

func fwscanf(f *Struct__IO_FILE, fmt *uint32, __cgo_args ...interface {
}) int32 {
	var ret int32
	var ap []interface {
	}
	ap = __cgo_args
	ret = vfwscanf(f, fmt, ap)
	return ret
}
