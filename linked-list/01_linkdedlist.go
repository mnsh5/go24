package main

import "fmt"

// Array tiene muchos elementos, y LinkedList tiene muchos nodes
// Un array se almacena de forma contigua en la memoria, los elementos se almacenen uno al lado del otro
// Esto puede traer consecuencias en el tiempo de ejecucion de la memoria en diferentes operaciones del array
// LinkedList no necesita que los nodos sean contiguas en la memoria
// tres operaciones tipicias en LinkedList -> Add(), Remove(), Get()

// Definición de la estructura de un nodo
type Node struct {
	Value string
	Next  *Node
}

// Definición de la estructura de la LinkedList
type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) printLinkedList() {
	currentNode := ll.Head
	for currentNode != nil {
		fmt.Printf("%s -> ", currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Println("nil")
}

func main() {
	a := &Node{Value: "A", Next: nil}

	// Crear una instancia de LinkedList
	ll := LinkedList{Head: a}

	// Llamar al método printLinkedList de la instancia de LinkedList
	ll.printLinkedList()

}
