package libc

import unsafe "unsafe"

func Getw(f *Struct__IO_FILE) int32 {
	var x int32
	return func() int32 {
		if Fread(unsafe.Pointer(&x), 4, uint64(1), f) != 0 {
			return x
		} else {
			return -1
		}
	}()
}
