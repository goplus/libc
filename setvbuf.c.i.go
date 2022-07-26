package libc

import unsafe "unsafe"

func Setvbuf(f *Struct__IO_FILE, buf *int8, type_ int32, size uint64) int32 {
	f.Lbf = -1
	if type_ == int32(2) {
		f.Buf_size = uint64(0)
	} else if type_ == int32(1) || type_ == int32(0) {
		if buf != nil && size >= uint64(8) {
			f.Buf = (*uint8)(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(int32(8))))))
			f.Buf_size = size - uint64(8)
		}
		if type_ == int32(1) && f.Buf_size != 0 {
			f.Lbf = int32('\n')
		}
	} else {
		return -1
	}
	f.Flags |= uint32(64)
	return int32(0)
}
