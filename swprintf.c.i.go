package libc

func swprintf(s *uint32, n uint64, fmt *uint32, __cgo_args ...interface {
}) int32 {
	var ret int32
	var ap []interface {
	}
	ap = __cgo_args
	ret = vswprintf(s, n, fmt, ap)
	return ret
}
