package main

import "fmt"

// Definición de la estructura de un nodo
type Node struct {
	Value string
	Next  *Node
}

// Definición de la estructura de la LinkedList
type LinkedList struct {
	Head *Node
}

// Método para imprimir la LinkedList de manera recursiva
func (ll *LinkedList) printRecursive() {
	if ll.Head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	fmt.Println(ll.Head.Value)
	ll.Head = ll.Head.Next
	ll.printRecursive()
}

func main() {
	a := &Node{Value: "A", Next: nil}
	b := &Node{Value: "B", Next: nil}
	c := &Node{Value: "C", Next: nil}
	d := &Node{Value: "D", Next: nil}

	a.Next = b
	b.Next = c
	c.Next = d

	ll := LinkedList{Head: a}

	// Llamar al método printRecursive
	fmt.Println("Printing LinkedList:")
	ll.printRecursive()

}
