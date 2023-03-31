package queue

type QueueRest interface {
	PushIn(name string, val string) (response string, err error)
	Pop(name string)
	BqPOP(name string, time string) (response string, err error)
}
