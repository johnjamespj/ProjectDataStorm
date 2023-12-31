package iterator

type MapIterator[T1 any, T2 any] struct {
	Iterator[T1]
	mapper func(T1) T2
}

func Map[T1 any, T2 any](iterator Iterator[T1], mapper func(T1) T2) Iterator[T2] {
	return &MapIterator[T1, T2]{
		Iterator: iterator,
		mapper:   mapper,
	}
}

func (it *MapIterator[T1, T2]) Current() T2 {
	return it.mapper(it.Iterator.Current())
}

func (it *MapIterator[T1, T2]) Clone() Iterator[T2] {
	return &MapIterator[T1, T2]{
		Iterator: it.Iterator.Clone(),
		mapper:   it.mapper,
	}
}
