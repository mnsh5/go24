package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["foo"] = 7
	m["bar"] = 13
	fmt.Println("map:", m)

	for clave, valor := range m {
		fmt.Println("Clave:", clave, "Valor:", valor)
	}

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	for clave, valor := range n {
		fmt.Println("Clave:", clave, "Valor:", valor)
	}
}
