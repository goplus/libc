package libc

import unsafe "unsafe"

type struct___dirstream struct {
}

func Execl(path *int8, argv0 *int8, __cgo_args ...interface {
}) int32 {
	panic("notimpl")
}
func Fdopen(int32, *int8) *Struct__IO_FILE {
	panic("notimpl")
}
func Fork() int32 {
	panic("notimpl")
}
func Ftrylockfile(*Struct__IO_FILE) int32 {
	panic("notimpl")
}
func Getcwd(buf *int8, size uint64) *int8 {
	panic("notimpl")
}
func Ioctl(fd int32, req int32, __cgo_args ...interface {
}) int32 {
	panic("notimpl")
}
func Mblen(*int8, uint64) int32 {
	panic("notimpl")
}
func Mbtowc(*uint32, *int8, uint64) int32 {
	panic("notimpl")
}
func Pipe2(fd *int32, flag int32) int32 {
	panic("notimpl")
}
func Vfscanf(*Struct__IO_FILE, *int8, []interface {
}) int32 {
	panic("notimpl")
}
func Wctomb(*int8, uint32) int32 {
	panic("notimpl")
}
func X_Exit(int32) {
	panic("notimpl")
}
func __aio_close(int32) int32 {
	panic("notimpl")
}
func __block_all_sigs(unsafe.Pointer) {
	panic("notimpl")
}
func __futexwait(addr unsafe.Pointer, val int32, priv int32) {
	panic("notimpl")
}
func __procfdname(*int8, uint32) {
	panic("notimpl")
}
func __randname(*int8) *int8 {
	panic("notimpl")
}
func __register_locked_file(*Struct__IO_FILE, *Struct___pthread) {
	panic("notimpl")
}
func __restore_sigs(unsafe.Pointer) {
	panic("notimpl")
}
func __synccall(func(unsafe.Pointer), unsafe.Pointer) {
	panic("notimpl")
}
func __unlist_locked_file(*Struct__IO_FILE) {
	panic("notimpl")
}
func __vm_wait() {
	panic("notimpl")
}
func __wake(addr unsafe.Pointer, cnt int32, priv int32) {
	panic("notimpl")
}
func closedir(*struct___dirstream) int32 {
	panic("notimpl")
}
func fcntl(int32, int32, ...interface {
}) int32 {
	panic("notimpl")
}
func fdopendir(int32) *struct___dirstream {
	panic("notimpl")
}
func getgrouplist(*int8, uint32, *uint32, *int32) int32 {
	panic("notimpl")
}
func getrandom(unsafe.Pointer, uint64, uint32) int64 {
	panic("notimpl")
}
func kill(int32, int32) int32 {
	panic("notimpl")
}
func lstat(*int8, *struct_stat) int32 {
	panic("notimpl")
}
func mbrtowc(*uint32, *int8, uint64, *Struct___mbstate_t) uint64 {
	panic("notimpl")
}
func open(*int8, int32, ...interface {
}) int32 {
	panic("notimpl")
}
func pthread_setcancelstate(int32, *int32) int32 {
	panic("notimpl")
}
func pthread_sigmask(int32, *Struct___sigset_t, *Struct___sigset_t) int32 {
	panic("notimpl")
}
func readdir(*struct___dirstream) *struct_dirent {
	panic("notimpl")
}
func setgroups(uint64, *uint32) int32 {
	panic("notimpl")
}
func setitimer(int32, *struct_itimerval, *struct_itimerval) int32 {
	panic("notimpl")
}
func sigfillset(*Struct___sigset_t) int32 {
	panic("notimpl")
}
func stat(*int8, *struct_stat) int32 {
	panic("notimpl")
}
func tcsetattr(int32, int32, *struct_termios) int32 {
	panic("notimpl")
}
func uname(*struct_utsname) int32 {
	panic("notimpl")
}
func vfwprintf(*Struct__IO_FILE, *uint32, []interface {
}) int32 {
	panic("notimpl")
}
func vfwscanf(*Struct__IO_FILE, *uint32, []interface {
}) int32 {
	panic("notimpl")
}
func waitpid(int32, *int32, int32) int32 {
	panic("notimpl")
}
func wcrtomb(*int8, uint32, *Struct___mbstate_t) uint64 {
	panic("notimpl")
}
func wcsrtombs(*int8, **uint32, uint64, *Struct___mbstate_t) uint64 {
	panic("notimpl")
}
