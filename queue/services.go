package queue

import (
	"sync"

	"github.com/gammazero/deque"
)

type QueueInstance struct {
	Q    deque.Deque[string]
	Lock sync.RWMutex
	Name string
}

var qumap = make(map[string][]string)

func NewQueue(queue QueueRest, name string) {

}
