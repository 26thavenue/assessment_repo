package main


type SinglyLinkedListNode[T any] struct {
	data T
	next *SinglyLinkedListNode[T]
}

func NewNode[T any](data T) *SinglyLinkedListNode[T] {
	return &SinglyLinkedListNode[T]{data: data, next: nil}
}

func (n *SinglyLinkedListNode[T]) GetData() T {
	return n.data
}

func (n *SinglyLinkedListNode[T]) GetNext() *SinglyLinkedListNode[T] {
	return n.next
}

func (n *SinglyLinkedListNode[T]) SetNext(next *SinglyLinkedListNode[T]) {
	n.next = next
}

type SinglyLinkedList[T any] struct {
	head *SinglyLinkedListNode[T]
	tail *SinglyLinkedListNode[T]
	size int
}

func NewList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (l *SinglyLinkedList[T]) AddFront(data T) {
	newNode := NewNode(data)

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.SetNext(l.head)
		l.head = newNode
	}

	l.size++
}

func (l *SinglyLinkedList[T]) AddBack(data T) {
	newNode := NewNode(data)

	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.SetNext(newNode)
		l.tail = newNode
	}

	l.size++
}

func (l *SinglyLinkedList[T]) GetHead() *SinglyLinkedListNode[T] {
	return l.head
}

func (l *SinglyLinkedList[T]) GetTail() *SinglyLinkedListNode[T] {
	return l.tail
}

func (l *SinglyLinkedList[T]) GetSize() int {
	return l.size
}