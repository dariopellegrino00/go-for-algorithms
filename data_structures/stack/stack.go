package stack

import "fmt"

type Stack[T comparable] struct {
	top *element[T]
}

type element[T comparable] struct {
	value T
	next  *element[T]
}

func New[T comparable]() Stack[T] {
	return Stack[T]{nil}
}

func newNode[T comparable](e T, p *element[T]) *element[T] {
	return &element[T]{e, p}
}

func (this *Stack[T]) IsEmpty() bool {
	return this.top == nil
}

func (this *Stack[T]) Push(e T) {
	r := newNode(e, this.top)
	this.top = r
}

func (this *Stack[T]) Top() T {
	if this.IsEmpty() {
		panic("called Top() on empty stack")
	}
	return this.top.value
}

func (this *Stack[T]) Pop() T {
	if this.IsEmpty() {
		panic("called Pop() on empty stack")
	}
	e := this.top.value
	this.top = this.top.next
	return e
}

func (this Stack[T]) String() string {
	s := "{ top -> "
	if !this.IsEmpty() {
		p := this.top
		s += fmt.Sprint(p.value)
		for p.next != nil {
			p = p.next
			s += " -> " + fmt.Sprint(p.value)
		}
		return s + " }"
	} else {
		return "{ }"
	}
}
