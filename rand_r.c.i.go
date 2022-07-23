package libc

func _cgos_temper__rand_r(x uint32) uint32 {
	x ^= x >> int32(11)
	x ^= x << int32(7) & uint32(2636928640)
	x ^= x << int32(15) & uint32(4022730752)
	x ^= x >> int32(18)
	return x
}
func Rand_r(seed *uint32) int32 {
	return int32(_cgos_temper__rand_r(func() (_cgo_ret uint32) {
		_cgo_addr := &*seed
		*_cgo_addr = *seed*uint32(1103515245) + uint32(12345)
		return *_cgo_addr
	}()) / uint32(2))
}
