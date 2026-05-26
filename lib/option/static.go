package option

import (
	"math/rand"
	"sync"
	"time"
)

var (
	TestTime  time.Time // For Tests
	Timezones = NewTimezoneMap()

	random  *rand.Rand
	getRand sync.Once
)

func GetRand() *rand.Rand { _ = "STUB: not implemented"; return nil }

func GetLocation(timezone string) (*time.Location, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Now(location *time.Location) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type TimezoneMap struct {
	m   *sync.Map
	mtx *sync.Mutex
}

func NewTimezoneMap() TimezoneMap { _ = "STUB: not implemented"; return *new(TimezoneMap) }

func (tzmap TimezoneMap) store(key string, value *time.Location) { _ = "STUB: not implemented"; return }

func (tzmap TimezoneMap) load(key string) (*time.Location, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (tzmap TimezoneMap) Get(timezone string) (*time.Location, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
