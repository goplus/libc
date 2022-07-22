package libc

func __ofl_add(f *struct__IO_FILE) *struct__IO_FILE {
	var head **struct__IO_FILE = __ofl_lock()
	f.next = *head
	if *head != nil {
		(*head).prev = f
	}
	*head = f
	__ofl_unlock()
	return f
}
