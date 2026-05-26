package file

import (
	"context"
	"math/rand"
	"sync"
	"time"
)

const rlockFileSuffixLen = 12

var dummyCancelFunc = func() {}

var (
	letterRunes    = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randForLock    *rand.Rand
	getRandForLock sync.Once
)

func randStrForLock() *rand.Rand { _ = "STUB: not implemented"; return nil }

func RandomString(length int) string { _ = "STUB: not implemented"; return "" }

func rlockFileSuffix() string { _ = "STUB: not implemented"; return "" }

func GetTimeoutContext(ctx context.Context, waitTimeOut time.Duration) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func RLockFilePath(path string) string { _ = "STUB: not implemented"; return "" }

func LockFilePath(path string) string { _ = "STUB: not implemented"; return "" }

func TempFilePath(path string) string { _ = "STUB: not implemented"; return "" }

func getFilePath(path string, suffix string) string { _ = "STUB: not implemented"; return "" }

func RLockExists(path string) bool { _ = "STUB: not implemented"; return false }

func LockExists(path string) bool { _ = "STUB: not implemented"; return false }

func Exists(path string) bool { _ = "STUB: not implemented"; return false }
