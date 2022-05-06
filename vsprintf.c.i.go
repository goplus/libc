package libc

func Vsprintf(s *int8, fmt *int8, ap []interface {
}) int32 {
	return Vsnprintf(s, uint64(2147483647), fmt, ap)
}
