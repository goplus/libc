package libc

import unsafe "unsafe"

func __freadahead(f *struct__IO_FILE) uint64 {
	return uint64(func() int64 {
		if f.rend != nil {
			return int64(uintptr(unsafe.Pointer(f.rend)) - uintptr(unsafe.Pointer(f.rpos)))
		} else {
			return int64(0)
		}
	}())
}
func __freadptr(f *struct__IO_FILE, sizep *uint64) *int8 {
	if uintptr(unsafe.Pointer(f.rpos)) == uintptr(unsafe.Pointer(f.rend)) {
		return (*int8)(nil)
	}
	*sizep = uint64(uintptr(unsafe.Pointer(f.rend)) - uintptr(unsafe.Pointer(f.rpos)))
	return (*int8)(unsafe.Pointer(f.rpos))
}
func __freadptrinc(f *struct__IO_FILE, inc uint64) {
	*(*uintptr)(unsafe.Pointer(&f.rpos)) += uintptr(inc)
}
func __fseterr(f *struct__IO_FILE) {
	f.flags |= uint32(32)
}
