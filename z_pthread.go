package libc

import (
	"sync"

	"github.com/petermattis/goid"
)

var (
	c2go_gtls sync.Map
)

func __pthread_self() (self Pthread_t) {
	id := goid.Get()
	if ret, ok := c2go_gtls.Load(id); ok {
		self = ret.(*Struct___pthread)
	} else {
		self = &Struct___pthread{Tid: int32(id), Locale: &__libc.global_locale}
		c2go_gtls.Store(id, self)
	}
	return
}
