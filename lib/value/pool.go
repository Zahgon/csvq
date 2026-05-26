package value

import (
	"sync"
)

var stringPool = &sync.Pool{
	New: func() interface{} {
		return &String{}
	},
}

var integerPool = &sync.Pool{
	New: func() interface{} {
		return &Integer{}
	},
}

var floatPool = &sync.Pool{
	New: func() interface{} {
		return &Float{}
	},
}

var datetimePool = &sync.Pool{
	New: func() interface{} {
		return &Datetime{}
	},
}

func getString() *String { _ = "STUB: not implemented"; return nil }

func getInteger() *Integer { _ = "STUB: not implemented"; return nil }

func getFloat() *Float { _ = "STUB: not implemented"; return nil }

func getDatetime() *Datetime { _ = "STUB: not implemented"; return nil }

func Discard(p Primary) { _ = "STUB: not implemented"; return }
