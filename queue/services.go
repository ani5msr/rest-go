package queue

import (
	"github.com/gammazero/deque"
)

type QueueInstance struct {
	queue QueueRest
	q     deque.Deque[string]
	name  string
}

var queuemap map[string]QueueInstance

func (m *QueueInstance) PushIn(request QPushRequest) {

}
func NewQueue(queue QueueRest, name string) {

}
