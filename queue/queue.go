package queue

type Queue interface {
	Produce(any) error
	Consume() (any, error)
}
