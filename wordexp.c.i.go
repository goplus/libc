package libc

import unsafe "unsafe"

func _cgos_reap_wordexp(pid int32) {
	var status int32
	for waitpid(pid, &status, int32(0)) < int32(0) && *X__errno_location() == int32(4) {
	}
}
func _cgos_getword_wordexp(f *Struct__IO_FILE) *int8 {
	var s *int8 = nil
	return func() *int8 {
		if Getdelim(&s, (*uint64)(unsafe.Pointer(&[1]uint64{uint64(0)})), int32(0), f) < int64(0) {
			return nil
		} else {
			return s
		}
	}()
}
func _cgos_do_wordexp_wordexp(s *int8, we *_cgoa_18_wordexp, flags int32) int32 {
	var i uint64
	var l uint64
	var sq int32 = int32(0)
	var dq int32 = int32(0)
	var np uint64 = uint64(0)
	var w *int8
	var tmp **int8
	var redir *int8 = func() *int8 {
		if flags&int32(16) != 0 {
			return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
		} else {
			return (*int8)(unsafe.Pointer(&[12]int8{'2', '>', '/', 'd', 'e', 'v', '/', 'n', 'u', 'l', 'l', '\x00'}))
		}
	}()
	var err int32 = int32(0)
	var f *Struct__IO_FILE
	var wc uint64 = uint64(0)
	var wv **int8 = nil
	var p [2]int32
	var pid int32
	var set Struct___sigset_t
	if flags&int32(8) != 0 {
		wordfree(we)
	}
	if flags&int32(4) != 0 {
		for i = uint64(0); *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i))) != 0; i++ {
			switch int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) {
			case '\\':
				if !(sq != 0) && !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(func() (_cgo_ret uint64) {
					_cgo_addr := &i
					*_cgo_addr++
					return *_cgo_addr
				}()))) != 0) {
					return int32(5)
				}
				break
			case '\'':
				if !(dq != 0) {
					sq ^= int32(1)
				}
				break
			case '"':
				if !(sq != 0) {
					dq ^= int32(1)
				}
				break
			case '(':
				if np != 0 {
					np++
					break
				}
			case ')':
				if np != 0 {
					np--
					break
				}
			case '\n':
				fallthrough
			case '|':
				fallthrough
			case '&':
				fallthrough
			case ';':
				fallthrough
			case '<':
				fallthrough
			case '>':
				fallthrough
			case '{':
				fallthrough
			case '}':
				if !(uint64(sq|dq)|np != 0) {
					return int32(2)
				}
				break
			case '$':
				if sq != 0 {
					break
				}
				if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i+uint64(1))))) == '(' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i+uint64(2))))) == '(' {
					i += uint64(2)
					np += uint64(2)
					break
				} else if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i+uint64(1))))) != '(' {
					break
				}
			case '`':
				if sq != 0 {
					break
				}
				return int32(4)
			}
		}
	}
	if flags&int32(2) != 0 {
		wc = we.we_wordc
		wv = we.we_wordv
	}
	i = wc
	if flags&int32(1) != 0 {
		if we.we_offs > 134217727 {
			goto nospace
		}
		i += we.we_offs
	} else {
		we.we_offs = uint64(0)
	}
	if Pipe2((*int32)(unsafe.Pointer(&p)), int32(524288)) < int32(0) {
		goto nospace
	}
	__block_all_sigs(unsafe.Pointer(&set))
	pid = Fork()
	__restore_sigs(unsafe.Pointer(&set))
	if pid < int32(0) {
		Close(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(0))*4)))
		Close(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1))*4)))
		goto nospace
	}
	if !(pid != 0) {
		if *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1))*4)) == int32(1) {
			fcntl(int32(1), int32(2), int32(0))
		} else {
			Dup2(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1))*4)), int32(1))
		}
		Execl((*int8)(unsafe.Pointer(&[8]int8{'/', 'b', 'i', 'n', '/', 's', 'h', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'s', 'h', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'-', 'c', '\x00'})), (*int8)(unsafe.Pointer(&[30]int8{'e', 'v', 'a', 'l', ' ', '"', 'p', 'r', 'i', 'n', 't', 'f', ' ', '%', 's', '\\', '\\', '\\', '\\', '0', ' ', 'x', ' ', '$', '1', ' ', '$', '2', '"', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'s', 'h', '\x00'})), s, redir, (*int8)(nil))
		X_exit(int32(1))
	}
	Close(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(1))*4)))
	f = Fdopen(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(0))*4)), (*int8)(unsafe.Pointer(&[2]int8{'r', '\x00'})))
	if !(f != nil) {
		Close(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&p)))) + uintptr(int32(0))*4)))
		kill(pid, int32(9))
		_cgos_reap_wordexp(pid)
		goto nospace
	}
	l = func() uint64 {
		if wv != nil {
			return i + uint64(1)
		} else {
			return uint64(0)
		}
	}()
	Free(unsafe.Pointer(_cgos_getword_wordexp(f)))
	if Feof(f) != 0 {
		Fclose(f)
		_cgos_reap_wordexp(pid)
		return int32(5)
	}
	for func() (_cgo_ret *int8) {
		_cgo_addr := &w
		*_cgo_addr = _cgos_getword_wordexp(f)
		return *_cgo_addr
	}() != nil {
		if i+uint64(1) >= l {
			l += l/uint64(2) + uint64(10)
			tmp = (**int8)(Realloc(unsafe.Pointer(wv), l*8))
			if !(tmp != nil) {
				break
			}
			wv = tmp
		}
		*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(wv)) + uintptr(func() (_cgo_ret uint64) {
			_cgo_addr := &i
			_cgo_ret = *_cgo_addr
			*_cgo_addr++
			return
		}())*8)) = w
		*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(wv)) + uintptr(i)*8)) = (*int8)(nil)
	}
	if !(Feof(f) != 0) {
		err = int32(1)
	}
	Fclose(f)
	_cgos_reap_wordexp(pid)
	if !(wv != nil) {
		wv = (**int8)(Calloc(i+uint64(1), 8))
	}
	we.we_wordv = wv
	we.we_wordc = i
	if flags&int32(1) != 0 {
		if wv != nil {
			for i = we.we_offs; i != 0; i-- {
				*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(we.we_wordv)) + uintptr(i-uint64(1))*8)) = (*int8)(nil)
			}
		}
		we.we_wordc -= we.we_offs
	}
	return err
nospace:
	if !(flags&int32(2) != 0) {
		we.we_wordc = uint64(0)
		we.we_wordv = (**int8)(nil)
	}
	return int32(1)
}
func wordexp(s *int8, we *_cgoa_18_wordexp, flags int32) int32 {
	var r int32
	var cs int32
	pthread_setcancelstate(int32(1), &cs)
	r = _cgos_do_wordexp_wordexp(s, we, flags)
	pthread_setcancelstate(cs, nil)
	return r
}
func wordfree(we *_cgoa_18_wordexp) {
	var i uint64
	if !(we.we_wordv != nil) {
		return
	}
	for i = uint64(0); i < we.we_wordc; i++ {
		Free(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(we.we_wordv)) + uintptr(we.we_offs+i)*8))))
	}
	Free(unsafe.Pointer(we.we_wordv))
	we.we_wordv = (**int8)(nil)
	we.we_wordc = uint64(0)
}
