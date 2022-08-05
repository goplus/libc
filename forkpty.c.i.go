package libc

import unsafe "unsafe"

func forkpty(pm *int32, name *int8, tio *struct_termios, ws *Struct_winsize) int32 {
	var m int32
	var s int32
	var ec int32 = int32(0)
	var p [2]int32
	var cs int32
	var pid int32 = -1
	var set Struct___sigset_t
	var oldset Struct___sigset_t
	if openpty(&m, &s, name, tio, ws) < int32(0) {
		return -1
	}
	sigfillset(&set)
	pthread_sigmask(int32(0), &set, &oldset)
	pthread_setcancelstate(int32(1), &cs)
	if Pipe2((*int32)(unsafe.Pointer(&p)), int32(524288)) != 0 {
		Close(s)
		goto out
	}
	pid = Fork()
	if !(pid != 0) {
		Close(m)
		Close(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(0))*4)))
		if login_tty(s) != 0 {
			Write(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1))*4)), unsafe.Pointer(&*X__errno_location()), 4)
			X_exit(int32(127))
		}
		Close(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1))*4)))
		pthread_setcancelstate(cs, nil)
		pthread_sigmask(int32(2), &oldset, nil)
		return int32(0)
	}
	Close(s)
	Close(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1))*4)))
	if Read(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(0))*4)), unsafe.Pointer(&ec), 4) > int64(0) {
		var status int32
		waitpid(pid, &status, int32(0))
		pid = -1
		*X__errno_location() = ec
	}
	Close(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(0))*4)))
out:
	if pid > int32(0) {
		*pm = m
	} else {
		Close(m)
	}
	pthread_setcancelstate(cs, nil)
	pthread_sigmask(int32(2), &oldset, nil)
	return pid
}
