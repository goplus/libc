package libc

func wprintf(fmt *uint32, __cgo_args ...interface {
}) int32 {
	var ret int32
	var ap []interface {
	}
	ap = __cgo_args
	ret = vwprintf(fmt, ap)
	return ret
}
