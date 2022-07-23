package libc

var _cgos_seed_rand uint64

func Srand(s uint32) {
	_cgos_seed_rand = uint64(s - uint32(1))
}
func Rand() int32 {
	_cgos_seed_rand = uint64(6364136223846793005)*_cgos_seed_rand + uint64(1)
	return int32(_cgos_seed_rand >> int32(33))
}
