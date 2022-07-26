package libc

func __ofl_add(f *Struct__IO_FILE) *Struct__IO_FILE {
	var head **Struct__IO_FILE = __ofl_lock()
	f.Next = *head
	if *head != nil {
		(*head).Prev = f
	}
	*head = f
	__ofl_unlock()
	return f
}
