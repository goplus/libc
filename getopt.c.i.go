package libc

import unsafe "unsafe"

var optarg *int8
var optind int32 = int32(1)
var opterr int32 = int32(1)
var optopt int32
var __optpos int32
var __optreset int32 = int32(0)

func __getopt_msg(a *int8, b *int8, c *int8, l uint64) {
	var f *struct__IO_FILE = &__stderr_FILE
	b = __lctrans_cur(b)
	var __need_unlock int32 = func() int32 {
		if f.lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if Fputs(a, f) >= int32(0) && Fwrite(unsafe.Pointer(b), Strlen(b), uint64(1), f) != 0 && Fwrite(unsafe.Pointer(c), uint64(1), l, f) == l {
		Putc('\n', f)
	}
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
}
func getopt(argc int32, argv **int8, optstring *int8) int32 {
	var i int32
	var c uint32
	var d uint32
	var k int32
	var l int32
	var optchar *int8
	if !(optind != 0) || __optreset != 0 {
		__optreset = int32(0)
		__optpos = int32(0)
		optind = int32(1)
	}
	if optind >= argc || !(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)) != nil) {
		return -1
	}
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)))) + uintptr(int32(0))))) != '-' {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + uintptr(int32(0))))) == '-' {
			optarg = *(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(func() (_cgo_ret int32) {
				_cgo_addr := &optind
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}())*8))
			return int32(1)
		}
		return -1
	}
	if !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)))) + uintptr(int32(1)))) != 0) {
		return -1
	}
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)))) + uintptr(int32(1))))) == '-' && !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)))) + uintptr(int32(2)))) != 0) {
		return func() int32 {
			optind++
			return -1
		}()
	}
	if !(__optpos != 0) {
		__optpos++
	}
	if func() (_cgo_ret int32) {
		_cgo_addr := &k
		*_cgo_addr = Mbtowc(&c, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8))))+uintptr(__optpos))), uint64(4))
		return *_cgo_addr
	}() < int32(0) {
		k = int32(1)
		c = uint32(65533)
	}
	optchar = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)))) + uintptr(__optpos)))
	__optpos += k
	if !(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)))) + uintptr(__optpos))) != 0) {
		optind++
		__optpos = int32(0)
	}
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + uintptr(int32(0))))) == '-' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + uintptr(int32(0))))) == '+' {
		*(*uintptr)(unsafe.Pointer(&optstring))++
	}
	i = int32(0)
	d = uint32(0)
	for {
		l = Mbtowc(&d, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring))+uintptr(i))), uint64(4))
		if l > int32(0) {
			i += l
		} else {
			i++
		}
		if !(l != 0 && d != c) {
			break
		}
	}
	if d != c || c == uint32(':') {
		optopt = int32(c)
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + uintptr(int32(0))))) != ':' && opterr != 0 {
			__getopt_msg(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(int32(0))*8)), (*int8)(unsafe.Pointer(&[24]int8{':', ' ', 'u', 'n', 'r', 'e', 'c', 'o', 'g', 'n', 'i', 'z', 'e', 'd', ' ', 'o', 'p', 't', 'i', 'o', 'n', ':', ' ', '\x00'})), optchar, uint64(k))
		}
		return int32('?')
	}
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + uintptr(i)))) == ':' {
		optarg = (*int8)(nil)
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + uintptr(i+int32(1))))) != ':' || __optpos != 0 {
			optarg = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(func() (_cgo_ret int32) {
				_cgo_addr := &optind
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}())*8)))) + uintptr(__optpos)))
			__optpos = int32(0)
		}
		if optind > argc {
			optopt = int32(c)
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + uintptr(int32(0))))) == ':' {
				return int32(':')
			}
			if opterr != 0 {
				__getopt_msg(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(int32(0))*8)), (*int8)(unsafe.Pointer(&[32]int8{':', ' ', 'o', 'p', 't', 'i', 'o', 'n', ' ', 'r', 'e', 'q', 'u', 'i', 'r', 'e', 's', ' ', 'a', 'n', ' ', 'a', 'r', 'g', 'u', 'm', 'e', 'n', 't', ':', ' ', '\x00'})), optchar, uint64(k))
			}
			return int32('?')
		}
	}
	return int32(c)
}
