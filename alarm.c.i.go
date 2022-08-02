package libc

func alarm(seconds uint32) uint32 {
	var it struct_itimerval = struct_itimerval{Struct_timeval{}, Struct_timeval{int64(seconds), 0}}
	var old struct_itimerval = struct_itimerval{Struct_timeval{int64(0), 0}, Struct_timeval{}}
	setitimer(int32(0), &it, &old)
	return uint32(old.it_value.Tv_sec + func() int64 {
		if !!(old.it_value.Tv_usec != 0) {
			return 1
		} else {
			return 0
		}
	}())
}
