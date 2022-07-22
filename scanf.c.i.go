package libc

func Scanf(fmt *int8, __cgo_args ...interface {
}) int32 {
	var ret int32
	var ap []interface {
	}
	ap = __cgo_args
	ret = Vscanf(fmt, ap)
	return ret
}
