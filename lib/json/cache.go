package json

import (
	"sync"
)

var Path = NewPathMap()
var Query = NewQueryMap()

type PathMap struct {
	m   *sync.Map
	mtx *sync.Mutex
}

func NewPathMap() PathMap { _ = "STUB: not implemented"; return *new(PathMap) }

func (pmap PathMap) store(key string, value PathExpression) { _ = "STUB: not implemented"; return }

func (pmap PathMap) load(key string) (PathExpression, bool) {
	_ = "STUB: not implemented"
	return *new(PathExpression), false
}

func (pmap PathMap) Parse(s string) (PathExpression, error) {
	_ = "STUB: not implemented"
	return *new(PathExpression), nil
}

type QueryMap struct {
	m   *sync.Map
	mtx *sync.Mutex
}

func NewQueryMap() QueryMap { _ = "STUB: not implemented"; return *new(QueryMap) }

func (qmap QueryMap) store(key string, value QueryExpression) { _ = "STUB: not implemented"; return }

func (qmap QueryMap) load(key string) (QueryExpression, bool) {
	_ = "STUB: not implemented"
	return *new(QueryExpression), false
}

func (qmap QueryMap) Parse(s string) (QueryExpression, error) {
	_ = "STUB: not implemented"
	return *new(QueryExpression), nil
}
