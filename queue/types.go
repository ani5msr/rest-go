package queue

type QPushRequest struct {
	name string
	val  string
}
type QPopRequest struct {
	name string
}
type BQPopRequest struct {
	name string
	time int
}
type Deque[T any] struct {
	buf    []T
	head   int
	tail   int
	count  int
	minCap int
	name   string
}
