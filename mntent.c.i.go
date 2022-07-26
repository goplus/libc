package libc

import unsafe "unsafe"

var _cgos_internal_buf_mntent *int8
var _cgos_internal_bufsize_mntent uint64

func setmntent(name *int8, mode *int8) *Struct__IO_FILE {
	return Fopen(name, mode)
}
func endmntent(f *Struct__IO_FILE) int32 {
	if f != nil {
		Fclose(f)
	}
	return int32(1)
}
func getmntent_r(f *Struct__IO_FILE, mnt *struct_mntent, linebuf *int8, buflen int32) *struct_mntent {
	var cnt int32
	var n [8]int32
	var use_internal int32 = func() int32 {
		if uintptr(unsafe.Pointer(linebuf)) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&_cgos_internal_buf_mntent)))) {
			return 1
		} else {
			return 0
		}
	}()
	mnt.mnt_freq = int32(0)
	mnt.mnt_passno = int32(0)
	for {
		if use_internal != 0 {
			Getline(&_cgos_internal_buf_mntent, &_cgos_internal_bufsize_mntent, f)
			linebuf = _cgos_internal_buf_mntent
		} else {
			Fgets(linebuf, buflen, f)
		}
		if Feof(f) != 0 || Ferror(f) != 0 {
			return (*struct_mntent)(nil)
		}
		if !(Strchr(linebuf, '\n') != nil) {
			Fscanf(f, (*int8)(unsafe.Pointer(&[12]int8{'%', '*', '[', '^', '\n', ']', '%', '*', '[', '\n', ']', '\x00'})))
			*__errno_location() = int32(34)
			return (*struct_mntent)(nil)
		}
		cnt = Sscanf(linebuf, (*int8)(unsafe.Pointer(&[39]int8{' ', '%', 'n', '%', '*', 's', '%', 'n', ' ', '%', 'n', '%', '*', 's', '%', 'n', ' ', '%', 'n', '%', '*', 's', '%', 'n', ' ', '%', 'n', '%', '*', 's', '%', 'n', ' ', '%', 'd', ' ', '%', 'd', '\x00'})), (*int32)(unsafe.Pointer(&n)), (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n))))+uintptr(int32(1))*4)), (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n))))+uintptr(int32(2))*4)), (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n))))+uintptr(int32(3))*4)), (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n))))+uintptr(int32(4))*4)), (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n))))+uintptr(int32(5))*4)), (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n))))+uintptr(int32(6))*4)), (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n))))+uintptr(int32(7))*4)), &mnt.mnt_freq, &mnt.mnt_passno)
		if !(cnt < int32(2) || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(linebuf)) + uintptr(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n)))) + uintptr(int32(0))*4)))))) == '#') {
			break
		}
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(linebuf)) + uintptr(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n)))) + uintptr(int32(1))*4))))) = int8(0)
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(linebuf)) + uintptr(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n)))) + uintptr(int32(3))*4))))) = int8(0)
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(linebuf)) + uintptr(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n)))) + uintptr(int32(5))*4))))) = int8(0)
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(linebuf)) + uintptr(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n)))) + uintptr(int32(7))*4))))) = int8(0)
	mnt.mnt_fsname = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(linebuf)) + uintptr(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n)))) + uintptr(int32(0))*4)))))
	mnt.mnt_dir = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(linebuf)) + uintptr(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n)))) + uintptr(int32(2))*4)))))
	mnt.mnt_type = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(linebuf)) + uintptr(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n)))) + uintptr(int32(4))*4)))))
	mnt.mnt_opts = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(linebuf)) + uintptr(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&n)))) + uintptr(int32(6))*4)))))
	return mnt
}
func getmntent(f *Struct__IO_FILE) *struct_mntent {
	return getmntent_r(f, &_cgos_mnt_mntent, (*int8)(unsafe.Pointer(&_cgos_internal_buf_mntent)), int32(0))
}

var _cgos_mnt_mntent struct_mntent

func addmntent(f *Struct__IO_FILE, mnt *struct_mntent) int32 {
	if Fseek(f, int64(0), int32(2)) != 0 {
		return int32(1)
	}
	return func() int32 {
		if Fprintf(f, (*int8)(unsafe.Pointer(&[19]int8{'%', 's', '\t', '%', 's', '\t', '%', 's', '\t', '%', 's', '\t', '%', 'd', '\t', '%', 'd', '\n', '\x00'})), mnt.mnt_fsname, mnt.mnt_dir, mnt.mnt_type, mnt.mnt_opts, mnt.mnt_freq, mnt.mnt_passno) < int32(0) {
			return 1
		} else {
			return 0
		}
	}()
}
func hasmntopt(mnt *struct_mntent, opt *int8) *int8 {
	return Strstr(mnt.mnt_opts, opt)
}
