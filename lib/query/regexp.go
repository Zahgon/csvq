package query

import (
	"regexp"
)

var validFlagsRegExp = regexp.MustCompile("[imsU]+")

var RegExps = NewRegExpMap()

type RegExpMap struct {
	*SyncMap
}

func NewRegExpMap() RegExpMap { _ = "STUB: not implemented"; return *new(RegExpMap) }

func (rem RegExpMap) Store(key string, value *regexp.Regexp) { _ = "STUB: not implemented"; return }

func (rem RegExpMap) Load(key string) (*regexp.Regexp, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (rem RegExpMap) Get(expr string) (*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
