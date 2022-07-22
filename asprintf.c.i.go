package libc

func Asprintf(s **int8, fmt *int8, __cgo_args ...interface {
}) int32 {
	var ret int32
	var ap []interface {
	}
	ap = __cgo_args
	ret = Vasprintf(s, fmt, ap)
	return ret
}
