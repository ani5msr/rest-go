package redis

import "sync"

type SetInterface struct {
	Key    string
	Value  string
	Expiry int64
	Lock   sync.Mutex
}
