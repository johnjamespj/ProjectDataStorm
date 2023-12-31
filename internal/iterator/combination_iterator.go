package iterator

import "github.com/johnjamespj/project_datastorm/internal/interfaces"

type FollowingIterator[T any] struct {
	Iterator[T]
	next        Iterator[T]
	nextStarted bool
}

func FollowedBy[T any](iterator Iterator[T], next Iterator[T]) Iterator[T] {
	return &FollowingIterator[T]{
		Iterator: iterator,
		next:     next,
	}
}

func (it *FollowingIterator[T]) MoveNext() bool {
	if !it.nextStarted && it.Iterator.MoveNext() {
		return true
	}

	it.nextStarted = true
	return it.next.MoveNext()
}

func (it *FollowingIterator[T]) Current() T {
	if !it.nextStarted {
		return it.Iterator.Current()
	}
	return it.next.Current()
}

func FollowedByAll[T any](iterator Iterator[T], next ...Iterator[T]) Iterator[T] {
	for _, n := range next {
		iterator = FollowedBy(iterator, n)
	}
	return iterator
}

type MergeSortedIterator[T any] struct {
	itrA, itrB       Iterator[T]
	compare          interfaces.CompareFunc[T, T]
	a_empty, b_empty bool
	moved            bool
}

func MergeSort[T any](itrA, itrB Iterator[T], compare func(T, T) int) Iterator[T] {
	return &MergeSortedIterator[T]{
		itrA:    itrA,
		itrB:    itrB,
		a_empty: false,
		b_empty: false,
		moved:   false,
		compare: compare,
	}
}

func (it *MergeSortedIterator[T]) MoveNext() bool {
	if !it.moved {
		it.moved = true
		it.a_empty = !it.itrA.MoveNext()
		it.b_empty = !it.itrB.MoveNext()
		return !it.a_empty || !it.b_empty
	}

	if it.a_empty && it.b_empty {
		return false
	}

	if it.a_empty {
		it.b_empty = !it.itrB.MoveNext()
		return !it.b_empty
	}

	if it.b_empty {
		it.a_empty = !it.itrA.MoveNext()
		return !it.a_empty
	}

	if it.compare(it.itrA.Current(), it.itrB.Current()) < 0 {
		it.a_empty = !it.itrA.MoveNext()
		return true
	} else {
		it.b_empty = !it.itrB.MoveNext()
		return true
	}
}

func (it *MergeSortedIterator[T]) Current() T {
	if it.a_empty {
		return it.itrB.Current()
	}
	if it.b_empty {
		return it.itrA.Current()
	}

	if it.compare(it.itrA.Current(), it.itrB.Current()) < 0 {
		return it.itrA.Current()
	} else {
		return it.itrB.Current()
	}
}

func (it *MergeSortedIterator[T]) Clone() Iterator[T] {
	return &MergeSortedIterator[T]{
		itrA:    it.itrA.Clone(),
		itrB:    it.itrB.Clone(),
		compare: it.compare,
	}
}

func MergeSortAll[T any](compare func(T, T) int, iterators ...Iterator[T]) Iterator[T] {
	if len(iterators) == 0 {
		return NewEmptyIterator[T]()
	}

	if len(iterators) == 1 {
		return iterators[0]
	}

	mid := len(iterators) / 2
	return MergeSort(
		MergeSortAll[T](compare, iterators[:mid]...),
		MergeSortAll[T](compare, iterators[mid:]...),
		compare,
	)
}
