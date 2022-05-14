package libc

import unsafe "unsafe"

func __stdio_write(f *struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var iovs [2]struct_iovec = [2]struct_iovec{struct_iovec{unsafe.Pointer(f.wbase), uint64(uintptr(unsafe.Pointer(f.wpos)) - uintptr(unsafe.Pointer(f.wbase)))}, struct_iovec{unsafe.Pointer(buf), len}}
	var iov *struct_iovec = (*struct_iovec)(unsafe.Pointer(&iovs))
	var rem uint64 = (*(*struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).iov_len + (*(*struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(1))*16))).iov_len
	var iovcnt int32 = int32(2)
	var cnt int64
	for {
		cnt = __syscall3_r1(int64(121), int64(f.fd), int64(uintptr(unsafe.Pointer(iov))), int64(iovcnt))
		if uint64(cnt) == rem {
			f.wend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.buf)) + uintptr(f.buf_size)))
			f.wpos = func() (_cgo_ret *uint8) {
				_cgo_addr := &f.wbase
				*_cgo_addr = f.buf
				return *_cgo_addr
			}()
			return len
		}
		if cnt < int64(0) {
			f.wpos = func() (_cgo_ret *uint8) {
				_cgo_addr := &f.wbase
				*_cgo_addr = func() (_cgo_ret *uint8) {
					_cgo_addr := &f.wend
					*_cgo_addr = (*uint8)(nil)
					return *_cgo_addr
				}()
				return *_cgo_addr
			}()
			f.flags |= uint32(32)
			return func() uint64 {
				if iovcnt == int32(2) {
					return uint64(0)
				} else {
					return len - (*(*struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).iov_len
				}
			}()
		}
		rem -= uint64(cnt)
		if uint64(cnt) > (*(*struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).iov_len {
			cnt -= int64((*(*struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).iov_len)
			*(*uintptr)(unsafe.Pointer(&iov)) += 16
			iovcnt--
		}
		(*(*struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).iov_base = unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)((*(*struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).iov_base))) + uintptr(cnt))))
		(*(*struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer(iov)) + uintptr(int32(0))*16))).iov_len -= uint64(cnt)
	}
	return 0
}
