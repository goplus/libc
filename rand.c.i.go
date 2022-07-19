package libc

var seed_cgo18_rand uint64

func Srand(s uint32) {
	seed_cgo18_rand = uint64(s - uint32(1))
}
func Rand() int32 {
	seed_cgo18_rand = uint64(6364136223846793005)*seed_cgo18_rand + uint64(1)
	return int32(seed_cgo18_rand >> int32(33))
}
