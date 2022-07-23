package libc

import unsafe "unsafe"

func _cgos_permute_getopt_long(argv **int8, dest int32, src int32) {
	var av **int8 = (**int8)(unsafe.Pointer(argv))
	var tmp *int8 = *(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(av)) + uintptr(src)*8))
	var i int32
	for i = src; i > dest; i-- {
		*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(av)) + uintptr(i)*8)) = *(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(av)) + uintptr(i-int32(1))*8))
	}
	*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(av)) + uintptr(dest)*8)) = tmp
}
func _cgos___getopt_long_getopt_long(argc int32, argv **int8, optstring *int8, longopts *struct_option, idx *int32, longonly int32) int32 {
	var ret int32
	var skipped int32
	var resumed int32
	if !(optind != 0) || __optreset != 0 {
		__optreset = int32(0)
		__optpos = int32(0)
		optind = int32(1)
	}
	if optind >= argc || !(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)) != nil) {
		return -1
	}
	skipped = optind
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + uintptr(int32(0))))) != '+' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + uintptr(int32(0))))) != '-' {
		var i int32
		for i = optind; ; i++ {
			if i >= argc || !(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(i)*8)) != nil) {
				return -1
			}
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(i)*8)))) + uintptr(int32(0))))) == '-' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(i)*8)))) + uintptr(int32(1))))) != 0 {
				break
			}
		}
		optind = i
	}
	resumed = optind
	ret = _cgos___getopt_long_core_getopt_long(argc, argv, optstring, longopts, idx, longonly)
	if resumed > skipped {
		var i int32
		var cnt int32 = optind - resumed
		for i = int32(0); i < cnt; i++ {
			_cgos_permute_getopt_long(argv, skipped, optind-int32(1))
		}
		optind = skipped + cnt
	}
	return ret
}
func _cgos___getopt_long_core_getopt_long(argc int32, argv **int8, optstring *int8, longopts *struct_option, idx *int32, longonly int32) int32 {
	optarg = (*int8)(nil)
	if longopts != nil && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)))) + uintptr(int32(0))))) == '-' && (longonly != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)))) + uintptr(int32(1))))) != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)))) + uintptr(int32(1))))) != '-' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)))) + uintptr(int32(1))))) == '-' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)))) + uintptr(int32(2))))) != 0) {
		var colon int32 = func() int32 {
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + func() uintptr {
				if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + uintptr(int32(0))))) == '+' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + uintptr(int32(0))))) == '-' {
					return 1
				} else {
					return 0
				}
			}()))) == ':' {
				return 1
			} else {
				return 0
			}
		}()
		var i int32
		var cnt int32
		var match int32
		var arg *int8
		var opt *int8
		var start *int8 = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)))) + uintptr(int32(1))))
		for cnt = func() (_cgo_ret int32) {
			_cgo_addr := &i
			*_cgo_addr = int32(0)
			return *_cgo_addr
		}(); (*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).name != nil; i++ {
			var name *int8 = (*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).name
			opt = start
			if int32(*opt) == '-' {
				*(*uintptr)(unsafe.Pointer(&opt))++
			}
			for int32(*opt) != 0 && int32(*opt) != '=' && int32(*opt) == int32(*name) {
				func() *int8 {
					*(*uintptr)(unsafe.Pointer(&name))++
					return func() (_cgo_ret *int8) {
						_cgo_addr := &opt
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}()
				}()
			}
			if int32(*opt) != 0 && int32(*opt) != '=' {
				continue
			}
			arg = opt
			match = i
			if !(*name != 0) {
				cnt = int32(1)
				break
			}
			cnt++
		}
		if cnt == int32(1) && longonly != 0 && int64(uintptr(unsafe.Pointer(arg))-uintptr(unsafe.Pointer(start))) == int64(Mblen(start, uint64(4))) {
			var l int32 = int32(uintptr(unsafe.Pointer(arg)) - uintptr(unsafe.Pointer(start)))
			for i = int32(0); *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + uintptr(i))) != 0; i++ {
				var j int32
				for j = int32(0); j < l && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(start)) + uintptr(j)))) == int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(optstring)) + uintptr(i+j)))); j++ {
				}
				if j == l {
					cnt++
					break
				}
			}
		}
		if cnt == int32(1) {
			i = match
			opt = arg
			optind++
			if int32(*opt) == '=' {
				if !((*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).has_arg != 0) {
					optopt = (*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).val
					if colon != 0 || !(opterr != 0) {
						return int32('?')
					}
					__getopt_msg(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(int32(0))*8)), (*int8)(unsafe.Pointer(&[37]int8{':', ' ', 'o', 'p', 't', 'i', 'o', 'n', ' ', 'd', 'o', 'e', 's', ' ', 'n', 'o', 't', ' ', 't', 'a', 'k', 'e', ' ', 'a', 'n', ' ', 'a', 'r', 'g', 'u', 'm', 'e', 'n', 't', ':', ' ', '\x00'})), (*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).name, Strlen((*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).name))
					return int32('?')
				}
				optarg = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(opt)) + uintptr(int32(1))))
			} else if (*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).has_arg == int32(1) {
				if !(func() (_cgo_ret *int8) {
					_cgo_addr := &optarg
					*_cgo_addr = *(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8))
					return *_cgo_addr
				}() != nil) {
					optopt = (*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).val
					if colon != 0 {
						return int32(':')
					}
					if !(opterr != 0) {
						return int32('?')
					}
					__getopt_msg(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(int32(0))*8)), (*int8)(unsafe.Pointer(&[32]int8{':', ' ', 'o', 'p', 't', 'i', 'o', 'n', ' ', 'r', 'e', 'q', 'u', 'i', 'r', 'e', 's', ' ', 'a', 'n', ' ', 'a', 'r', 'g', 'u', 'm', 'e', 'n', 't', ':', ' ', '\x00'})), (*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).name, Strlen((*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).name))
					return int32('?')
				}
				optind++
			}
			if idx != nil {
				*idx = i
			}
			if (*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).flag != nil {
				*(*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).flag = (*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).val
				return int32(0)
			}
			return (*(*struct_option)(unsafe.Pointer(uintptr(unsafe.Pointer(longopts)) + uintptr(i)*28))).val
		}
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8)))) + uintptr(int32(1))))) == '-' {
			optopt = int32(0)
			if !(colon != 0) && opterr != 0 {
				__getopt_msg(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(int32(0))*8)), func() *int8 {
					if cnt != 0 {
						return (*int8)(unsafe.Pointer(&[24]int8{':', ' ', 'o', 'p', 't', 'i', 'o', 'n', ' ', 'i', 's', ' ', 'a', 'm', 'b', 'i', 'g', 'u', 'o', 'u', 's', ':', ' ', '\x00'}))
					} else {
						return (*int8)(unsafe.Pointer(&[24]int8{':', ' ', 'u', 'n', 'r', 'e', 'c', 'o', 'g', 'n', 'i', 'z', 'e', 'd', ' ', 'o', 'p', 't', 'i', 'o', 'n', ':', ' ', '\x00'}))
					}
				}(), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8))))+uintptr(int32(2)))), Strlen((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*(**int8)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + uintptr(optind)*8))))+uintptr(int32(2))))))
			}
			optind++
			return int32('?')
		}
	}
	return getopt(argc, argv, optstring)
}
func getopt_long(argc int32, argv **int8, optstring *int8, longopts *struct_option, idx *int32) int32 {
	return _cgos___getopt_long_getopt_long(argc, argv, optstring, longopts, idx, int32(0))
}
func getopt_long_only(argc int32, argv **int8, optstring *int8, longopts *struct_option, idx *int32) int32 {
	return _cgos___getopt_long_getopt_long(argc, argv, optstring, longopts, idx, int32(1))
}
