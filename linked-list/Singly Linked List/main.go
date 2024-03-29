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

func (list *LinkedList) Add(value string) {
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

func (list *LinkedList) Get(value string) string {
	if list.Head == nil {
		fmt.Println("Linked list is empty")
		return ""
	}

	currentNode := list.Head
	for currentNode != nil {
		if currentNode.Value == value {
			return currentNode.Value
		}
		currentNode = currentNode.Next
	}
	fmt.Println("Value not found in the linked list")
	return ""
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

	list.Add("a")
	list.Add("b")
	list.Add("c")
	list.Add("Square")

	list.PrintLinkedList()
	fmt.Println("Get List:", list.Get("Square"))
}
