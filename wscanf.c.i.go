package libc

func wscanf(fmt *uint32, __cgo_args ...interface {
}) int32 {
	var ret int32
	var ap []interface {
	}
	ap = __cgo_args
	ret = vwscanf(fmt, ap)
	return ret
}
