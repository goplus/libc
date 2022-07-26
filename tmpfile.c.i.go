package libc

import unsafe "unsafe"

func Tmpfile() *Struct__IO_FILE {
	var s [20]int8 = [20]int8{'/', 't', 'm', 'p', '/', 't', 'm', 'p', 'f', 'i', 'l', 'e', '_', 'X', 'X', 'X', 'X', 'X', 'X', '\x00'}
	var fd int32
	var f *Struct__IO_FILE
	var try int32
	for try = int32(0); try < int32(100); try++ {
		__randname((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&s)))) + uintptr(int32(13)))))
		fd = int32(__syscall3_r1(int64(5), int64(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))))), int64(131266), int64(384)))
		if fd >= int32(0) {
			__syscall1(int64(10), int64(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&s))))))
			f = __fdopen(fd, (*int8)(unsafe.Pointer(&[3]int8{'w', '+', '\x00'})))
			if !(f != nil) {
				__syscall1(int64(6), int64(fd))
			}
			return f
		}
	}
	return (*Struct__IO_FILE)(nil)
}
