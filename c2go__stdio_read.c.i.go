package libc

import unsafe "unsafe"

func __stdio_read(f *struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var iov [2]struct_iovec = [2]struct_iovec{struct_iovec{unsafe.Pointer(buf), len - func() uint64 {
		if !!(f.buf_size != 0) {
			return 1
		} else {
			return 0
		}
	}()}, struct_iovec{unsafe.Pointer(f.buf), f.buf_size}}
	var cnt int64
	cnt = func() int64 {
		if (*(*struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer((*struct_iovec)(unsafe.Pointer(&iov)))) + uintptr(0)*16))).iov_len != 0 {
			return __syscall3_r1(int64(120), int64(f.fd), int64(uintptr(unsafe.Pointer((*struct_iovec)(unsafe.Pointer(&iov))))), int64(2))
		} else {
			return __syscall3_r1(int64(3), int64(f.fd), int64(uintptr((*(*struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer((*struct_iovec)(unsafe.Pointer(&iov)))) + uintptr(1)*16))).iov_base)), int64((*(*struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer((*struct_iovec)(unsafe.Pointer(&iov)))) + uintptr(1)*16))).iov_len))
		}
	}()
	if cnt <= int64(0) {
		f.flags |= uint32(func() int32 {
			if cnt != 0 {
				return 32
			} else {
				return 16
			}
		}())
		return uint64(0)
	}
	if uint64(cnt) <= (*(*struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer((*struct_iovec)(unsafe.Pointer(&iov)))) + uintptr(0)*16))).iov_len {
		return uint64(cnt)
	}
	cnt -= int64((*(*struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer((*struct_iovec)(unsafe.Pointer(&iov)))) + uintptr(0)*16))).iov_len)
	f.rpos = f.buf
	f.rend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.buf)) + uintptr(cnt)))
	if f.buf_size != 0 {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(len-uint64(1)))) = *func() (_cgo_ret *uint8) {
			_cgo_addr := &f.rpos
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}
	return len
}
