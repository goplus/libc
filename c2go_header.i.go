package libc

import unsafe "unsafe"

type Size_t = uint64
type Time_t = int64
type Clockid_t = int32
type Struct_timespec struct {
	Tv_sec  int64
	Xbf_0   int32
	Tv_nsec int64
	Xbf_1   int32
}
type Pthread_t = *Struct___pthread
type Pthread_once_t = int32
type Pthread_key_t = uint32
type Pthread_spinlock_t = int32
type _cgoa_1_a64l struct {
	X__attr uint32
}
type Pthread_mutexattr_t = _cgoa_1_a64l
type _cgoa_2_a64l struct {
	X__attr uint32
}
type Pthread_condattr_t = _cgoa_2_a64l
type _cgoa_3_a64l struct {
	X__attr uint32
}
type Pthread_barrierattr_t = _cgoa_3_a64l
type _cgoa_4_a64l struct {
	X__attr [2]uint32
}
type Pthread_rwlockattr_t = _cgoa_4_a64l
type Struct___sigset_t struct {
	X__bits [16]uint64
}
type Sigset_t = Struct___sigset_t
type _cgoa_6_a64l struct {
	X__i [14]int32
}
type _cgoa_5_a64l struct {
	X__u _cgoa_6_a64l
}
type Pthread_attr_t = _cgoa_5_a64l
type _cgoa_8_a64l struct {
	X__i [10]int32
}
type _cgoa_7_a64l struct {
	X__u _cgoa_8_a64l
}
type Pthread_mutex_t = _cgoa_7_a64l
type _cgoa_10_a64l struct {
	X__i [12]int32
}
type _cgoa_9_a64l struct {
	X__u _cgoa_10_a64l
}
type Pthread_cond_t = _cgoa_9_a64l
type _cgoa_12_a64l struct {
	X__i [14]int32
}
type _cgoa_11_a64l struct {
	X__u _cgoa_12_a64l
}
type Pthread_rwlock_t = _cgoa_11_a64l
type _cgoa_14_a64l struct {
	X__i [8]int32
}
type _cgoa_13_a64l struct {
	X__u _cgoa_14_a64l
}
type Pthread_barrier_t = _cgoa_13_a64l
type Pid_t = int32
type struct_sched_param struct {
	sched_priority int32
	__reserved1    int32
	__reserved2    [4]int64
	__reserved3    int32
}
type Timer_t = unsafe.Pointer
type Clock_t = int64
type Locale_t = *Struct___locale_struct
type struct_tm struct {
	tm_sec    int32
	tm_min    int32
	tm_hour   int32
	tm_mday   int32
	tm_mon    int32
	tm_year   int32
	tm_wday   int32
	tm_yday   int32
	tm_isdst  int32
	tm_gmtoff int64
	tm_zone   *int8
}
type struct_itimerspec struct {
	it_interval Struct_timespec
	it_value    Struct_timespec
}
type struct___ptcb struct {
	__f    func(unsafe.Pointer)
	__x    unsafe.Pointer
	__next *struct___ptcb
}
type Ssize_t = int64
type Off_t = int64
type FILE = Struct__IO_FILE
type Va_list = []interface {
}
type X__isoc_va_list = []interface {
}
type Union__G_fpos64_t struct {
	X__opaque [16]int8
}
type Fpos_t = Union__G_fpos64_t
type Struct___locale_struct struct {
	Cat [6]*struct___locale_map
}
type struct_tls_module struct {
	next   *struct_tls_module
	image  unsafe.Pointer
	len    uint64
	size   uint64
	align  uint64
	offset uint64
}
type struct___libc struct {
	can_do_threads  int8
	threaded        int8
	secure          int8
	need_locks      int8
	threads_minus_1 int32
	auxv            *uint64
	tls_head        *struct_tls_module
	tls_size        uint64
	tls_align       uint64
	tls_cnt         uint64
	page_size       uint64
	global_locale   Struct___locale_struct
}
type Struct___pthread struct {
	Tid       int32
	Errno_val int32
	Sys_r1    int64
	Locale    *Struct___locale_struct
}
type Wchar_t = uint32
type _cgoa_15_a64l struct {
	Quot int32
	Rem  int32
}
type Div_t = _cgoa_15_a64l
type _cgoa_16_a64l struct {
	Quot int64
	Rem  int64
}
type Ldiv_t = _cgoa_16_a64l
type _cgoa_17_a64l struct {
	Quot int64
	Rem  int64
}
type Lldiv_t = _cgoa_17_a64l
type Uintptr_t = uint64
type Intptr_t = int64
type Int8_t = int8
type Int16_t = int16
type Int32_t = int32
type Int64_t = int64
type Intmax_t = int64
type Uint8_t = uint8
type Uint16_t = uint16
type Uint32_t = uint32
type Uint64_t = uint64
type Uintmax_t = uint64
type int_fast8_t = int8
type int_fast64_t = int64
type int_least8_t = int8
type int_least16_t = int16
type int_least32_t = int32
type int_least64_t = int64
type uint_fast8_t = uint8
type uint_fast64_t = uint64
type uint_least8_t = uint8
type uint_least16_t = uint16
type uint_least32_t = uint32
type uint_least64_t = uint64
type int_fast16_t = int32
type int_fast32_t = int32
type uint_fast16_t = uint32
type uint_fast32_t = uint32

func a_fetch_and(p *int32, v int32) int32 {
	var old int32
	for {
		old = a_ll(p)
		if !!(a_sc(p, old&v) != 0) {
			break
		}
	}
	return old
}
func a_fetch_or(p *int32, v int32) int32 {
	var old int32
	for {
		old = a_ll(p)
		if !!(a_sc(p, old|v) != 0) {
			break
		}
	}
	return old
}
func a_and(p *int32, v int32) {
	a_fetch_and(p, v)
}
func a_or(p *int32, v int32) {
	a_fetch_or(p, v)
}
func a_inc(p *int32) {
	a_fetch_add(p, int32(1))
}
func a_dec(p *int32) {
	a_fetch_add(p, -1)
}
func a_store(p *int32, v int32) {
	a_swap(p, v)
}
func a_barrier() {
	var tmp int32 = int32(0)
	_ = tmp
	a_cas(&tmp, int32(0), int32(0))
}
func a_and_64(p *uint64, v uint64) {
	type _cgoa_1_ffs struct {
		v uint64
	}
	var u _cgoa_1_ffs
	u.v = v
	if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(0))*4))+uint32(1) != 0 {
		a_and((*int32)(unsafe.Pointer(p)), int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(0))*4))))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(1))*4))+uint32(1) != 0 {
		a_and((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(p))))+uintptr(int32(1))*4)), int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(1))*4))))
	}
}
func a_or_64(p *uint64, v uint64) {
	type _cgoa_2_ffs struct {
		v uint64
	}
	var u _cgoa_2_ffs
	u.v = v
	if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(0))*4)) != 0 {
		a_or((*int32)(unsafe.Pointer(p)), int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(0))*4))))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(1))*4)) != 0 {
		a_or((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(p))))+uintptr(int32(1))*4)), int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[2]uint32)(unsafe.Pointer(&u)))))) + uintptr(int32(1))*4))))
	}
}
func a_or_l(p unsafe.Pointer, v int64) {
	if 8 == 4 {
		a_or((*int32)(p), int32(v))
	} else {
		a_or_64((*uint64)(p), uint64(v))
	}
}
func a_crash() {
	*(*int8)(nil) = int8(0)
}
func a_ctz_32(x uint32) int32 {
	return int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&_cgos_debruijn32_ffs)))) + uintptr(x&-x*uint32(124511785)>>int32(27)))))
}

var _cgos_debruijn32_ffs [32]int8 = [32]int8{int8(0), int8(1), int8(23), int8(2), int8(29), int8(24), int8(19), int8(3), int8(30), int8(27), int8(25), int8(11), int8(20), int8(8), int8(4), int8(13), int8(31), int8(22), int8(28), int8(18), int8(26), int8(10), int8(7), int8(12), int8(21), int8(17), int8(9), int8(6), int8(16), int8(5), int8(15), int8(14)}

func a_ctz_64(x uint64) int32 {
	if false {
		var y uint32 = uint32(x)
		if !(y != 0) {
			y = uint32(x >> int32(32))
			return int32(32) + a_ctz_32(y)
		}
		return a_ctz_32(y)
	}
	return int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&_cgos_debruijn64_ffs)))) + uintptr(x&-x*uint64(157587932685088877)>>int32(58)))))
}

var _cgos_debruijn64_ffs [64]int8 = [64]int8{int8(0), int8(1), int8(2), int8(53), int8(3), int8(7), int8(54), int8(27), int8(4), int8(38), int8(41), int8(8), int8(34), int8(55), int8(48), int8(28), int8(62), int8(5), int8(39), int8(46), int8(44), int8(42), int8(22), int8(9), int8(24), int8(35), int8(59), int8(56), int8(49), int8(18), int8(29), int8(11), int8(63), int8(52), int8(6), int8(26), int8(37), int8(40), int8(33), int8(47), int8(61), int8(45), int8(43), int8(21), int8(23), int8(58), int8(17), int8(10), int8(51), int8(25), int8(36), int8(32), int8(60), int8(20), int8(57), int8(16), int8(50), int8(31), int8(19), int8(15), int8(30), int8(14), int8(13), int8(12)}

func a_ctz_l(x uint64) int32 {
	return func() int32 {
		if false {
			return a_ctz_32(uint32(x))
		} else {
			return a_ctz_64(uint64(x))
		}
	}()
}
func a_clz_64(x uint64) int32 {
	var y uint32
	var r int32
	if x>>int32(32) != 0 {
		func() int32 {
			y = uint32(x >> int32(32))
			return func() (_cgo_ret int32) {
				_cgo_addr := &r
				*_cgo_addr = int32(0)
				return *_cgo_addr
			}()
		}()
	} else {
		func() int32 {
			y = uint32(x)
			return func() (_cgo_ret int32) {
				_cgo_addr := &r
				*_cgo_addr = int32(32)
				return *_cgo_addr
			}()
		}()
	}
	if y>>int32(16) != 0 {
		y >>= int32(16)
	} else {
		r |= int32(16)
	}
	if y>>int32(8) != 0 {
		y >>= int32(8)
	} else {
		r |= int32(8)
	}
	if y>>int32(4) != 0 {
		y >>= int32(4)
	} else {
		r |= int32(4)
	}
	if y>>int32(2) != 0 {
		y >>= int32(2)
	} else {
		r |= int32(2)
	}
	return r | func() int32 {
		if !(y>>int32(1) != 0) {
			return 1
		} else {
			return 0
		}
	}()
}
func a_clz_32(x uint32) int32 {
	x >>= int32(1)
	x |= x >> int32(1)
	x |= x >> int32(2)
	x |= x >> int32(4)
	x |= x >> int32(8)
	x |= x >> int32(16)
	x++
	return int32(31) - a_ctz_32(x)
}

type Mode_t = uint32
type struct_flock struct {
	l_type   int16
	l_whence int16
	l_start  int64
	l_len    int64
	l_pid    int32
}
type Uid_t = uint32
type Gid_t = uint32
type Useconds_t = uint32
type Struct_winsize struct {
	Ws_row    uint16
	Ws_col    uint16
	Ws_xpixel uint16
	Ws_ypixel uint16
}
type cc_t = uint8
type speed_t = uint32
type tcflag_t = uint32
type struct_termios struct {
	c_iflag    uint32
	c_oflag    uint32
	c_cflag    uint32
	c_lflag    uint32
	c_line     uint8
	c_cc       [32]uint8
	__c_ispeed uint32
	__c_ospeed uint32
}
type Suseconds_t = int64
type Struct_timeval struct {
	Tv_sec  int64
	Tv_usec int64
}
type _cgoa_18_forkpty struct {
	__e_termination int16
	__e_exit        int16
}
type struct_utmpx struct {
	ut_type    int16
	__ut_pad1  int16
	ut_pid     int32
	ut_line    [32]int8
	ut_id      [4]int8
	ut_user    [32]int8
	ut_host    [256]int8
	ut_exit    _cgoa_18_forkpty
	ut_session int32
	__ut_pad2  int32
	ut_tv      Struct_timeval
	ut_addr_v6 [4]uint32
	__unused   [20]int8
}
type struct_lastlog struct {
	ll_time int64
	ll_line [32]int8
	ll_host [256]int8
}
type Id_t = uint32

const (
	P_ALL   int32 = int32(0)
	P_PID   int32 = int32(1)
	P_PGID  int32 = int32(2)
	P_PIDFD int32 = int32(3)
)

type idtype_t = int32
type stack_t = struct_sigaltstack
type greg_t = int32
type gregset_t = [18]int32
type struct_sigcontext struct {
	trap_no       uint64
	error_code    uint64
	oldmask       uint64
	arm_r0        uint64
	arm_r1        uint64
	arm_r2        uint64
	arm_r3        uint64
	arm_r4        uint64
	arm_r5        uint64
	arm_r6        uint64
	arm_r7        uint64
	arm_r8        uint64
	arm_r9        uint64
	arm_r10       uint64
	arm_fp        uint64
	arm_ip        uint64
	arm_sp        uint64
	arm_lr        uint64
	arm_pc        uint64
	arm_cpsr      uint64
	fault_address uint64
}
type mcontext_t = struct_sigcontext
type struct_sigaltstack struct {
	ss_sp    unsafe.Pointer
	ss_flags int32
	ss_size  uint64
}
type struct___ucontext struct {
	uc_flags    uint64
	uc_link     *struct___ucontext
	uc_stack    struct_sigaltstack
	uc_mcontext struct_sigcontext
	uc_sigmask  Struct___sigset_t
	uc_regspace [64]uint64
}
type ucontext_t = struct___ucontext
type union_sigval struct {
	sival_ptr unsafe.Pointer
}
type _cgoa_23_forkpty struct {
	si_pid int32
	si_uid uint32
}
type _cgoa_24_forkpty struct {
	si_timerid int32
	si_overrun int32
}
type _cgoa_22_forkpty struct {
	__piduid _cgoa_23_forkpty
}
type _cgoa_26_forkpty struct {
	si_status int32
	si_utime  int64
	si_stime  int64
}
type _cgoa_25_forkpty struct {
	__sigchld _cgoa_26_forkpty
}
type _cgoa_21_forkpty struct {
	__first  _cgoa_22_forkpty
	__second _cgoa_25_forkpty
}
type _cgoa_29_forkpty struct {
	si_lower unsafe.Pointer
	si_upper unsafe.Pointer
}
type _cgoa_28_forkpty struct {
	__addr_bnd _cgoa_29_forkpty
}
type _cgoa_27_forkpty struct {
	si_addr     unsafe.Pointer
	si_addr_lsb int16
	__first     _cgoa_28_forkpty
}
type _cgoa_30_forkpty struct {
	si_band int64
	si_fd   int32
}
type _cgoa_31_forkpty struct {
	si_call_addr unsafe.Pointer
	si_syscall   int32
	si_arch      uint32
}
type _cgoa_20_forkpty struct {
	__pad [112]int8
}
type _cgoa_19_forkpty struct {
	si_signo    int32
	si_errno    int32
	si_code     int32
	__si_fields _cgoa_20_forkpty
}
type siginfo_t = _cgoa_19_forkpty
type _cgoa_32_forkpty struct {
	sa_handler func(int32)
}
type struct_sigaction struct {
	__sa_handler _cgoa_32_forkpty
	sa_mask      Struct___sigset_t
	sa_flags     int32
	sa_restorer  func()
}
type _cgoa_34_forkpty struct {
	sigev_notify_function   func(union_sigval)
	sigev_notify_attributes *_cgoa_5_a64l
}
type _cgoa_33_forkpty struct {
	__pad [48]int8
}
type struct_sigevent struct {
	sigev_value  union_sigval
	sigev_signo  int32
	sigev_notify int32
	__sev_fields _cgoa_33_forkpty
}
type sig_t = func(int32)
type sig_atomic_t = int32
type fd_mask = uint64
type _cgoa_35_forkpty struct {
	fds_bits [16]uint64
}
type fd_set = _cgoa_35_forkpty
type struct_itimerval struct {
	it_interval Struct_timeval
	it_value    Struct_timeval
}
type struct_timezone struct {
	tz_minuteswest int32
	tz_dsttime     int32
}
type rlim_t = uint64
type struct_rlimit struct {
	rlim_cur uint64
	rlim_max uint64
}
type struct_rusage struct {
	ru_utime    Struct_timeval
	ru_stime    Struct_timeval
	ru_maxrss   int64
	ru_ixrss    int64
	ru_idrss    int64
	ru_isrss    int64
	ru_minflt   int64
	ru_majflt   int64
	ru_nswap    int64
	ru_inblock  int64
	ru_oublock  int64
	ru_msgsnd   int64
	ru_msgrcv   int64
	ru_nsignals int64
	ru_nvcsw    int64
	ru_nivcsw   int64
	__reserved  [16]int64
}
type struct_cpu_set_t struct {
	__bits [16]uint64
}
type cpu_set_t = struct_cpu_set_t

func __CPU_AND_S(__size uint64, __dest *struct_cpu_set_t, __src1 *struct_cpu_set_t, __src2 *struct_cpu_set_t) {
	var __i uint64
	for __i = uint64(0); __i < __size/8; __i++ {
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(__dest)))) + uintptr(__i)*8)) = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(__src1)))) + uintptr(__i)*8)) & *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(__src2)))) + uintptr(__i)*8))
	}
}
func __CPU_OR_S(__size uint64, __dest *struct_cpu_set_t, __src1 *struct_cpu_set_t, __src2 *struct_cpu_set_t) {
	var __i uint64
	for __i = uint64(0); __i < __size/8; __i++ {
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(__dest)))) + uintptr(__i)*8)) = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(__src1)))) + uintptr(__i)*8)) | *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(__src2)))) + uintptr(__i)*8))
	}
}
func __CPU_XOR_S(__size uint64, __dest *struct_cpu_set_t, __src1 *struct_cpu_set_t, __src2 *struct_cpu_set_t) {
	var __i uint64
	for __i = uint64(0); __i < __size/8; __i++ {
		*(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(__dest)))) + uintptr(__i)*8)) = *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(__src1)))) + uintptr(__i)*8)) ^ *(*uint64)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint64)(unsafe.Pointer(__src2)))) + uintptr(__i)*8))
	}
}

type Cookie_read_function_t = func(unsafe.Pointer, *int8, uint64) int64
type Cookie_write_function_t = func(unsafe.Pointer, *int8, uint64) int64
type Cookie_seek_function_t = func(unsafe.Pointer, *int64, int32) int32
type Cookie_close_function_t = func(unsafe.Pointer) int32
type Struct__IO_cookie_io_functions_t struct {
	Read  func(unsafe.Pointer, *int8, uint64) int64
	Write func(unsafe.Pointer, *int8, uint64) int64
	Seek  func(unsafe.Pointer, *int64, int32) int32
	Close func(unsafe.Pointer) int32
}
type Cookie_io_functions_t = Struct__IO_cookie_io_functions_t
type Nlink_t = uint32
type Ino_t = uint64
type Dev_t = uint64
type Blksize_t = int64
type Blkcnt_t = int64
type _cgoa_18_get_current_dir_name struct {
	tv_sec  int64
	tv_nsec int64
}
type struct_stat struct {
	st_dev             uint64
	__st_dev_padding   int32
	__st_ino_truncated int64
	st_mode            uint32
	st_nlink           uint32
	st_uid             uint32
	st_gid             uint32
	st_rdev            uint64
	__st_rdev_padding  int32
	st_size            int64
	st_blksize         int64
	st_blocks          int64
	__st_atim32        _cgoa_18_get_current_dir_name
	__st_mtim32        _cgoa_18_get_current_dir_name
	__st_ctim32        _cgoa_18_get_current_dir_name
	st_ino             uint64
	st_atim            Struct_timespec
	st_mtim            Struct_timespec
	st_ctim            Struct_timespec
}
type Elf32_Half = uint16
type Elf64_Half = uint16
type Elf32_Word = uint32
type Elf32_Sword = int32
type Elf64_Word = uint32
type Elf64_Sword = int32
type Elf32_Xword = uint64
type Elf32_Sxword = int64
type Elf64_Xword = uint64
type Elf64_Sxword = int64
type Elf32_Addr = uint32
type Elf64_Addr = uint64
type Elf32_Off = uint32
type Elf64_Off = uint64
type Elf32_Section = uint16
type Elf64_Section = uint16
type Elf32_Versym = uint16
type Elf64_Versym = uint16
type _cgoa_1_getauxval struct {
	e_ident     [16]uint8
	e_type      uint16
	e_machine   uint16
	e_version   uint32
	e_entry     uint32
	e_phoff     uint32
	e_shoff     uint32
	e_flags     uint32
	e_ehsize    uint16
	e_phentsize uint16
	e_phnum     uint16
	e_shentsize uint16
	e_shnum     uint16
	e_shstrndx  uint16
}
type Elf32_Ehdr = _cgoa_1_getauxval
type _cgoa_2_getauxval struct {
	e_ident     [16]uint8
	e_type      uint16
	e_machine   uint16
	e_version   uint32
	e_entry     uint64
	e_phoff     uint64
	e_shoff     uint64
	e_flags     uint32
	e_ehsize    uint16
	e_phentsize uint16
	e_phnum     uint16
	e_shentsize uint16
	e_shnum     uint16
	e_shstrndx  uint16
}
type Elf64_Ehdr = _cgoa_2_getauxval
type _cgoa_3_getauxval struct {
	sh_name      uint32
	sh_type      uint32
	sh_flags     uint32
	sh_addr      uint32
	sh_offset    uint32
	sh_size      uint32
	sh_link      uint32
	sh_info      uint32
	sh_addralign uint32
	sh_entsize   uint32
}
type Elf32_Shdr = _cgoa_3_getauxval
type _cgoa_4_getauxval struct {
	sh_name      uint32
	sh_type      uint32
	sh_flags     uint64
	sh_addr      uint64
	sh_offset    uint64
	sh_size      uint64
	sh_link      uint32
	sh_info      uint32
	sh_addralign uint64
	sh_entsize   uint64
}
type Elf64_Shdr = _cgoa_4_getauxval
type _cgoa_5_getauxval struct {
	ch_type      uint32
	ch_size      uint32
	ch_addralign uint32
}
type Elf32_Chdr = _cgoa_5_getauxval
type _cgoa_6_getauxval struct {
	ch_type      uint32
	ch_reserved  uint32
	ch_size      uint64
	ch_addralign uint64
}
type Elf64_Chdr = _cgoa_6_getauxval
type _cgoa_7_getauxval struct {
	st_name  uint32
	st_value uint32
	st_size  uint32
	st_info  uint8
	st_other uint8
	st_shndx uint16
}
type Elf32_Sym = _cgoa_7_getauxval
type _cgoa_8_getauxval struct {
	st_name  uint32
	st_info  uint8
	st_other uint8
	st_shndx uint16
	st_value uint64
	st_size  uint64
}
type Elf64_Sym = _cgoa_8_getauxval
type _cgoa_9_getauxval struct {
	si_boundto uint16
	si_flags   uint16
}
type Elf32_Syminfo = _cgoa_9_getauxval
type _cgoa_10_getauxval struct {
	si_boundto uint16
	si_flags   uint16
}
type Elf64_Syminfo = _cgoa_10_getauxval
type _cgoa_11_getauxval struct {
	r_offset uint32
	r_info   uint32
}
type Elf32_Rel = _cgoa_11_getauxval
type _cgoa_12_getauxval struct {
	r_offset uint64
	r_info   uint64
}
type Elf64_Rel = _cgoa_12_getauxval
type _cgoa_13_getauxval struct {
	r_offset uint32
	r_info   uint32
	r_addend int32
}
type Elf32_Rela = _cgoa_13_getauxval
type _cgoa_14_getauxval struct {
	r_offset uint64
	r_info   uint64
	r_addend int64
}
type Elf64_Rela = _cgoa_14_getauxval
type _cgoa_15_getauxval struct {
	p_type   uint32
	p_offset uint32
	p_vaddr  uint32
	p_paddr  uint32
	p_filesz uint32
	p_memsz  uint32
	p_flags  uint32
	p_align  uint32
}
type Elf32_Phdr = _cgoa_15_getauxval
type _cgoa_16_getauxval struct {
	p_type   uint32
	p_flags  uint32
	p_offset uint64
	p_vaddr  uint64
	p_paddr  uint64
	p_filesz uint64
	p_memsz  uint64
	p_align  uint64
}
type Elf64_Phdr = _cgoa_16_getauxval
type _cgoa_18_getauxval struct {
	d_val uint32
}
type _cgoa_17_getauxval struct {
	d_tag int32
	d_un  _cgoa_18_getauxval
}
type Elf32_Dyn = _cgoa_17_getauxval
type _cgoa_20_getauxval struct {
	d_val uint64
}
type _cgoa_19_getauxval struct {
	d_tag int64
	d_un  _cgoa_20_getauxval
}
type Elf64_Dyn = _cgoa_19_getauxval
type _cgoa_21_getauxval struct {
	vd_version uint16
	vd_flags   uint16
	vd_ndx     uint16
	vd_cnt     uint16
	vd_hash    uint32
	vd_aux     uint32
	vd_next    uint32
}
type Elf32_Verdef = _cgoa_21_getauxval
type _cgoa_22_getauxval struct {
	vd_version uint16
	vd_flags   uint16
	vd_ndx     uint16
	vd_cnt     uint16
	vd_hash    uint32
	vd_aux     uint32
	vd_next    uint32
}
type Elf64_Verdef = _cgoa_22_getauxval
type _cgoa_23_getauxval struct {
	vda_name uint32
	vda_next uint32
}
type Elf32_Verdaux = _cgoa_23_getauxval
type _cgoa_24_getauxval struct {
	vda_name uint32
	vda_next uint32
}
type Elf64_Verdaux = _cgoa_24_getauxval
type _cgoa_25_getauxval struct {
	vn_version uint16
	vn_cnt     uint16
	vn_file    uint32
	vn_aux     uint32
	vn_next    uint32
}
type Elf32_Verneed = _cgoa_25_getauxval
type _cgoa_26_getauxval struct {
	vn_version uint16
	vn_cnt     uint16
	vn_file    uint32
	vn_aux     uint32
	vn_next    uint32
}
type Elf64_Verneed = _cgoa_26_getauxval
type _cgoa_27_getauxval struct {
	vna_hash  uint32
	vna_flags uint16
	vna_other uint16
	vna_name  uint32
	vna_next  uint32
}
type Elf32_Vernaux = _cgoa_27_getauxval
type _cgoa_28_getauxval struct {
	vna_hash  uint32
	vna_flags uint16
	vna_other uint16
	vna_name  uint32
	vna_next  uint32
}
type Elf64_Vernaux = _cgoa_28_getauxval
type _cgoa_30_getauxval struct {
	a_val uint32
}
type _cgoa_29_getauxval struct {
	a_type uint32
	a_un   _cgoa_30_getauxval
}
type Elf32_auxv_t = _cgoa_29_getauxval
type _cgoa_32_getauxval struct {
	a_val uint64
}
type _cgoa_31_getauxval struct {
	a_type uint64
	a_un   _cgoa_32_getauxval
}
type Elf64_auxv_t = _cgoa_31_getauxval
type _cgoa_33_getauxval struct {
	n_namesz uint32
	n_descsz uint32
	n_type   uint32
}
type Elf32_Nhdr = _cgoa_33_getauxval
type _cgoa_34_getauxval struct {
	n_namesz uint32
	n_descsz uint32
	n_type   uint32
}
type Elf64_Nhdr = _cgoa_34_getauxval
type _cgoa_35_getauxval struct {
	m_value   uint64
	m_info    uint32
	m_poffset uint32
	m_repeat  uint16
	m_stride  uint16
}
type Elf32_Move = _cgoa_35_getauxval
type _cgoa_36_getauxval struct {
	m_value   uint64
	m_info    uint64
	m_poffset uint64
	m_repeat  uint16
	m_stride  uint16
}
type Elf64_Move = _cgoa_36_getauxval
type _cgoa_38_getauxval struct {
	gt_current_g_value uint32
	gt_unused          uint32
}
type _cgoa_39_getauxval struct {
	gt_g_value uint32
	gt_bytes   uint32
}
type _cgoa_37_getauxval struct {
	gt_header _cgoa_38_getauxval
}
type Elf32_gptab = _cgoa_37_getauxval
type _cgoa_40_getauxval struct {
	ri_gprmask  uint32
	ri_cprmask  [4]uint32
	ri_gp_value int32
}
type Elf32_RegInfo = _cgoa_40_getauxval
type _cgoa_41_getauxval struct {
	kind    uint8
	size    uint8
	section uint16
	info    uint32
}
type Elf_Options = _cgoa_41_getauxval
type _cgoa_42_getauxval struct {
	hwp_flags1 uint32
	hwp_flags2 uint32
}
type Elf_Options_Hw = _cgoa_42_getauxval
type _cgoa_43_getauxval struct {
	l_name       uint32
	l_time_stamp uint32
	l_checksum   uint32
	l_version    uint32
	l_flags      uint32
}
type Elf32_Lib = _cgoa_43_getauxval
type _cgoa_44_getauxval struct {
	l_name       uint32
	l_time_stamp uint32
	l_checksum   uint32
	l_version    uint32
	l_flags      uint32
}
type Elf64_Lib = _cgoa_44_getauxval
type Elf32_Conflict = uint32
type _cgoa_45_getauxval struct {
	version   uint16
	isa_level uint8
	isa_rev   uint8
	gpr_size  uint8
	cpr1_size uint8
	cpr2_size uint8
	fp_abi    uint8
	isa_ext   uint32
	ases      uint32
	flags1    uint32
	flags2    uint32
}
type Elf_MIPS_ABIFlags_v0 = _cgoa_45_getauxval

const (
	Val_GNU_MIPS_ABI_FP_ANY    int32 = int32(0)
	Val_GNU_MIPS_ABI_FP_DOUBLE int32 = int32(1)
	Val_GNU_MIPS_ABI_FP_SINGLE int32 = int32(2)
	Val_GNU_MIPS_ABI_FP_SOFT   int32 = int32(3)
	Val_GNU_MIPS_ABI_FP_OLD_64 int32 = int32(4)
	Val_GNU_MIPS_ABI_FP_XX     int32 = int32(5)
	Val_GNU_MIPS_ABI_FP_64     int32 = int32(6)
	Val_GNU_MIPS_ABI_FP_64A    int32 = int32(7)
	Val_GNU_MIPS_ABI_FP_MAX    int32 = int32(7)
)

type struct_utsname struct {
	sysname    [65]int8
	nodename   [65]int8
	release    [65]int8
	version    [65]int8
	machine    [65]int8
	domainname [65]int8
}
type Wint_t = uint32
type Wctype_t = uint64
type Struct___mbstate_t struct {
	X__opaque1 uint32
	X__opaque2 uint32
}
type Mbstate_t = Struct___mbstate_t
type struct_lconv struct {
	decimal_point      *int8
	thousands_sep      *int8
	grouping           *int8
	int_curr_symbol    *int8
	currency_symbol    *int8
	mon_decimal_point  *int8
	mon_thousands_sep  *int8
	mon_grouping       *int8
	positive_sign      *int8
	negative_sign      *int8
	int_frac_digits    int8
	frac_digits        int8
	p_cs_precedes      int8
	p_sep_by_space     int8
	n_cs_precedes      int8
	n_sep_by_space     int8
	p_sign_posn        int8
	n_sign_posn        int8
	int_p_cs_precedes  int8
	int_p_sep_by_space int8
	int_n_cs_precedes  int8
	int_n_sep_by_space int8
	int_p_sign_posn    int8
	int_n_sign_posn    int8
}
type struct___locale_map struct {
	map_     unsafe.Pointer
	map_size uint64
	name     [24]int8
	next     *struct___locale_map
}
type syscall_arg_t = int64

func __alt_socketcall(sys int32, sock int32, cp int32, a int64, b int64, c int64, d int64, e int64, f int64) int64 {
	var r int64
	if cp != 0 {
		r = __syscall_cp(int64(sys), int64(a), int64(b), int64(c), int64(d), int64(e), int64(f))
	} else {
		r = __syscall6(int64(sys), int64(a), int64(b), int64(c), int64(d), int64(e), int64(f))
	}
	if r != int64(-38) {
		return r
	}
	return r
}

type Struct__IO_FILE struct {
	Flags        uint32
	Rpos         *uint8
	Rend         *uint8
	Close        func(*Struct__IO_FILE) int32
	Wend         *uint8
	Wpos         *uint8
	Mustbezero_1 *uint8
	Wbase        *uint8
	Read         func(*Struct__IO_FILE, *uint8, uint64) uint64
	Write        func(*Struct__IO_FILE, *uint8, uint64) uint64
	Seek         func(*Struct__IO_FILE, int64, int32) int64
	Buf          *uint8
	Buf_size     uint64
	Prev         *Struct__IO_FILE
	Next         *Struct__IO_FILE
	Fd           int32
	Pipe_pid     int32
	Lockcount    int64
	Mode         int32
	Lock         int32
	Lbf          int32
	Cookie       unsafe.Pointer
	Off          int64
	Getln_buf    *int8
	Mustbezero_2 unsafe.Pointer
	Shend        *uint8
	Shlim        int64
	Shcnt        int64
	Prev_locked  *Struct__IO_FILE
	Next_locked  *Struct__IO_FILE
	Locale       *Struct___locale_struct
}
type _cgoa_1_getopt_long struct {
	X__ll int64
	X__ld float64
}
type Max_align_t = _cgoa_1_getopt_long
type Ptrdiff_t = int64
type struct_option struct {
	name    *int8
	has_arg int32
	flag    *int32
	val     int32
}
type struct_group struct {
	gr_name   *int8
	gr_passwd *int8
	gr_gid    uint32
	gr_mem    **int8
}
type struct_mntent struct {
	mnt_fsname *int8
	mnt_dir    *int8
	mnt_type   *int8
	mnt_opts   *int8
	mnt_freq   int32
	mnt_passno int32
}
type struct_FTW struct {
	base  int32
	level int32
}
type struct_dirent struct {
	d_ino    uint64
	d_off    int64
	d_reclen uint16
	d_type   uint8
	d_name   [256]int8
}
type DIR = struct___dirstream
type _cgoa_18_wordexp struct {
	we_wordc uint64
	we_wordv **int8
	we_offs  uint64
}
type wordexp_t = _cgoa_18_wordexp
type Float_t = float32
type Double_t = float64

func X__FLOAT_BITS(__f float32) uint32 {
	type _cgoa_18___cos struct {
		__f float32
	}
	var __u _cgoa_18___cos
	__u.__f = __f
	return *(*uint32)(unsafe.Pointer(&__u))
}
func X__DOUBLE_BITS(__f float64) uint64 {
	type _cgoa_19___cos struct {
		__f float64
	}
	var __u _cgoa_19___cos
	__u.__f = __f
	return *(*uint64)(unsafe.Pointer(&__u))
}
func X__islessf(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x < __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__isless(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(__x)&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(__y)&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x < __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__islessl(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(__x) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(__y) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x < __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__islessequalf(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x <= __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__islessequal(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(__x)&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(__y)&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x <= __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__islessequall(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(__x) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(__y) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x <= __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__islessgreaterf(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x != __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__islessgreater(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(__x)&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(__y)&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x != __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__islessgreaterl(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(__x) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(__y) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x != __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__isgreaterf(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x > __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__isgreater(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(__x)&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(__y)&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x > __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__isgreaterl(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(__x) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(__y) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x > __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__isgreaterequalf(__x float32, __y float32) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 4 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(__x)&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 4 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 4 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(__y)&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 4 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x >= __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__isgreaterequal(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(__x)&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(float64(__x)) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(__y)&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(float64(__y)) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x >= __y {
			return 1
		} else {
			return 0
		}
	}()
}
func X__isgreaterequall(__x float64, __y float64) int32 {
	return func() int32 {
		if !(func() int32 {
			if func() int32 {
				if 8 == 4 {
					return func() int32 {
						if X__FLOAT_BITS(float32(__x))&uint32(2147483647) > uint32(2139095040) {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if 8 == 8 {
							return func() int32 {
								if X__DOUBLE_BITS(float64(__x))&9223372036854775807 > 9218868437227405312 {
									return 1
								} else {
									return 0
								}
							}()
						} else {
							return func() int32 {
								if X__fpclassifyl(__x) == int32(0) {
									return 1
								} else {
									return 0
								}
							}()
						}
					}()
				}
			}() != 0 {
				return func() int32 {
					func() int {
						_ = __y
						return 0
					}()
					return int32(1)
				}()
			} else {
				return func() int32 {
					if 8 == 4 {
						return func() int32 {
							if X__FLOAT_BITS(float32(__y))&uint32(2147483647) > uint32(2139095040) {
								return 1
							} else {
								return 0
							}
						}()
					} else {
						return func() int32 {
							if 8 == 8 {
								return func() int32 {
									if X__DOUBLE_BITS(float64(__y))&9223372036854775807 > 9218868437227405312 {
										return 1
									} else {
										return 0
									}
								}()
							} else {
								return func() int32 {
									if X__fpclassifyl(__y) == int32(0) {
										return 1
									} else {
										return 0
									}
								}()
							}
						}()
					}
				}()
			}
		}() != 0) && __x >= __y {
			return 1
		} else {
			return 0
		}
	}()
}
func __bswap16(__x uint16) uint16 {
	return uint16(int32(__x)<<int32(8) | int32(__x)>>int32(8))
}
func __bswap32(__x uint32) uint32 {
	return __x>>int32(24) | __x>>int32(8)&uint32(65280) | __x<<int32(8)&uint32(16711680) | __x<<int32(24)
}
func __bswap64(__x uint64) uint64 {
	return (uint64(__bswap32(uint32(__x)))+uint64(0))<<int32(32) | uint64(__bswap32(uint32(__x>>int32(32))))
}
func eval_as_float(x float32) float32 {
	var y float32 = x
	return y
}
func eval_as_double(x float64) float64 {
	var y float64 = x
	return y
}
func fp_barrierf(x float32) float32 {
	var y float32 = x
	_ = y
	return y
}
func fp_barrier(x float64) float64 {
	var y float64 = x
	_ = y
	return y
}
func fp_barrierl(x float64) float64 {
	var y float64 = x
	_ = y
	return y
}
func fp_force_evalf(x float32) {
	var y float32
	_ = y
	y = x
}
func fp_force_eval(x float64) {
	var y float64
	_ = y
	y = x
}
func fp_force_evall(x float64) {
	var y float64
	_ = y
	y = x
}

type struct_exp_data struct {
	invln2N    float64
	shift      float64
	negln2hiN  float64
	negln2loN  float64
	poly       [4]float64
	exp2_shift float64
	exp2_poly  [5]float64
	tab        [256]uint64
}
type struct_exp2f_data struct {
	tab           [32]uint64
	shift_scaled  float64
	poly          [3]float64
	shift         float64
	invln2_scaled float64
	poly_scaled   [3]float64
}
type fexcept_t = uint64
type _cgoa_1_fmaf struct {
	__cw uint64
}
type fenv_t = _cgoa_1_fmaf
type _cgoa_18_log struct {
	invc float64
	logc float64
}
type _cgoa_19_log struct {
	chi float64
	clo float64
}
type struct_log_data struct {
	ln2hi float64
	ln2lo float64
	poly  [5]float64
	poly1 [11]float64
	tab   [128]_cgoa_18_log
	tab2  [128]_cgoa_19_log
}
type _cgoa_18_log2 struct {
	invc float64
	logc float64
}
type _cgoa_19_log2 struct {
	chi float64
	clo float64
}
type struct_log2_data struct {
	invln2hi float64
	invln2lo float64
	poly     [6]float64
	poly1    [10]float64
	tab      [64]_cgoa_18_log2
	tab2     [64]_cgoa_19_log2
}
type _cgoa_18_log2f struct {
	invc float64
	logc float64
}
type struct_log2f_data struct {
	tab  [16]_cgoa_18_log2f
	poly [4]float64
}
type _cgoa_18_logf struct {
	invc float64
	logc float64
}
type struct_logf_data struct {
	tab  [16]_cgoa_18_logf
	ln2  float64
	poly [3]float64
}
type _cgoa_18_pow struct {
	invc     float64
	pad      float64
	logc     float64
	logctail float64
}
type struct_pow_log_data struct {
	ln2hi float64
	ln2lo float64
	poly  [7]float64
	tab   [128]_cgoa_18_pow
}
type _cgoa_18_powf struct {
	invc float64
	logc float64
}
type struct_powf_log2_data struct {
	tab  [16]_cgoa_18_powf
	poly [5]float64
}
type Struct_iovec struct {
	Iov_base unsafe.Pointer
	Iov_len  uint64
}

func locking_getc(f *Struct__IO_FILE) int32 {
	if a_cas(&f.Lock, int32(0), 1073741823) != 0 {
		__lockfile(f)
	}
	var c int32 = func() int32 {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __uflow(f)
		}
	}()
	if a_swap(&f.Lock, int32(0))&int32(1073741824) != 0 {
		__wake(unsafe.Pointer(&f.Lock), int32(1), int32(1))
	}
	return c
}
func do_getc(f *Struct__IO_FILE) int32 {
	var l int32 = f.Lock
	if l < int32(0) || l != 0 && l&-1073741825 == __pthread_self().Tid {
		return func() int32 {
			if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) {
				return int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return __uflow(f)
			}
		}()
	}
	return locking_getc(f)
}

type _cgoa_19_fmemopen struct {
	quot int64
	rem  int64
}
type imaxdiv_t = _cgoa_19_fmemopen
type struct_file_handle struct {
	handle_bytes uint32
	handle_type  int32
	f_handle     [0]uint8
}
type struct_f_owner_ex struct {
	type_ int32
	pid   int32
}

func locking_putc(c int32, f *Struct__IO_FILE) int32 {
	if a_cas(&f.Lock, int32(0), 1073741823) != 0 {
		__lockfile(f)
	}
	c = func() int32 {
		if int32(uint8(c)) != f.Lbf && uintptr(unsafe.Pointer(f.Wpos)) != uintptr(unsafe.Pointer(f.Wend)) {
			return int32(func() (_cgo_ret uint8) {
				_cgo_addr := &*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Wpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
				*_cgo_addr = uint8(c)
				return *_cgo_addr
			}())
		} else {
			return __overflow(f, int32(uint8(c)))
		}
	}()
	if a_swap(&f.Lock, int32(0))&int32(1073741824) != 0 {
		__wake(unsafe.Pointer(&f.Lock), int32(1), int32(1))
	}
	return c
}
func do_putc(c int32, f *Struct__IO_FILE) int32 {
	var l int32 = f.Lock
	if l < int32(0) || l != 0 && l&-1073741825 == __pthread_self().Tid {
		return func() int32 {
			if int32(uint8(c)) != f.Lbf && uintptr(unsafe.Pointer(f.Wpos)) != uintptr(unsafe.Pointer(f.Wend)) {
				return int32(func() (_cgo_ret uint8) {
					_cgo_addr := &*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Wpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}()
					*_cgo_addr = uint8(c)
					return *_cgo_addr
				}())
			} else {
				return __overflow(f, int32(uint8(c)))
			}
		}()
	}
	return locking_putc(c, f)
}
func X__isspace(_c int32) int32 {
	return func() int32 {
		if _c == ' ' || uint32(_c)-uint32('\t') < uint32(5) {
			return 1
		} else {
			return 0
		}
	}()
}

type struct_kstat struct {
	st_dev             uint64
	__st_dev_padding   int32
	__st_ino_truncated int64
	st_mode            uint32
	st_nlink           uint32
	st_uid             uint32
	st_gid             uint32
	st_rdev            uint64
	__st_rdev_padding  int32
	st_size            int64
	st_blksize         int64
	st_blocks          int64
	st_atime_sec       int64
	st_atime_nsec      int64
	st_mtime_sec       int64
	st_mtime_nsec      int64
	st_ctime_sec       int64
	st_ctime_nsec      int64
	st_ino             uint64
}
type wctrans_t = *int32

var _cgos_tab_towctrans [2666]uint8 = [2666]uint8{uint8(7), uint8(8), uint8(9), uint8(10), uint8(11), uint8(12), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(13), uint8(6), uint8(6), uint8(14), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(15), uint8(16), uint8(17), uint8(18), uint8(6), uint8(19), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(20), uint8(21), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(22), uint8(23), uint8(6), uint8(6), uint8(6), uint8(24), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(25), uint8(6), uint8(6), uint8(6), uint8(6), uint8(26), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(27), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(28), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(29), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(30), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(6), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(36), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(1), uint8(0), uint8(84), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(24), uint8(0), uint8(0), uint8(0), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(7), uint8(43), uint8(43), uint8(91), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(74), uint8(86), uint8(86), uint8(5), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(36), uint8(80), uint8(121), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(56), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(78), uint8(49), uint8(2), uint8(78), uint8(13), uint8(13), uint8(78), uint8(3), uint8(78), uint8(0), uint8(36), uint8(110), uint8(0), uint8(78), uint8(49), uint8(38), uint8(110), uint8(81), uint8(78), uint8(36), uint8(80), uint8(78), uint8(57), uint8(20), uint8(129), uint8(27), uint8(29), uint8(29), uint8(83), uint8(49), uint8(80), uint8(49), uint8(80), uint8(13), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(27), uint8(83), uint8(36), uint8(80), uint8(49), uint8(2), uint8(92), uint8(123), uint8(92), uint8(123), uint8(92), uint8(123), uint8(92), uint8(123), uint8(92), uint8(123), uint8(20), uint8(121), uint8(92), uint8(123), uint8(92), uint8(123), uint8(92), uint8(45), uint8(43), uint8(73), uint8(3), uint8(72), uint8(3), uint8(120), uint8(92), uint8(123), uint8(20), uint8(0), uint8(150), uint8(10), uint8(1), uint8(43), uint8(40), uint8(6), uint8(6), uint8(0), uint8(42), uint8(6), uint8(42), uint8(42), uint8(43), uint8(7), uint8(187), uint8(181), uint8(43), uint8(30), uint8(0), uint8(43), uint8(7), uint8(43), uint8(43), uint8(43), uint8(1), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(1), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(42), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(205), uint8(70), uint8(205), uint8(43), uint8(0), uint8(37), uint8(43), uint8(7), uint8(1), uint8(6), uint8(1), uint8(85), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(85), uint8(86), uint8(86), uint8(2), uint8(36), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(21), uint8(129), uint8(129), uint8(129), uint8(0), uint8(0), uint8(43), uint8(0), uint8(178), uint8(209), uint8(178), uint8(209), uint8(178), uint8(209), uint8(178), uint8(209), uint8(0), uint8(0), uint8(205), uint8(204), uint8(1), uint8(0), uint8(215), uint8(215), uint8(215), uint8(215), uint8(215), uint8(131), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(28), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(2), uint8(0), uint8(0), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(78), uint8(49), uint8(80), uint8(49), uint8(80), uint8(78), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(2), uint8(135), uint8(166), uint8(135), uint8(166), uint8(135), uint8(166), uint8(135), uint8(166), uint8(135), uint8(166), uint8(135), uint8(166), uint8(135), uint8(166), uint8(135), uint8(166), uint8(42), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(0), uint8(0), uint8(0), uint8(84), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(84), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(12), uint8(0), uint8(12), uint8(42), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(7), uint8(42), uint8(1), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(42), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(86), uint8(86), uint8(108), uint8(129), uint8(21), uint8(0), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(7), uint8(108), uint8(3), uint8(65), uint8(43), uint8(43), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(44), uint8(86), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(1), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(12), uint8(108), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(6), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(86), uint8(122), uint8(158), uint8(38), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(37), uint8(6), uint8(1), uint8(43), uint8(43), uint8(79), uint8(86), uint8(86), uint8(44), uint8(43), uint8(127), uint8(86), uint8(86), uint8(57), uint8(43), uint8(43), uint8(85), uint8(86), uint8(86), uint8(43), uint8(43), uint8(79), uint8(86), uint8(86), uint8(44), uint8(43), uint8(127), uint8(86), uint8(86), uint8(129), uint8(55), uint8(117), uint8(91), uint8(123), uint8(92), uint8(43), uint8(43), uint8(79), uint8(86), uint8(86), uint8(2), uint8(172), uint8(4), uint8(0), uint8(0), uint8(57), uint8(43), uint8(43), uint8(85), uint8(86), uint8(86), uint8(43), uint8(43), uint8(79), uint8(86), uint8(86), uint8(44), uint8(43), uint8(43), uint8(86), uint8(86), uint8(50), uint8(19), uint8(129), uint8(87), uint8(0), uint8(111), uint8(129), uint8(126), uint8(201), uint8(215), uint8(126), uint8(45), uint8(129), uint8(129), uint8(14), uint8(126), uint8(57), uint8(127), uint8(111), uint8(87), uint8(0), uint8(129), uint8(129), uint8(126), uint8(21), uint8(0), uint8(126), uint8(3), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(7), uint8(43), uint8(36), uint8(43), uint8(151), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(42), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(128), uint8(129), uint8(129), uint8(129), uint8(129), uint8(57), uint8(187), uint8(42), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(1), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(129), uint8(201), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(172), uint8(208), uint8(13), uint8(0), uint8(78), uint8(49), uint8(2), uint8(180), uint8(193), uint8(193), uint8(215), uint8(215), uint8(36), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(215), uint8(215), uint8(83), uint8(193), uint8(71), uint8(212), uint8(215), uint8(215), uint8(215), uint8(5), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(7), uint8(1), uint8(0), uint8(1), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(78), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(13), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(36), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(49), uint8(80), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(121), uint8(92), uint8(123), uint8(92), uint8(123), uint8(79), uint8(123), uint8(92), uint8(123), uint8(92), uint8(123), uint8(92), uint8(123), uint8(92), uint8(123), uint8(92), uint8(123), uint8(92), uint8(123), uint8(92), uint8(123), uint8(92), uint8(123), uint8(92), uint8(123), uint8(92), uint8(45), uint8(43), uint8(43), uint8(121), uint8(20), uint8(92), uint8(123), uint8(92), uint8(45), uint8(121), uint8(42), uint8(92), uint8(39), uint8(92), uint8(123), uint8(92), uint8(123), uint8(92), uint8(123), uint8(164), uint8(0), uint8(10), uint8(180), uint8(92), uint8(123), uint8(92), uint8(123), uint8(79), uint8(3), uint8(42), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(1), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(72), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(42), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(7), uint8(0), uint8(72), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(2), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(85), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(14), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(36), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(7), uint8(0), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(36), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(7), uint8(0), uint8(0), uint8(0), uint8(0), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(42), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(14), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(42), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(14), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(43), uint8(85), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(86), uint8(14), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0)}
var _cgos_rules_towctrans [240]int32 = [240]int32{int32(0), int32(8193), -8192, int32(1949440), int32(190208), int32(30976), int32(9218), int32(257), -256, int32(0), int32(513), -512, -50943, -59392, -30975, -76800, int32(49920), int32(53761), int32(52737), int32(52481), int32(20225), int32(51713), int32(51969), int32(52993), int32(24832), int32(54017), int32(53505), int32(41728), int32(54529), int32(33280), int32(54785), int32(55809), int32(55553), int32(56065), int32(14336), int32(3), -20224, -24831, -14335, int32(2369538), int32(0), int32(257), -256, -52480, -55808, -33279, int32(2763521), -41727, int32(2762753), int32(2768640), -49919, int32(17665), int32(18177), int32(2760448), int32(2759680), int32(2760192), -53760, -52736, -51712, -51968, int32(10833664), int32(10832640), -52992, int32(10823680), int32(10830848), -53504, -54016, int32(2750208), int32(10830080), int32(2751744), -54528, -54784, int32(2746112), int32(10830592), int32(10824192), -17664, -55552, -18176, -56064, int32(10818816), int32(10818048), int32(4989954), int32(0), int32(8193), -8192, int32(257), -256, int32(21504), int32(29697), int32(9729), int32(9473), int32(16385), int32(16129), -9728, -9472, -7936, -16384, -16128, int32(2049), -15872, -14592, -12032, -13824, -2048, -22016, -20480, int32(1792), -29696, -15359, -24576, -1791, int32(7346690), int32(257), -256, int32(8193), -8192, int32(20481), int32(3841), -3840, int32(0), int32(12289), -12288, int32(257), -256, int32(0), int32(770048), int32(1859585), int32(0), int32(9949185), int32(2049), -2048, int32(9045250), int32(0), -770047, -1597952, int32(9028096), -1582336, -1601024, -1600768, -1598464, -1598208, -1596416, int32(0), int32(9058304), int32(9044992), int32(976384), int32(257), -256, int32(0), -15104, -1949439, int32(9379074), int32(2048), -2047, int32(0), int32(22016), -22015, int32(18944), int32(25600), int32(32768), int32(28672), int32(32256), int32(2304), -18943, -2303, -1844480, -25599, -28671, -32767, -32255, int32(11273474), int32(0), int32(4097), -4096, int32(7169), int32(257), -1924351, -2146047, -2115071, -7168, int32(11602690), int32(257), -256, int32(12289), -12288, int32(0), -2750207, -976383, -2746111, -2763520, -2762752, -2759679, -2751743, -2760447, -2760191, -2768639, int32(0), -1859584, int32(0), int32(257), -256, int32(12323842), int32(0), int32(257), -256, -10830847, int32(237569), -9044991, -10823679, int32(12288), -10833663, -10832639, -10830079, -10818047, -10824191, -10818815, -12287, -10830591, -9058303, int32(0), -9949184, -237568, int32(0), int32(8193), -8192, int32(0), int32(10241), -10240, int32(0), int32(16385), -16384, int32(0), int32(8193), -8192, int32(0), int32(8193), -8192, int32(0), int32(8705), -8704}
var _cgos_rulebases_towctrans [512]uint8 = [512]uint8{uint8(0), uint8(6), uint8(39), uint8(81), uint8(111), uint8(119), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(124), uint8(0), uint8(0), uint8(127), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(131), uint8(142), uint8(146), uint8(151), uint8(0), uint8(170), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(180), uint8(196), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(198), uint8(201), uint8(0), uint8(0), uint8(0), uint8(219), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(222), uint8(0), uint8(0), uint8(0), uint8(0), uint8(225), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(228), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(231), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(234), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(237), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0)}
var _cgos_exceptions_towctrans [200][2]uint8 = [200][2]uint8{[2]uint8{uint8(48), uint8(12)}, [2]uint8{uint8(49), uint8(13)}, [2]uint8{uint8(120), uint8(14)}, [2]uint8{uint8(127), uint8(15)}, [2]uint8{uint8(128), uint8(16)}, [2]uint8{uint8(129), uint8(17)}, [2]uint8{uint8(134), uint8(18)}, [2]uint8{uint8(137), uint8(19)}, [2]uint8{uint8(138), uint8(19)}, [2]uint8{uint8(142), uint8(20)}, [2]uint8{uint8(143), uint8(21)}, [2]uint8{uint8(144), uint8(22)}, [2]uint8{uint8(147), uint8(19)}, [2]uint8{uint8(148), uint8(23)}, [2]uint8{uint8(149), uint8(24)}, [2]uint8{uint8(150), uint8(25)}, [2]uint8{uint8(151), uint8(26)}, [2]uint8{uint8(154), uint8(27)}, [2]uint8{uint8(156), uint8(25)}, [2]uint8{uint8(157), uint8(28)}, [2]uint8{uint8(158), uint8(29)}, [2]uint8{uint8(159), uint8(30)}, [2]uint8{uint8(166), uint8(31)}, [2]uint8{uint8(169), uint8(31)}, [2]uint8{uint8(174), uint8(31)}, [2]uint8{uint8(177), uint8(32)}, [2]uint8{uint8(178), uint8(32)}, [2]uint8{uint8(183), uint8(33)}, [2]uint8{uint8(191), uint8(34)}, [2]uint8{uint8(197), uint8(35)}, [2]uint8{uint8(200), uint8(35)}, [2]uint8{uint8(203), uint8(35)}, [2]uint8{uint8(221), uint8(36)}, [2]uint8{uint8(242), uint8(35)}, [2]uint8{uint8(246), uint8(37)}, [2]uint8{uint8(247), uint8(38)}, [2]uint8{uint8(32), uint8(45)}, [2]uint8{uint8(58), uint8(46)}, [2]uint8{uint8(61), uint8(47)}, [2]uint8{uint8(62), uint8(48)}, [2]uint8{uint8(63), uint8(49)}, [2]uint8{uint8(64), uint8(49)}, [2]uint8{uint8(67), uint8(50)}, [2]uint8{uint8(68), uint8(51)}, [2]uint8{uint8(69), uint8(52)}, [2]uint8{uint8(80), uint8(53)}, [2]uint8{uint8(81), uint8(54)}, [2]uint8{uint8(82), uint8(55)}, [2]uint8{uint8(83), uint8(56)}, [2]uint8{uint8(84), uint8(57)}, [2]uint8{uint8(89), uint8(58)}, [2]uint8{uint8(91), uint8(59)}, [2]uint8{uint8(92), uint8(60)}, [2]uint8{uint8(97), uint8(61)}, [2]uint8{uint8(99), uint8(62)}, [2]uint8{uint8(101), uint8(63)}, [2]uint8{uint8(102), uint8(64)}, [2]uint8{uint8(104), uint8(65)}, [2]uint8{uint8(105), uint8(66)}, [2]uint8{uint8(106), uint8(64)}, [2]uint8{uint8(107), uint8(67)}, [2]uint8{uint8(108), uint8(68)}, [2]uint8{uint8(111), uint8(66)}, [2]uint8{uint8(113), uint8(69)}, [2]uint8{uint8(114), uint8(70)}, [2]uint8{uint8(117), uint8(71)}, [2]uint8{uint8(125), uint8(72)}, [2]uint8{uint8(130), uint8(73)}, [2]uint8{uint8(135), uint8(74)}, [2]uint8{uint8(137), uint8(75)}, [2]uint8{uint8(138), uint8(76)}, [2]uint8{uint8(139), uint8(76)}, [2]uint8{uint8(140), uint8(77)}, [2]uint8{uint8(146), uint8(78)}, [2]uint8{uint8(157), uint8(79)}, [2]uint8{uint8(158), uint8(80)}, [2]uint8{uint8(69), uint8(87)}, [2]uint8{uint8(123), uint8(29)}, [2]uint8{uint8(124), uint8(29)}, [2]uint8{uint8(125), uint8(29)}, [2]uint8{uint8(127), uint8(88)}, [2]uint8{uint8(134), uint8(89)}, [2]uint8{uint8(136), uint8(90)}, [2]uint8{uint8(137), uint8(90)}, [2]uint8{uint8(138), uint8(90)}, [2]uint8{uint8(140), uint8(91)}, [2]uint8{uint8(142), uint8(92)}, [2]uint8{uint8(143), uint8(92)}, [2]uint8{uint8(172), uint8(93)}, [2]uint8{uint8(173), uint8(94)}, [2]uint8{uint8(174), uint8(94)}, [2]uint8{uint8(175), uint8(94)}, [2]uint8{uint8(194), uint8(95)}, [2]uint8{uint8(204), uint8(96)}, [2]uint8{uint8(205), uint8(97)}, [2]uint8{uint8(206), uint8(97)}, [2]uint8{uint8(207), uint8(98)}, [2]uint8{uint8(208), uint8(99)}, [2]uint8{uint8(209), uint8(100)}, [2]uint8{uint8(213), uint8(101)}, [2]uint8{uint8(214), uint8(102)}, [2]uint8{uint8(215), uint8(103)}, [2]uint8{uint8(240), uint8(104)}, [2]uint8{uint8(241), uint8(105)}, [2]uint8{uint8(242), uint8(106)}, [2]uint8{uint8(243), uint8(107)}, [2]uint8{uint8(244), uint8(108)}, [2]uint8{uint8(245), uint8(109)}, [2]uint8{uint8(249), uint8(110)}, [2]uint8{uint8(253), uint8(45)}, [2]uint8{uint8(254), uint8(45)}, [2]uint8{uint8(255), uint8(45)}, [2]uint8{uint8(80), uint8(105)}, [2]uint8{uint8(81), uint8(105)}, [2]uint8{uint8(82), uint8(105)}, [2]uint8{uint8(83), uint8(105)}, [2]uint8{uint8(84), uint8(105)}, [2]uint8{uint8(85), uint8(105)}, [2]uint8{uint8(86), uint8(105)}, [2]uint8{uint8(87), uint8(105)}, [2]uint8{uint8(88), uint8(105)}, [2]uint8{uint8(89), uint8(105)}, [2]uint8{uint8(90), uint8(105)}, [2]uint8{uint8(91), uint8(105)}, [2]uint8{uint8(92), uint8(105)}, [2]uint8{uint8(93), uint8(105)}, [2]uint8{uint8(94), uint8(105)}, [2]uint8{uint8(95), uint8(105)}, [2]uint8{uint8(130), uint8(0)}, [2]uint8{uint8(131), uint8(0)}, [2]uint8{uint8(132), uint8(0)}, [2]uint8{uint8(133), uint8(0)}, [2]uint8{uint8(134), uint8(0)}, [2]uint8{uint8(135), uint8(0)}, [2]uint8{uint8(136), uint8(0)}, [2]uint8{uint8(137), uint8(0)}, [2]uint8{uint8(192), uint8(117)}, [2]uint8{uint8(207), uint8(118)}, [2]uint8{uint8(128), uint8(137)}, [2]uint8{uint8(129), uint8(138)}, [2]uint8{uint8(130), uint8(139)}, [2]uint8{uint8(133), uint8(140)}, [2]uint8{uint8(134), uint8(141)}, [2]uint8{uint8(112), uint8(157)}, [2]uint8{uint8(113), uint8(157)}, [2]uint8{uint8(118), uint8(158)}, [2]uint8{uint8(119), uint8(158)}, [2]uint8{uint8(120), uint8(159)}, [2]uint8{uint8(121), uint8(159)}, [2]uint8{uint8(122), uint8(160)}, [2]uint8{uint8(123), uint8(160)}, [2]uint8{uint8(124), uint8(161)}, [2]uint8{uint8(125), uint8(161)}, [2]uint8{uint8(179), uint8(162)}, [2]uint8{uint8(186), uint8(163)}, [2]uint8{uint8(187), uint8(163)}, [2]uint8{uint8(188), uint8(164)}, [2]uint8{uint8(190), uint8(165)}, [2]uint8{uint8(195), uint8(162)}, [2]uint8{uint8(204), uint8(164)}, [2]uint8{uint8(218), uint8(166)}, [2]uint8{uint8(219), uint8(166)}, [2]uint8{uint8(229), uint8(106)}, [2]uint8{uint8(234), uint8(167)}, [2]uint8{uint8(235), uint8(167)}, [2]uint8{uint8(236), uint8(110)}, [2]uint8{uint8(243), uint8(162)}, [2]uint8{uint8(248), uint8(168)}, [2]uint8{uint8(249), uint8(168)}, [2]uint8{uint8(250), uint8(169)}, [2]uint8{uint8(251), uint8(169)}, [2]uint8{uint8(252), uint8(164)}, [2]uint8{uint8(38), uint8(176)}, [2]uint8{uint8(42), uint8(177)}, [2]uint8{uint8(43), uint8(178)}, [2]uint8{uint8(78), uint8(179)}, [2]uint8{uint8(132), uint8(8)}, [2]uint8{uint8(98), uint8(186)}, [2]uint8{uint8(99), uint8(187)}, [2]uint8{uint8(100), uint8(188)}, [2]uint8{uint8(101), uint8(189)}, [2]uint8{uint8(102), uint8(190)}, [2]uint8{uint8(109), uint8(191)}, [2]uint8{uint8(110), uint8(192)}, [2]uint8{uint8(111), uint8(193)}, [2]uint8{uint8(112), uint8(194)}, [2]uint8{uint8(126), uint8(195)}, [2]uint8{uint8(127), uint8(195)}, [2]uint8{uint8(125), uint8(207)}, [2]uint8{uint8(141), uint8(208)}, [2]uint8{uint8(148), uint8(209)}, [2]uint8{uint8(171), uint8(210)}, [2]uint8{uint8(172), uint8(211)}, [2]uint8{uint8(173), uint8(212)}, [2]uint8{uint8(176), uint8(213)}, [2]uint8{uint8(177), uint8(214)}, [2]uint8{uint8(178), uint8(215)}, [2]uint8{uint8(196), uint8(216)}, [2]uint8{uint8(197), uint8(217)}, [2]uint8{uint8(198), uint8(218)}}
