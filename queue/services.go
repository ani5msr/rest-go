package queue

import (
	"sync"
)

type QueueInstance struct {
	Q    []string
	Lock sync.RWMutex
	Name string
}

var qumap = make(map[string][]string)

func NewQueue(queue QueueRest, name string) {

}
