package libc

import unsafe "unsafe"

func Strerror_r(err int32, buf *int8, buflen uint64) int32 {
	var msg *int8 = Strerror(err)
	var l uint64 = Strlen(msg)
	if l >= buflen {
		if buflen != 0 {
			Memcpy(unsafe.Pointer(buf), unsafe.Pointer(msg), buflen-uint64(1))
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(buflen-uint64(1)))) = int8(0)
		}
		return int32(34)
	}
	Memcpy(unsafe.Pointer(buf), unsafe.Pointer(msg), l+uint64(1))
	return int32(0)
}
