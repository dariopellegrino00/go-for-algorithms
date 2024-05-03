package binarytree

type Tree[T comparable] struct {
	value T
	left  *Tree[T]
	right *Tree[T]
}

func (this *Tree[T]) TODO() {

}
