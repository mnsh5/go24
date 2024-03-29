package linkedlist

// tres operaciones tipicias en LinkedList -> Add(), Remove(), Get()
type Node struct {
	Current string // nodo actual
	Head    string
	Next    string
	Tail    string
}

// Array tiene muchos elementos, y LinkedList tiene muchos nodes
// Un array se almacena de forma contigua en la memoria, los elementos se almacenen uno al lado del otro
// Esto puede traer consecuencias en el tiempo de ejecucion de la memoria en diferentes operaciones del array
// LinkedList no necesita que los nodos sean contiguas en la memoria
