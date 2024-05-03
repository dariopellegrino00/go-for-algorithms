package circularlist

import (
	"fmt"
)

type CircNode struct {
	value int
	next  *CircNode
	prev  *CircNode
}

type Circularlist struct {
	head *CircNode
}

func (list *Circularlist) Append(e int) {
	var r CircNode
	r.value = e
	if list.head == nil {
		r.next = &r
		r.prev = r.next
		list.head = &r
	} else {
		if list.head == list.head.prev {
			list.head.next = &r
			r.prev = list.head
		} else {
			prev := list.head.prev
			prev.next = &r
			r.prev = prev
		}
		list.head.prev = &r
		r.next = list.head
	}
}

// TODO match old f return domain
func (list *Circularlist) F2(k int) {
	var p *CircNode
	if list.head != nil && k > 0 {
		p = list.head.prev
		for i := 0; i < k && p != list.head; i++ {
			p = p.prev
		}
	}

	if p != nil {
		fmt.Println(p.value)
	}
}

//not safe but for exercise purpose (to use the func below)
func (list *Circularlist) GetPoiter(p int) *CircNode {
	node := list.head
	for i := 0; node != nil && i < p; i++ {
		node = node.next
	}

	return node
}

// we assume that the node p is part of the list
func (list *Circularlist) Move(p *CircNode) {
	if list.head != nil && p != nil && p.next != p.prev {
		if p.value > 0 {
			for i := 0; i < p.value; i++ {
				p.ShiftForward()
			}
		} else {
			// it goes into loop find why
			for i := p.value; i < 0; i++ {
				// back shift
				p.ShiftBackwards()
			}
		}
	}
}

func (node *CircNode) ShiftForward() {
	// shift next node and backup his next node
	node.next.prev = node.prev
	oldNext := node.next
	newNext := node.next.next
	node.next.next = node
	// update prev node
	node.prev.next = node.next
	// update this node
	node.next = newNext
	node.prev = oldNext
	newNext.prev = node

}

func (node *CircNode) ShiftBackwards() {
	// shift back node and backup his prev node
	oldPrev := node.prev
	newPrev := node.prev.prev
	oldPrev.next = node.next
	oldPrev.prev = node
	// update next node
	node.next.prev = node.prev
	// update this node
	node.prev = newPrev
	node.next = oldPrev
	newPrev.next = node
}

//ALL the priting below...

func (list Circularlist) String() string {

	if list.head != nil {
		s := "[" + fmt.Sprint(list.head.value)
		p := list.head
		for p != nil && p.next != list.head {
			p = p.next
			s += " -> " + fmt.Sprint(p.value)
		}
		s += "]"
		return s
	} else {
		return "[ ]"
	}
}

// we assume that the circular list cotains a zero to start printing from
func (list *Circularlist) PrintFromZero() {
	if list.head != nil {
		var zero *CircNode
		p := list.head
		for p != nil && p.next != list.head {
			if p.value == 0 {
				zero = p
				break
			}
			p = p.next
		}

		p = zero.next
		fmt.Print(zero.value, " ")
		for p != nil && p != zero {
			fmt.Print(p.value, " ")
			p = p.next
		}
		fmt.Println()
	}
}

func (node CircNode) String() string {

	return fmt.Sprint("{", node.value, ", prev: ", node.prev.value, ", next: ", node.next.value, " }")
}

func (list *Circularlist) String2() string {
	if list.head != nil {
		s := fmt.Sprintln(list.head)
		p := list.head
		for p != nil && p.next != list.head {
			p = p.next
			s += fmt.Sprintln(p)
		}
		return s
	} else {
		return ""
	}
}
