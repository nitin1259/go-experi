package linkedlist

type LinkNode[T any] struct {
	Val  T
	Next *LinkNode[T]
}

func ReverseLinkedList(head *LinkNode[int])*LinkNode[int]{

		var curr, prev *LinkNode[int]
		curr = head

	 for curr!=nil{
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	 }
	 return prev
}