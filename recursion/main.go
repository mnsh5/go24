package main

import "fmt"

func factorial(n int) int {
	// Caso base: factorial de 0 es 1
	if n == 0 {
		return 1
	}
	// Llamada recursiva: factorial de n es n * factorial(n-1)
	return n * factorial(n-1)
}

func imprimirNumeros(n int) {
	// Caso base: cuando n llega a 0, terminamos la recursión
	if n == 0 {
		return
	}

	// Imprimimos el número actual
	fmt.Println(n)

	// Llamada recursiva con n-1
	imprimirNumeros(n - 1)
}

func main() {
	// Calcular el factorial de 5
	resultado := factorial(5)
	fmt.Println("El factorial de 5 es:", resultado)

	imprimirNumeros(2)
}
