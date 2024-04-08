package main

import "fmt"

// Función sortResult ordena un slice de enteros usando el algoritmo de ordenamiento de burbuja
func sortResult(result []int) {
	// Itera sobre todos los elementos menos el último
	for i := 0; i < len(result)-1; i++ {
		// Itera sobre los elementos a partir del elemento i+1
		for j := i + 1; j < len(result); j++ {
			// Compara los elementos en las posiciones i y j
			if result[i] > result[j] {
				// Si el elemento en la posición i es mayor que el elemento en la posición j, los intercambia
				result[i], result[j] = result[j], result[i]
			}
		}
	}
}

// Función SortMerge fusiona y ordena dos arrays de enteros
func SortMerge(arr1, arr2 []int) []int {
	// Calcula la longitud total del resultado
	totalLen := len(arr1) + len(arr2)
	// Crea un slice vacío con una capacidad inicial igual a la longitud total del resultado
	result := make([]int, 0, totalLen)

	// Inicializa índices para recorrer arr1 y arr2
	i, j := 0, 0

	// Fusiona los elementos de arr1 y arr2 en el resultado
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			// Si el elemento actual de arr1 es menor, lo agrega al resultado y mueve el índice i
			result = append(result, arr1[i])
			i++
		} else {
			// Si el elemento actual de arr2 es menor o igual, lo agrega al resultado y mueve el índice j
			result = append(result, arr2[j])
			j++
		}
	}

	// Agrega los elementos restantes de arr1 al resultado
	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}

	// Agrega los elementos restantes de arr2 al resultado
	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}

	// Ordena el resultado utilizando la función sortResult
	sortResult(result)

	return result
}

func main() {
	msa1 := []int{0, 3, 4, 31, 12}
	msa2 := []int{5, 35, 24, 2, 15}
	sortedResult := SortMerge(msa1, msa2)
	fmt.Println(sortedResult)
}
