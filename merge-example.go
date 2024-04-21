package main

import (
	"fmt"
	"sort"
)

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Iterar sobre ambos arrays mientras haya elementos en ambos
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i]) // Agregar el elemento del array izquierdo al resultado
			i++                              // Mover al siguiente elemento del array izquierdo
		} else {
			result = append(result, right[j]) // Agregar el elemento del array derecho al resultado
			j++                               // Mover al siguiente elemento del array derecho
		}
	}

	// Agregar los elementos restantes del array izquierdo al resultado
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Agregar los elementos restantes del array derecho al resultado
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func main() {
	arr := []int{1, 95, 6, 45, 3, 789, 0}
	l := []int{54, 2, 15, 5, 4}

	// Ordenar los dos arrays antes de pasarlos a la función merge
	sort.Ints(arr)
	sort.Ints(l)

	fmt.Println(merge(l, arr))
}
