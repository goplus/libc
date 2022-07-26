package libc

import unsafe "unsafe"

func Fgetpos(f *Struct__IO_FILE, pos *Union__G_fpos64_t) int32 {
	var off int64 = __ftello(f)
	if off < int64(0) {
		return -1
	}
	*(*int64)(unsafe.Pointer(pos)) = off
	return int32(0)
}
