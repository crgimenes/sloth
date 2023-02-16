package queue

type Queue interface {
	Publish(any) error
	Consume() (any, error)
}
