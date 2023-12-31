package datastructures

import (
	"sort"

	"github.com/johnjamespj/project_datastorm/internal/interfaces"
	"github.com/johnjamespj/project_datastorm/internal/iterator"
)

type SortedList[V any] struct {
	list       []V
	comparator interfaces.CompareFunc[V, V]
}

func NewSortedList[V any](list []V, comparator interfaces.CompareFunc[V, V]) *SortedList[V] {
	sortedList := &SortedList[V]{
		list:       list,
		comparator: comparator,
	}
	return sortedList
}

// Add adds a value to the list efficiently
func (v *SortedList[V]) Add(value V) {
	v.list = append(v.list, value)
	sort.Slice(v.list, func(i, j int) bool {
		return v.comparator(v.list[i], v.list[j]) < 0
	})
}

// Add all values to the list efficiently
func (v *SortedList[V]) AddAll(values []V) {
	v.list = append(v.list, values...)
	sort.Slice(v.list, func(i, j int) bool {
		return v.comparator(v.list[i], v.list[j]) < 0
	})
}

func (v *SortedList[V]) Merge(other []V) {
	// TODO: merge two sorted lists
	newList := make([]V, 0)
	i, j := 0, 0
	for i < len(v.list) && j < len(other) {
		if v.comparator(v.list[i], other[j]) < 0 {
			newList = append(newList, v.list[i])
			i++
		} else {
			newList = append(newList, other[j])
			j++
		}
	}

	for i < len(v.list) {
		newList = append(newList, v.list[i])
		i++
	}

	for j < len(other) {
		newList = append(newList, other[j])
		j++
	}

	v.list = newList
}

// Remove removes a value from the list efficiently
func (v *SortedList[V]) Remove(value V) *V {
	i := sort.Search(len(v.list), func(i int) bool {
		return v.comparator(v.list[i], value) >= 0
	})

	if i >= len(v.list) {
		return nil
	}

	if v.comparator(v.list[i], value) == 0 {
		v.list = append(v.list[:i], v.list[i+1:]...)
	}

	return &value
}

// Remove removes a value from the list efficiently
func (v *SortedList[V]) RemoveWhere(predicate func(V) bool) []V {
	var removed []V

	for i := 0; i < len(v.list); i++ {
		if predicate(v.list[i]) {
			removed = append(removed, v.list[i])
			v.list = append(v.list[:i], v.list[i+1:]...)
			i--
		}
	}

	return removed
}

func (v *SortedList[V]) Clear() {
	v.list = make([]V, 0)
}

func (v *SortedList[V]) ToList() []V {
	ret := make([]V, len(v.list))
	copy(ret, v.list)
	return ret
}

// Returns the first value. O(1)
func (v *SortedList[V]) First() *V {
	if len(v.list) == 0 {
		return nil
	}

	return &v.list[0]
}

// Returns the last value. O(1)
func (v *SortedList[V]) Last() *V {
	if len(v.list) == 0 {
		return nil
	}

	return &v.list[len(v.list)-1]
}

// Returns a key-value mapping associated with the least key greater
// than or equal to the given key, or null if there is no such key. O(log n)
func (v *SortedList[V]) Floor(k V) *V {
	i, _ := sort.Find(len(v.list), func(i int) int {
		return v.comparator(k, v.list[i])
	})

	if i >= len(v.list) {
		return &v.list[len(v.list)-1]
	}

	for i >= 0 && v.comparator(v.list[i], k) > 0 {
		i--
	}

	if i < 0 {
		return nil
	}

	return &v.list[i]
}

// Returns a key-value mapping associated with the greatest key less
// than or equal to the given key, or null if there is no such key.
func (v *SortedList[V]) Lower(k V) *V {
	i, found := sort.Find(len(v.list), func(i int) int {
		return v.comparator(k, v.list[i])
	})

	if found && i > 0 {
		return &v.list[i-1]
	} else if found {
		return nil
	} else if i >= len(v.list) {
		return &v.list[len(v.list)-1]
	}

	return &v.list[i]
}

// Returns a key-value mapping associated with the least key greater
// than or equal to the given key, or null if there is no such key.
func (v *SortedList[V]) Ceiling(k V) *V {
	i, _ := sort.Find(len(v.list), func(i int) int {
		return v.comparator(k, v.list[i])
	})

	for i+1 < len(v.list) {
		if v.comparator(v.list[i+1], k) > 0 {
			break
		}
		i++
	}

	if i >= len(v.list) || v.comparator(v.list[i], k) < 0 {
		return nil
	}

	return &v.list[i]
}

// Returns a key-value mapping associated with the least key
// strictly greater than the given key, or null if there is
// no such key.
func (v *SortedList[V]) Higher(k V) *V {
	i := sort.Search(len(v.list), func(i int) bool {
		return v.comparator(v.list[i], k) > 0
	})

	if i >= len(v.list) || v.comparator(v.list[i], k) == 0 {
		return nil
	}

	return &v.list[i]
}

// Returns a view of the portion of this map whose keys are
// strictly less than toKey.
func (v *SortedList[V]) Tail(fromKey V, inclusive bool) iterator.Iterator[V] {
	i := sort.Search(len(v.list), func(i int) bool {
		return v.comparator(v.list[i], fromKey) >= 0
	})

	for i < len(v.list)-1 && !inclusive {
		if v.comparator(v.list[i+1], fromKey) > 0 {
			i++
			break
		}
		i++
	}

	return &SortedListIterator[V]{
		list:  v.list[i:],
		index: -1,
	}
}

// Returns a view of the portion of this map whose keys are
// greater than or equal to toKey.
func (v *SortedList[V]) Head(toKey V, inclusive bool) iterator.Iterator[V] {
	i, found := sort.Find(len(v.list), func(i int) int {
		return v.comparator(toKey, v.list[i])
	})

	if !inclusive {
		if found && i > 0 {
			i--
		} else if found {
			i = -1
		}
	} else {
		for i < len(v.list)-1 && v.comparator(v.list[i+1], toKey) == 0 {
			i++
		}
	}

	return &SortedListIterator[V]{
		list:  v.list[:i+1],
		index: -1,
	}
}

// Returns a view of the portion of this map whose keys range
// from fromKey, inclusive, to toKey, exclusive.
func (v *SortedList[V]) Sub(fromKey V, toKey V, fromInclusive bool, toInclusive bool) iterator.Iterator[V] {
	// find the first element that is greater than fromKey
	i, _ := sort.Find(len(v.list), func(i int) int {
		return v.comparator(fromKey, v.list[i])
	})

	for i < len(v.list) && !fromInclusive {
		if v.comparator(v.list[i+1], fromKey) > 0 {
			i++
			break
		}
		i++
	}

	// find the first element that is greater than toKey
	j, _ := sort.Find(len(v.list), func(i int) int {
		return v.comparator(toKey, v.list[i])
	})

	for j < len(v.list) && !toInclusive {
		if v.comparator(v.list[j+1], toKey) > 0 {
			j++
			break
		}
		j++
	}

	var itr iterator.Iterator[V]
	if i > j || i >= len(v.list) || j < 0 {
		emptyArray := make([]V, 0)
		itr = &SortedListIterator[V]{
			list:  emptyArray,
			index: -1,
		}
	} else {
		itr = &SortedListIterator[V]{
			list:  v.list[i : j+1],
			index: -1,
		}
	}

	return itr
}

// Returns all the entry matching the value. O(log n)
func (v *SortedList[V]) Get(value V) iterator.Iterator[V] {
	return v.Sub(value, value, true, true)
}

func (v *SortedList[V]) GetSize() int {
	return len(v.list)
}

// returns sortedlist iterator
func (v *SortedList[V]) GetIterator() iterator.Iterator[V] {
	return &SortedListIterator[V]{
		list:  v.list,
		index: -1,
	}
}

// Iterator for a sortedlist
type SortedListIterator[V any] struct {
	index int
	list  []V
}

// move next for SortedListIterator
func (i *SortedListIterator[V]) MoveNext() bool {
	if i.index < len(i.list)-1 {
		i.index++
		return true
	}
	return false
}

// get current for SortedListIterator
func (i *SortedListIterator[V]) Current() V {
	if i.index >= len(i.list) || i.index < 0 {
		panic("Iterator: No more items left or the first MoveNext() is called")
	}

	return i.list[i.index]
}

// clone the SortedListIterator
func (i *SortedListIterator[V]) Clone() iterator.Iterator[V] {
	return &SortedListIterator[V]{
		index: i.index,
		list:  i.list,
	}
}
