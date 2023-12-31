package datastructures

import "github.com/johnjamespj/project_datastorm/internal/iterator"

type DoublyLinkedListNode[V any] struct {
	Value V
	Next  *DoublyLinkedListNode[V]
	Prev  *DoublyLinkedListNode[V]
}

type DoublyLinkedList[V any] struct {
	Head *DoublyLinkedListNode[V]
	Tail *DoublyLinkedListNode[V]
}

func NewDoublyLinkedList[V any]() *DoublyLinkedList[V] {
	return &DoublyLinkedList[V]{}
}

func (l *DoublyLinkedList[V]) AddBeforeNode(node *DoublyLinkedListNode[V], value V) {
	newNode := &DoublyLinkedListNode[V]{Value: value}
	if node.Prev == nil {
		l.Head = newNode
	} else {
		node.Prev.Next = newNode
	}

	newNode.Prev = node.Prev
	newNode.Next = node
	node.Prev = newNode
}

func (l *DoublyLinkedList[V]) AddAfterNode(node *DoublyLinkedListNode[V], value V) {
	newNode := &DoublyLinkedListNode[V]{Value: value}
	if node.Next == nil {
		l.Tail = newNode
	} else {
		node.Next.Prev = newNode
	}

	newNode.Next = node.Next
	newNode.Prev = node
	node.Next = newNode
}

func (l *DoublyLinkedList[V]) Add(value V) {
	newNode := &DoublyLinkedListNode[V]{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		l.Tail.Next = newNode
	}

	newNode.Prev = l.Tail
	l.Tail = newNode
}

func (l *DoublyLinkedList[V]) Remove(node *DoublyLinkedListNode[V]) {
	if node.Prev == nil {
		l.Head = node.Next
	} else {
		node.Prev.Next = node.Next
	}

	if node.Next == nil {
		l.Tail = node.Prev
	} else {
		node.Next.Prev = node.Prev
	}
}

func (l *DoublyLinkedList[V]) ItrFromHead() *DoublyLinkedListIterator[V] {
	return &DoublyLinkedListIterator[V]{CurrentNode: l.Head, direction: true}
}

func (l *DoublyLinkedList[V]) ItrFromTail() *DoublyLinkedListIterator[V] {
	return &DoublyLinkedListIterator[V]{CurrentNode: l.Tail, direction: false}
}

type DoublyLinkedListIterator[V any] struct {
	CurrentNode *DoublyLinkedListNode[V]
	direction   bool
}

func (l *DoublyLinkedList[V]) Iterator() *DoublyLinkedListIterator[V] {
	return &DoublyLinkedListIterator[V]{CurrentNode: l.Head}
}

func (i *DoublyLinkedListIterator[V]) MoveNext() bool {
	if i.CurrentNode == nil {
		return false
	}

	if i.direction {
		i.CurrentNode = i.CurrentNode.Next
	} else {
		i.CurrentNode = i.CurrentNode.Prev
	}

	return i.CurrentNode != nil
}

func (i *DoublyLinkedListIterator[V]) Current() V {
	return i.CurrentNode.Value
}

func (i *DoublyLinkedListIterator[V]) Clone() iterator.Iterator[V] {
	return &DoublyLinkedListIterator[V]{CurrentNode: i.CurrentNode, direction: i.direction}
}

func (i *DoublyLinkedListIterator[V]) Reverse() {
	i.direction = !i.direction
}
