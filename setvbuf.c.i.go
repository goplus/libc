package libc

import unsafe "unsafe"

func Setvbuf(f *struct__IO_FILE, buf *int8, type_ int32, size uint64) int32 {
	f.lbf = -1
	if type_ == int32(2) {
		f.buf_size = uint64(0)
	} else if type_ == int32(1) || type_ == int32(0) {
		if buf != nil && size >= uint64(8) {
			f.buf = (*uint8)(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(8))))))
			f.buf_size = size - uint64(8)
		}
		if type_ == int32(1) && f.buf_size != 0 {
			f.lbf = int32('\n')
		}
	} else {
		return -1
	}
	f.flags |= uint32(64)
	return int32(0)
}
