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
