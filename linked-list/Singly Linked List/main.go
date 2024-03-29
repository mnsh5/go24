package main

import "fmt"

// Singly Linked List
// https://edwinsiby.medium.com/understanding-linked-lists-in-go-a-comprehensive-guide-for-all-skill-levels-1e3d91a24d08

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

func (list *LinkedList) Remove(value string) {
	if list.Head == nil {
		fmt.Println("Linked list is empty")
		return
	}

	// Si el valor a eliminar está en el primer nodo
	if list.Head.Value == value {
		list.Head = list.Head.Next
		return
	}

	// Buscar el nodo que contiene el valor a eliminar
	prevNode := list.Head
	currentNode := list.Head.Next
	for currentNode != nil {
		if currentNode.Value == value {
			prevNode.Next = currentNode.Next
			return
		}
		prevNode = currentNode
		currentNode = currentNode.Next
	}

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

	list.Add("Messi")
	list.Add("Pokemon")
	list.Add("Thomas Müller")
	list.Add("Square")

	list.PrintLinkedList()
	fmt.Println("Get List:", list.Get("Square"))
	list.Remove("Pokemon")
	list.PrintLinkedList()
}
