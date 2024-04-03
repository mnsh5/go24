package main

import "fmt"

// Array -> organiza los elementos de forma secuencial
// Metodos para un Array -> Insert(), Delete(), append(), Lookup(), Get()

type MyArray struct {
	length int
	data   map[int]string
}

func NewArray() *MyArray {
	return &MyArray{
		length: 0,
		data:   make(map[int]string),
	}
}

func (a *MyArray) get(index int) string {
	return a.data[index]
}

func (a *MyArray) append(item string) {
	a.data[a.length] = item
	a.length++
}

func main() {
	arr := NewArray()
	arr.append("Madrid")
	arr.append("Lisboa")
	arr.append("Berlin")

	fmt.Println(arr)
	fmt.Println(arr.get(0))
	fmt.Println(arr.get(1))
	fmt.Println(arr.get(2))
}
