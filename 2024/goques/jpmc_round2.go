package goques

import "fmt"


type LinkNode struct{
	val int
	next *LinkNode
}

// reverser
// -> nil
// nil <- 2   <- p5 -> n,c7 -> 3<- 4p


func reverseALinkedList(head *LinkNode) *LinkNode{

	var curr, prev *LinkNode
	curr= head
	prev = nil

	for curr!=nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev

}

func main_linkedList() {

	head := &LinkNode{
		val: 2,
		next: &LinkNode{
			val: 5,
			next: &LinkNode{
				val: 7,
				next: &LinkNode{
					val: 4,
					next: nil,
				},
			},
		},
	}

	// fmt.Println(head)

	printLinkedlist(head)
	fmt.Println("---reverse linked oist---")
	printLinkedlist(reverseALinkedList(head))
	
}

func printLinkedlist(head *LinkNode){
	curr := head
	for curr != nil {
		fmt.Println(curr.val)
		curr = curr.next
	}
}