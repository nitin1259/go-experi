package linkedlist

type LinkNode[T any] struct {
	Val  T
	Next *LinkNode[T]
}
