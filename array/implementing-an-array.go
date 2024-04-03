package main

import "fmt"

// Array -> organiza los elementos de forma secuencial
// Metodos para un Array -> Insert(), Delete(), Append(), Lookup(), Get()

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

func (a *MyArray) pop() string {
	lastItem := a.data[a.length-1]
	delete(a.data, a.length-1)
	a.length--
	return lastItem
}

func (a *MyArray) getAll() []string {
	// slice -> primer elemento es un array de string, el cero indica el valor en que iniciamos el slice, y el len sera el tamaño del slice en base a los datos guardados
	value := make([]string, 0, len(a.data))
	for _, v := range a.data {
		value = append(value, v)
	}
	return value
}

func (a *MyArray) delete(index int) string {
	item := a.data[index]
	a.shiftItems(index)
	return item
}

func (a *MyArray) shiftItems(index int) {
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	delete(a.data, a.length-1)
	a.length--
}

func main() {
	arr := NewArray()
	arr.append("Madrid")
	arr.append("Lisboa")
	arr.append("Berlin")
	arr.append("Penguin")
	arr.append("Amsterdam")
	arr.append("Buenos Aires")
	arr.append("Thomas Müller")

	fmt.Println(arr.get(0))
	fmt.Println(arr.get(1))
	fmt.Println(arr.get(2))
	fmt.Println(arr.get(3))
	fmt.Println(arr.get(4))
	fmt.Println(arr.get(5))

	fmt.Println(arr.getAll())

	arr.pop()

	fmt.Println(arr.getAll())

	arr.delete(3)
	fmt.Println(arr.getAll())
	fmt.Println(arr)

}
