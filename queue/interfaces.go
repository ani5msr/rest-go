package queue

type QueueRest interface {
	PushIn(request QPushRequest) (response string, err error)
	Pop(request QPopRequest)
	BqPOP(request BQPopRequest) (response string, err error)
}
