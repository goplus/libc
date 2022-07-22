package libc

func swscanf(s *uint32, fmt *uint32, __cgo_args ...interface {
}) int32 {
	var ret int32
	var ap []interface {
	}
	ap = __cgo_args
	ret = vswscanf(s, fmt, ap)
	return ret
}
