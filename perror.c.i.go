package libc

import unsafe "unsafe"

func Perror(msg *int8) {
	var f *Struct__IO_FILE = &__stderr_FILE
	var errstr *int8 = Strerror(*__errno_location())
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	var old_locale unsafe.Pointer = unsafe.Pointer(f.Locale)
	var old_mode int32 = f.Mode
	if msg != nil && int32(*msg) != 0 {
		Fwrite(unsafe.Pointer(msg), Strlen(msg), uint64(1), f)
		Fputc(':', f)
		Fputc(' ', f)
	}
	Fwrite(unsafe.Pointer(errstr), Strlen(errstr), uint64(1), f)
	Fputc('\n', f)
	f.Mode = old_mode
	f.Locale = (*Struct___locale_struct)(old_locale)
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
}
