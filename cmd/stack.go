package cmd

type Node struct {
	next *Node
	val  *Timer
}

var top *Node
var length int
var Stack *Node

func (s *Node) Push(value *Timer) {
	node := &Node{
		val: value,
	}
	
	length++
	node.next = top
	top = node
}

func (s *Node) Peek() *Timer {
	return top.val
}

func (s *Node) Pop() *Timer {
	result := top
	top = top.next
	length--
	return result.val
}
