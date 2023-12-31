package iterator

type WhileIterator[T any] struct {
	Iterator[T]
	predicate func(T) bool
}

func While[T any](iterator Iterator[T], predicate func(T) bool) Iterator[T] {
	return &WhileIterator[T]{
		Iterator:  iterator,
		predicate: predicate,
	}
}

func (it *WhileIterator[T]) MoveNext() bool {
	if !it.Iterator.MoveNext() {
		return false
	}

	if !it.predicate(it.Iterator.Current()) {
		return false
	}

	return true
}
