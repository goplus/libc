package libc

import unsafe "unsafe"

func __map_file(pathname *int8, size *uint64) *uint8 {
	var st struct_kstat
	var map_ *uint8 = (*uint8)(unsafe.Pointer(uintptr(18446744073709551615)))
	var fd int32 = int32(__syscall2_r1(int64(5), int64(uintptr(unsafe.Pointer(pathname))), int64(0|524288|2048|131072)))
	if fd < 0 {
		return (*uint8)(nil)
	}
	if !(__syscall2_r1(int64(339), int64(fd), int64(uintptr(unsafe.Pointer(&st)))) != 0) {
		map_ = (*uint8)(__mmap(nil, uint64(st.st_size), 1, 1, fd, int64(0)))
		*size = uint64(st.st_size)
	}
	__syscall1(int64(6), int64(fd))
	return func() *uint8 {
		if uintptr(unsafe.Pointer(map_)) == uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(18446744073709551615))))) {
			return nil
		} else {
			return map_
		}
	}()
}
