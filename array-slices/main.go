package main

import "fmt"

func main() {
	// Definir un array de tamaño fijo
	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	// Imprimir el array
	fmt.Println("Array:", array)

	arr := []int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// slice de 0 elementos con 0 de capacidad
	nums := make([]int, 0, 0)
	for i := 0; i < len(nums); i++ {
		nums = append(nums, i)
		fmt.Println(nums[i])
	}

}
