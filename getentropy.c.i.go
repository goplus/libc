package libc

import unsafe "unsafe"

func getentropy(buffer unsafe.Pointer, len uint64) int32 {
	var cs int32
	var ret int32 = int32(0)
	var pos *int8 = (*int8)(buffer)
	if len > uint64(256) {
		*X__errno_location() = int32(5)
		return -1
	}
	pthread_setcancelstate(int32(1), &cs)
	for len != 0 {
		ret = int32(getrandom(unsafe.Pointer(pos), len, uint32(0)))
		if ret < int32(0) {
			if *X__errno_location() == int32(4) {
				continue
			} else {
				break
			}
		}
		*(*uintptr)(unsafe.Pointer(&pos)) += uintptr(ret)
		len -= uint64(ret)
		ret = int32(0)
	}
	pthread_setcancelstate(cs, nil)
	return ret
}
