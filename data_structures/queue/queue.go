package queue

import "fmt"

type Queue[T comparable] struct {
	first *element[T]
	last  *element[T]
}

type element[T comparable] struct {
	value T
	next  *element[T]
}

func New[T comparable]() Queue[T] {
	return Queue[T]{nil, nil}
}

func (this *Queue[T]) IsEmpty() bool {
	return this.first == nil
}

func (this *Queue[T]) First() T {
	if this.IsEmpty() {
		panic("called First() on enpty queue")
	}
	return this.first.value
}

func newNode[T comparable](e T, p *element[T]) *element[T] {
	return &element[T]{e, p}
}

func (this *Queue[T]) Enqueue(e T) {
	r := newNode(e, nil)
	if this.IsEmpty() {
		this.first = r
		this.last = r
	} else {
		this.last.next = r
		this.last = r
	}
}

func (this *Queue[T]) Dequeue() T {
	if this.IsEmpty() {
		panic("called Dequeue() on enpty queue")
	}
	x := this.first.value
	this.first = this.first.next
	if this.IsEmpty() {
		this.last = nil
	}
	return x
}

func (this Queue[T]) String() string {
	s := "{ first -> "
	if !this.IsEmpty() {
		p := this.first
		s += fmt.Sprint(p.value)
		for p.next != nil {
			p = p.next
			s += " -> " + fmt.Sprint(p.value)
		}
		return s + " <- last }"
	} else {
		return "{ }"
	}
}
