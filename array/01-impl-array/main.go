package main

import "fmt"

// Array -> organiza los elementos de forma secuencial
// Metodos para un Array -> Insert(), Delete(), append(), Lookup(), Get()

type MyArray struct {
	length int
	data   map[int]string
}

func NewMyArray() *MyArray {
	return &MyArray{
		length: 0,
		data:   make(map[int]string),
	}
}

func (a *MyArray) Get(index int) string {
	return a.data[index]
}

func main() {
	arr := NewMyArray()
	arr.data[0] = "Madrid"
	arr.data[1] = "Lisboa"

	fmt.Println(arr)
	fmt.Println(arr.Get(0))
	fmt.Println(arr.Get(1))
}
