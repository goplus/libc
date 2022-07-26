package libc

import unsafe "unsafe"

func __freadahead(f *Struct__IO_FILE) uint64 {
	return uint64(func() int64 {
		if f.Rend != nil {
			return int64(uintptr(unsafe.Pointer(f.Rend)) - uintptr(unsafe.Pointer(f.Rpos)))
		} else {
			return int64(0)
		}
	}())
}
func __freadptr(f *Struct__IO_FILE, sizep *uint64) *int8 {
	if uintptr(unsafe.Pointer(f.Rpos)) == uintptr(unsafe.Pointer(f.Rend)) {
		return (*int8)(nil)
	}
	*sizep = uint64(uintptr(unsafe.Pointer(f.Rend)) - uintptr(unsafe.Pointer(f.Rpos)))
	return (*int8)(unsafe.Pointer(f.Rpos))
}
func __freadptrinc(f *Struct__IO_FILE, inc uint64) {
	*(*uintptr)(unsafe.Pointer(&f.Rpos)) += uintptr(inc)
}
func __fseterr(f *Struct__IO_FILE) {
	f.Flags |= uint32(32)
}
