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

func (m *QueueInstance) PushIn(request QPushRequest) {
	m.Q.PushBack(request.val)

}
func NewQueue(queue QueueRest, name string) {

}
