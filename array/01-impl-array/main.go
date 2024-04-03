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

func (a *MyArray) getAll() []string {
	// slice -> primer elemento es un array de string, el cero indica el valor en que iniciamos el slice, y el len sera el tamaño del slice en base a los datos guardados
	value := make([]string, 0, len(a.data))
	for _, v := range a.data {
		value = append(value, v)
	}
	return value
}

func main() {
	arr := NewArray()
	arr.append("Madrid")
	arr.append("Lisboa")
	arr.append("Berlin")
	arr.append("Amsterdam")
	arr.append("Buenos Aires")
	arr.append("Thomas Müller")

	fmt.Println(arr)
	fmt.Println(arr.get(0))
	fmt.Println(arr.get(1))
	fmt.Println(arr.get(2))
	fmt.Println(arr.get(3))
	fmt.Println(arr.get(4))

	fmt.Println(arr.getAll())
}
