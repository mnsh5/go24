package main

import "fmt"

// Array tiene muchos elementos, y LinkedList tiene muchos nodes
// Un array se almacena de forma contigua en la memoria, los elementos se almacenen uno al lado del otro
// Esto puede traer consecuencias en el tiempo de ejecucion de la memoria en diferentes operaciones del array
// LinkedList no necesita que los nodos sean contiguas en la memoria
// tres operaciones tipicias en LinkedList -> Add(), Remove(), Get()

// Definición de la estructura de un nodo de tipo generico
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// Definición de la estructura de la LinkedList de tipo generico
type LinkedList[T any] struct {
	Head *Node[T]
}

func (ll *LinkedList[T]) printLinkedList() {
	currentNode := ll.Head

	if currentNode == nil {
		fmt.Println("Linked list is empty")
		return
	}
	for currentNode != nil {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}

}

func main() {
	a := &Node[string]{Value: "A", Next: nil}
	b := &Node[string]{Value: "B", Next: nil}
	c := &Node[string]{Value: "C", Next: nil}
	d := &Node[string]{Value: "D", Next: nil}

	a.Next = b
	b.Next = c
	c.Next = d

	// Crear una instancia de LinkedList
	ll := LinkedList[string]{Head: a}

	// Llamar al método printLinkedList de la instancia de LinkedList
	ll.printLinkedList()
}
