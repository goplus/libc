package libc

import unsafe "unsafe"

func Fsetpos(f *Struct__IO_FILE, pos *Union__G_fpos64_t) int32 {
	return __fseeko(f, *(*int64)(unsafe.Pointer(pos)), int32(0))
}
