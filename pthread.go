package libc

import (
	"sync"

	"github.com/petermattis/goid"
)

var (
	gtls sync.Map
)

func __pthread_self() *struct___pthread {
	var self *struct___pthread
	id := goid.Get()
	ret, ok := gtls.Load(id)
	if !ok {
		self = &struct___pthread{tid: int32(id)}
		gtls.Store(id, self)
	} else {
		self = ret.(*struct___pthread)
	}
	return self
}
