package iterator

type SliceIterator[T any] struct {
	slice     []T
	index     int
	direction bool
}

func NewSliceIterator[T any](slice []T) Iterator[T] {
	return &SliceIterator[T]{
		slice:     slice,
		index:     -1,
		direction: true,
	}
}

func (it *SliceIterator[T]) MoveNext() bool {
	if it.direction {
		it.index++
		return it.index < len(it.slice)
	}
	it.index--
	return it.index >= 0
}

func (it *SliceIterator[T]) Current() T {
	if it.index < 0 || it.index >= len(it.slice) {
		panic("SliceIterator: iterator out of bounds, this could means the iterator is not initialized or has been exhausted")
	}

	return it.slice[it.index]
}

func (it *SliceIterator[T]) Reset() {
	if it.direction {
		it.index = -1
	} else {
		it.index = len(it.slice)
	}
}

func (it *SliceIterator[T]) Clone() Iterator[T] {
	return &SliceIterator[T]{
		slice:     it.slice,
		index:     it.index,
		direction: it.direction,
	}
}

func (it *SliceIterator[T]) Reverse() {
	it.direction = !it.direction
}

func (it *SliceIterator[T]) IsReversable() bool {
	return true
}
