package queue

import (
	"github.com/gammazero/deque"
)

type QueueRest interface {
	PushIn(string name, string val) (string response, error)
	Pop(string name) (string response, error)
	BqPOP(string name, string time) (string response, error)
}
