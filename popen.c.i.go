package libc

import unsafe "unsafe"

func Popen(cmd *int8, mode *int8) *Struct__IO_FILE {
	var p [2]int32
	var op int32
	var e int32
	var pid int32
	var f *Struct__IO_FILE
	var fa _cgoa_19_popen
	if int32(*mode) == 'r' {
		op = int32(0)
	} else if int32(*mode) == 'w' {
		op = int32(1)
	} else {
		*__errno_location() = int32(22)
		return (*Struct__IO_FILE)(nil)
	}
	if pipe2((*int32)(unsafe.Pointer(&p)), int32(524288)) != 0 {
		return (*Struct__IO_FILE)(nil)
	}
	f = Fdopen(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(op)*4)), mode)
	if !(f != nil) {
		__syscall1(int64(6), int64(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(0))*4))))
		__syscall1(int64(6), int64(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1))*4))))
		return (*Struct__IO_FILE)(nil)
	}
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1)-op)*4)) == int32(1)-op {
		var tmp int32 = fcntl(int32(1)-op, int32(1030), int32(0))
		if tmp < int32(0) {
			e = *__errno_location()
			goto fail
		}
		__syscall1(int64(6), int64(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1)-op)*4))))
		*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1)-op)*4)) = tmp
	}
	e = int32(12)
	if !(posix_spawn_file_actions_init(&fa) != 0) {
		if !(posix_spawn_file_actions_adddup2(&fa, *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1)-op)*4)), int32(1)-op) != 0) {
			if !(func() (_cgo_ret int32) {
				_cgo_addr := &e
				*_cgo_addr = posix_spawn(&pid, (*int8)(unsafe.Pointer(&[8]int8{'/', 'b', 'i', 'n', '/', 's', 'h', '\x00'})), &fa, nil, (**int8)(unsafe.Pointer(&[4]*int8{(*int8)(unsafe.Pointer(&[3]int8{'s', 'h', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'-', 'c', '\x00'})), (*int8)(unsafe.Pointer(cmd)), nil})), __environ)
				return *_cgo_addr
			}() != 0) {
				posix_spawn_file_actions_destroy(&fa)
				f.Pipe_pid = pid
				if !(Strchr(mode, 'e') != nil) {
					fcntl(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(op)*4)), int32(2), int32(0))
				}
				__syscall1(int64(6), int64(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1)-op)*4))))
				for {
					if __need_unlock != 0 {
						__unlockfile(f)
					}
					if true {
						break
					}
				}
				return f
			}
		}
		posix_spawn_file_actions_destroy(&fa)
	}
fail:
	Fclose(f)
	__syscall1(int64(6), int64(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1)-op)*4))))
	*__errno_location() = e
	return (*Struct__IO_FILE)(nil)
}
