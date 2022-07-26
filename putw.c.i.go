package libc

import unsafe "unsafe"

func Putw(x int32, f *Struct__IO_FILE) int32 {
	return int32(Fwrite(unsafe.Pointer(&x), 4, uint64(1), f)) - int32(1)
}
