package libc

import unsafe "unsafe"

func __stdio_write(f *Struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var iovs [2]Struct_iovec = [2]Struct_iovec{Struct_iovec{unsafe.Pointer(f.Wbase), uint64(uintptr(unsafe.Pointer(f.Wpos)) - uintptr(unsafe.Pointer(f.Wbase)))}, Struct_iovec{unsafe.Pointer(buf), len}}
	var iov *Struct_iovec = (*Struct_iovec)(unsafe.Pointer(&iovs))
	var rem uint64 = (*(*Struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).Iov_len + (*(*Struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(1))*16))).Iov_len
	var iovcnt int32 = int32(2)
	var cnt int64
	for {
		cnt = __syscall3_r1(int64(121), int64(f.Fd), int64(uintptr(unsafe.Pointer(iov))), int64(iovcnt))
		if uint64(cnt) == rem {
			f.Wend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.Buf)) + uintptr(f.Buf_size)))
			f.Wpos = func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Wbase
				*_cgo_addr = f.Buf
				return *_cgo_addr
			}()
			return len
		}
		if cnt < int64(0) {
			f.Wpos = func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Wbase
				*_cgo_addr = func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Wend
					*_cgo_addr = (*uint8)(nil)
					return *_cgo_addr
				}()
				return *_cgo_addr
			}()
			f.Flags |= uint32(32)
			return func() uint64 {
				if iovcnt == int32(2) {
					return uint64(0)
				} else {
					return len - (*(*Struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).Iov_len
				}
			}()
		}
		rem -= uint64(cnt)
		if uint64(cnt) > (*(*Struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).Iov_len {
			cnt -= int64((*(*Struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).Iov_len)
			*(*uintptr)(unsafe.Pointer(&iov)) += 16
			iovcnt--
		}
		(*(*Struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).Iov_base = unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)((*(*Struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).Iov_base))) + uintptr(cnt))))
		(*(*Struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).Iov_len -= uint64(cnt)
	}
	return 0
}
