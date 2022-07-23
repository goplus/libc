package libc

import unsafe "unsafe"

type _cgos_history_nftw struct {
	chain *_cgos_history_nftw
	dev   uint64
	ino   uint64
	level int32
	base  int32
}

func _cgos_do_nftw_nftw(path *int8, fn func(*int8, *struct_stat, int32, *struct_FTW) int32, fd_limit int32, flags int32, h *_cgos_history_nftw) int32 {
	var l uint64 = Strlen(path)
	var j uint64 = func() uint64 {
		if l != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(path)) + uintptr(l-uint64(1))))) == '/' {
			return l - uint64(1)
		} else {
			return l
		}
	}()
	var st struct_stat
	var new _cgos_history_nftw
	var type_ int32
	var r int32
	var dfd int32
	var err int32
	var lev struct_FTW
	if func() int32 {
		if flags&int32(1) != 0 {
			return lstat(path, &st)
		} else {
			return func() int32 {
				if stat(path, &st) < int32(0) {
					return 1
				} else {
					return 0
				}
			}()
		}
	}() != 0 {
		if !(flags&int32(1) != 0) && *__errno_location() == int32(2) && !(lstat(path, &st) != 0) {
			type_ = int32(7)
		} else if *__errno_location() != int32(13) {
			return -1
		} else {
			type_ = int32(4)
		}
	} else if st.st_mode&uint32(61440) == uint32(16384) {
		if flags&int32(8) != 0 {
			type_ = int32(6)
		} else {
			type_ = int32(2)
		}
	} else if st.st_mode&uint32(61440) == uint32(40960) {
		if flags&int32(1) != 0 {
			type_ = int32(5)
		} else {
			type_ = int32(7)
		}
	} else {
		type_ = int32(1)
	}
	if flags&int32(2) != 0 && h != nil && st.st_dev != h.dev {
		return int32(0)
	}
	new.chain = h
	new.dev = st.st_dev
	new.ino = st.st_ino
	new.level = func() int32 {
		if h != nil {
			return h.level + int32(1)
		} else {
			return int32(0)
		}
	}()
	new.base = int32(j + uint64(1))
	lev.level = new.level
	if h != nil {
		lev.base = h.base
	} else {
		var k uint64
		for k = j; k != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(path)) + uintptr(k)))) == '/'; k-- {
		}
		for ; k != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(path)) + uintptr(k-uint64(1))))) != '/'; k-- {
		}
		lev.base = int32(k)
	}
	if type_ == int32(2) || type_ == int32(6) {
		dfd = open(path, int32(0))
		err = *__errno_location()
		if dfd < int32(0) && err == int32(13) {
			type_ = int32(3)
		}
		if !(fd_limit != 0) {
			close(dfd)
		}
	}
	if !(flags&int32(8) != 0) && func() (_cgo_ret int32) {
		_cgo_addr := &r
		*_cgo_addr = fn(path, &st, type_, &lev)
		return *_cgo_addr
	}() != 0 {
		return r
	}
	for ; h != nil; h = h.chain {
		if h.dev == st.st_dev && h.ino == st.st_ino {
			return int32(0)
		}
	}
	if (type_ == int32(2) || type_ == int32(6)) && fd_limit != 0 {
		if dfd < int32(0) {
			*__errno_location() = err
			return -1
		}
		var d *struct___dirstream = fdopendir(dfd)
		if d != nil {
			var de *struct_dirent
			for func() (_cgo_ret *struct_dirent) {
				_cgo_addr := &de
				*_cgo_addr = readdir(d)
				return *_cgo_addr
			}() != nil {
				if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&de.d_name)))) + uintptr(int32(0))))) == '.' && (!(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&de.d_name)))) + uintptr(int32(1)))) != 0) || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&de.d_name)))) + uintptr(int32(1))))) == '.' && !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&de.d_name)))) + uintptr(int32(2)))) != 0)) {
					continue
				}
				if Strlen((*int8)(unsafe.Pointer(&de.d_name))) >= uint64(4096)-l {
					*__errno_location() = int32(36)
					closedir(d)
					return -1
				}
				*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(path)) + uintptr(j))) = int8('/')
				Strcpy((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(path))+uintptr(j)))))+uintptr(int32(1)))), (*int8)(unsafe.Pointer(&de.d_name)))
				if func() (_cgo_ret int32) {
					_cgo_addr := &r
					*_cgo_addr = _cgos_do_nftw_nftw(path, fn, fd_limit-int32(1), flags, &new)
					return *_cgo_addr
				}() != 0 {
					closedir(d)
					return r
				}
			}
			closedir(d)
		} else {
			close(dfd)
			return -1
		}
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(path)) + uintptr(l))) = int8(0)
	if flags&int32(8) != 0 && func() (_cgo_ret int32) {
		_cgo_addr := &r
		*_cgo_addr = fn(path, &st, type_, &lev)
		return *_cgo_addr
	}() != 0 {
		return r
	}
	return int32(0)
}
func nftw(path *int8, fn func(*int8, *struct_stat, int32, *struct_FTW) int32, fd_limit int32, flags int32) int32 {
	var r int32
	var cs int32
	var l uint64
	var pathbuf [4097]int8
	if fd_limit <= int32(0) {
		return int32(0)
	}
	l = Strlen(path)
	if l > uint64(4096) {
		*__errno_location() = int32(36)
		return -1
	}
	Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(&pathbuf))), unsafe.Pointer(path), l+uint64(1))
	pthread_setcancelstate(int32(1), &cs)
	r = _cgos_do_nftw_nftw((*int8)(unsafe.Pointer(&pathbuf)), fn, fd_limit, flags, nil)
	pthread_setcancelstate(cs, nil)
	return r
}
