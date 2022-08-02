package libc

func ualarm(value uint32, interval uint32) uint32 {
	var it struct_itimerval = struct_itimerval{Struct_timeval{0, int64(interval)}, Struct_timeval{0, int64(value)}}
	var it_old struct_itimerval
	setitimer(int32(0), &it, &it_old)
	return uint32(it_old.it_value.Tv_sec*int64(1000000) + it_old.it_value.Tv_usec)
}
