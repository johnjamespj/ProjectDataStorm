package iterator

import "github.com/johnjamespj/project_datastorm/internal/interfaces"

type Iterator[T any] interface {
	MoveNext() bool

	Current() T

	Clone() Iterator[T]
}

type BidirectionalIterator[T any] interface {
	Iterator[T]

	Reverse() BidirectionalIterator[T]
}

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

func ToSlice[T any](it Iterator[T]) []T {
	var slice []T
	for it.MoveNext() {
		slice = append(slice, it.Current())
	}
	return slice
}

func Contains[T1 any, T2 any](it Iterator[T1], compare interfaces.CompareFunc[T1, T2], other T2) bool {
	for it.MoveNext() {
		if compare(it.Current(), other) == 0 {
			return true
		}
	}
	return false
}

func FirstWhere[T any](it Iterator[T], predicate func(T) bool) (T, bool) {
	for it.MoveNext() {
		if predicate(it.Current()) {
			return it.Current(), true
		}
	}
	return it.Current(), false
}

func LastWhere[T any](it Iterator[T], predicate func(T) bool) (T, bool) {
	var last T
	found := false
	for it.MoveNext() {
		if predicate(it.Current()) {
			last = it.Current()
			found = true
		}
	}
	return last, found
}

func Every[T any](it Iterator[T], predicate func(T) bool) bool {
	for it.MoveNext() {
		if !predicate(it.Current()) {
			return false
		}
	}
	return true
}

func Any[T any](it Iterator[T], predicate func(T) bool) bool {
	for it.MoveNext() {
		if predicate(it.Current()) {
			return true
		}
	}
	return false
}

func SingleWhere[T any](it Iterator[T], predicate func(T) bool) bool {
	found := false
	for it.MoveNext() {
		if predicate(it.Current()) {
			if found {
				return false
			}
			found = true
		}
	}
	return found
}

func ElementAt[T any](it Iterator[T], index int) (T, bool) {
	for it.MoveNext() {
		if index == 0 {
			return it.Current(), true
		}
		index--
	}
	return it.Current(), false
}
