package libc

import unsafe "unsafe"

func __stdio_read(f *Struct__IO_FILE, buf *uint8, len uint64) uint64 {
	var iov [2]Struct_iovec = [2]Struct_iovec{Struct_iovec{unsafe.Pointer(buf), len - func() uint64 {
		if !!(f.Buf_size != 0) {
			return 1
		} else {
			return 0
		}
	}()}, Struct_iovec{unsafe.Pointer(f.Buf), f.Buf_size}}
	var cnt int64
	cnt = func() int64 {
		if (*(*Struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer((*Struct_iovec)(unsafe.Pointer(&iov)))) + uintptr(int32(0))*16))).Iov_len != 0 {
			return __syscall3_r1(int64(120), int64(f.Fd), int64(uintptr(unsafe.Pointer((*Struct_iovec)(unsafe.Pointer(&iov))))), int64(2))
		} else {
			return __syscall3_r1(int64(3), int64(f.Fd), int64(uintptr((*(*Struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer((*Struct_iovec)(unsafe.Pointer(&iov)))) + uintptr(int32(1))*16))).Iov_base)), int64((*(*Struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer((*Struct_iovec)(unsafe.Pointer(&iov)))) + uintptr(int32(1))*16))).Iov_len))
		}
	}()
	if cnt <= int64(0) {
		f.Flags |= uint32(func() int32 {
			if cnt != 0 {
				return int32(32)
			} else {
				return int32(16)
			}
		}())
		return uint64(0)
	}
	if uint64(cnt) <= (*(*Struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer((*Struct_iovec)(unsafe.Pointer(&iov)))) + uintptr(int32(0))*16))).Iov_len {
		return uint64(cnt)
	}
	cnt -= int64((*(*Struct_iovec)(unsafe.Pointer(uintptr(unsafe.Pointer((*Struct_iovec)(unsafe.Pointer(&iov)))) + uintptr(int32(0))*16))).Iov_len)
	f.Rpos = f.Buf
	f.Rend = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(f.Buf)) + uintptr(cnt)))
	if f.Buf_size != 0 {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(buf)) + uintptr(len-uint64(1)))) = *func() (_cgo_ret *uint8) {
			_cgo_addr := &f.Rpos
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}
	return len
}
