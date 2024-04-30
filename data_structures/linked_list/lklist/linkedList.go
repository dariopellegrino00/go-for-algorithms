package lklist

import (
	"errors"
	"fmt"
)

type listNode[T comparable] struct {
	item T
	next *listNode[T]
}

type Linkedlist[T comparable] struct {
	head *listNode[T]
}

func NewList[T comparable]() Linkedlist[T] {
	return Linkedlist[T]{nil}
}

//returns a node of type T for a linked list (private)
func newNode[T comparable](e T) *listNode[T] {
	return &listNode[T]{e, nil}
}

type ListIterator[T comparable] struct {
	current *listNode[T]
}

func (this *Linkedlist[T]) Iterator() ListIterator[T] {
	if this.IsEmpty() {
		return ListIterator[T]{nil}
	} else {
		return ListIterator[T]{this.head}
	}

}

func (it *ListIterator[T]) HasNext() bool {
	return it.current != nil
}

func (it *ListIterator[T]) Next() T {
	if !it.HasNext() {
		var null T
		return null
	} else {
		node := it.current.item
		it.current = it.current.next
		return node
	}
}

// returns true if the linked list contains e, false otherwise
func (this *Linkedlist[T]) Contains(e T) bool {
	n, _ := this.Search(e)
	return n != -1
}

// return the size of the linked list:
// an int representing the exact number of element contained in the list
func (this *Linkedlist[T]) Size() int {
	if this.IsEmpty() {
		return 0
	} else {
		p := this.head
		size := 0
		for p != nil {
			p = p.next
			size++
		}
		return size
	}
}

//returns the element in the given index position of the linked list,
// returns an error if the given index is negative or greater than this list size
func (this *Linkedlist[T]) Geti(pos int) (T, error) {
	if pos < 0 {
		panic("Negative index for linked list, pos is less than zero")
	}
	var null T
	if this.IsEmpty() {
		return null, errors.New("the list is empty")
	} else {
		p := this.head

		for i := 0; p != nil && i < pos; i++ {
			p = p.next
		}
		if p != nil {
			return p.item, nil
		} else {
			return null, errors.New("could not find element value i is bigger than the linkedList size")
		}
	}

}

// return true if the list is empty, false otherwise
func (this *Linkedlist[T]) IsEmpty() bool {
	return this.head == nil
}

//Attach the given element to the head of the linked list
func (this *Linkedlist[T]) Cons(e T) {
	if this.IsEmpty() {
		this.head = newNode(e)
	} else {
		oldHead := this.head
		this.head = newNode(e)
		this.head.next = oldHead
	}
}

// append the element e to the list tail
func (this *Linkedlist[T]) Append(e T) {
	if this.IsEmpty() {
		this.Cons(e)
	} else {
		p := this.head
		var prec *listNode[T] = nil
		for p != nil {
			prec = p
			p = p.next
		}
		newNode := newNode(e)
		prec.next = newNode
	}

}

// return the position of the first occurence of e in this linked list,
// returns -1 and an error if the element is not present
func (this *Linkedlist[T]) Search(e T) (int, error) {
	if !this.IsEmpty() {
		i := 0
		p := this.head
		for p != nil && p.item != e {
			p = p.next
			i++
		}
		if p == nil {
			return -1, errors.New("element not found")
		} else {
			return i, nil
		}
	} else {
		return -1, errors.New("element not found the list is empty")
	}
}

//remove the element in the given index, if the given index is
//greater than the list size it does nothing
func (this *Linkedlist[T]) Removei(pos int) {
	if !this.IsEmpty() {
		p := this.head
		var prec *listNode[T]
		i := 0
		for p != nil && i < pos {
			prec = p
			p = p.next
			i++
		}
		if p != nil {
			if prec == nil {
				this.head = p.next
			} else {
				prec.next = p.next
			}
		}
	}
}

//remove the first occurrance of the element e in the list
//or does nothing if the element is not present in the linked list
func (this *Linkedlist[T]) Remove(e T) {
	if !this.IsEmpty() {
		p := this.head
		var prec *listNode[T]
		for p != nil && p.item != e {
			prec = p
			p = p.next
		}

		if p != nil {
			if prec == nil {
				this.head = p.next
			} else {
				prec.next = p.next
			}
		}
	}
}

func (this *Linkedlist[T]) Clear() {
	this.head = nil
}

func (this Linkedlist[T]) String() string {
	if !this.IsEmpty() {
		p := this.head
		s := "[" + fmt.Sprint(p.item)
		for p.next != nil {
			p = p.next
			s += " -> " + fmt.Sprint(p.item)
		}
		s += "]"
		return s
	} else {
		return "[ ]"
	}

}

func (this *Linkedlist[T]) F(k int) {
	f(this.head, k)
}

func f[T comparable](p *listNode[T], k int) int {
	var a int
	if p == nil {
		return 0
	}
	a = 1 + f(p.next, k)
	if a == k {
		fmt.Println(p.item)
	}
	return a
}
