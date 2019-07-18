package main

import "fmt"

type Node struct{
	Value int
	Prev *Node
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func(l *LinkedList) GetHead() *Node {
	return l.Head
}
func(l *LinkedList) Insert(value int) {
	current := l.GetHead()
	newNode := Node{value,nil,nil}
	for current.Next != nil {
		current = current.Next
	}
	newNode.Prev = current
	current.Next = &newNode
}

func (l *LinkedList) Print() {
	current := l.GetHead()
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func (l *LinkedList) PrintReverse() {
	current := l.GetHead()
	for current.Next != nil {
		current = current.Next
	}
	for current != nil {
		fmt.Println(current.Value)
		current = current.Prev
	}
}

func main() {
	head:=Node {1,nil,nil}
	linkedList := LinkedList {&head}
	linkedList.Insert(2)
	linkedList.Insert(3)
	linkedList.Insert(4)
	linkedList.Insert(5)
	linkedList.PrintReverse()
}