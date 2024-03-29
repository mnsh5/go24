package main

import "fmt"

// Singly Linked List

type Node struct {
	Value string
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Insert(value string) {
	newNode := &Node{Value: value, Next: nil}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	currentNode := list.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode
}

func (list *LinkedList) PrintLinkedList() {
	currentNode := list.Head

	if currentNode == nil {
		fmt.Println("Linked list is empty")
		return
	}

	for currentNode != nil {
		fmt.Print("Linked list: ")
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}
}

func main() {
	list := LinkedList{}

	list.Insert("a")
	list.Insert("b")
	list.Insert("c")
	list.Insert("d")

	list.PrintLinkedList()
}
