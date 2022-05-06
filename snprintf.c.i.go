package libc

func Snprintf(s *int8, n uint64, fmt *int8, __cgo_args ...interface {
}) int32 {
	var ret int32
	var ap []interface {
	}
	ap = __cgo_args
	ret = Vsnprintf(s, n, fmt, ap)
	return ret
}
