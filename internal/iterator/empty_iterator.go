package iterator

type EmptyIterator[T any] struct{}

func NewEmptyIterator[T any]() *EmptyIterator[T] {
	return &EmptyIterator[T]{}
}

func (it *EmptyIterator[T]) MoveNext() bool {
	return false
}

func (it *EmptyIterator[T]) Current() T {
	panic("EmptyIterator has no current element")
}

func (it *EmptyIterator[T]) Clone() Iterator[T] {
	return NewEmptyIterator[T]()
}
