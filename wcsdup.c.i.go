package libc

func wcsdup(s *uint32) *uint32 {
	var l uint64 = wcslen(s)
	var d *uint32 = (*uint32)(Malloc((l + uint64(1)) * 4))
	if !(d != nil) {
		return (*uint32)(nil)
	}
	return wmemcpy(d, s, l+uint64(1))
}
