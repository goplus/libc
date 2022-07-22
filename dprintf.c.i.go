package libc

func Dprintf(fd int32, fmt *int8, __cgo_args ...interface {
}) int32 {
	var ret int32
	var ap []interface {
	}
	ap = __cgo_args
	ret = Vdprintf(fd, fmt, ap)
	return ret
}
