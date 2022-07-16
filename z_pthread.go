package libc

import (
	"sync"

	"github.com/petermattis/goid"
)

var (
	c2go_gtls sync.Map
)

func __pthread_self() (self *struct___pthread) {
	id := goid.Get()
	if ret, ok := c2go_gtls.Load(id); ok {
		self = ret.(*struct___pthread)
	} else {
		self = &struct___pthread{tid: int32(id)}
		c2go_gtls.Store(id, self)
	}
	return
}
