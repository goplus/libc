package libc

func Vasprintf(s **int8, fmt *int8, ap []interface {
}) int32 {
	var ap2 []interface {
	}
	ap2 = ap
	var l int32 = Vsnprintf(nil, uint64(0), fmt, ap2)
	if l < int32(0) || !(func() (_cgo_ret *int8) {
		_cgo_addr := &*s
		*_cgo_addr = (*int8)(Malloc(uint64(uint32(l) + uint32(1))))
		return *_cgo_addr
	}() != nil) {
		return -1
	}
	return Vsnprintf(*s, uint64(uint32(l)+uint32(1)), fmt, ap)
}
