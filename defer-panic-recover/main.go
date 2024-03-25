package main

import "fmt"

func main() {
	// Utilizamos defer para asegurarnos de que la función cleanup() se llame al final de la ejecución
	defer cleanup()

	fmt.Println("Inicio del programa")

	// Simulamos una operación que podría causar un pánico
	doSomething()

	fmt.Println("Fin del programa")
}

func doSomething() {
	// Utilizamos defer para recuperarnos del pánico si ocurre
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperando de pánico:", r)
		}
	}()

	// Simulamos un pánico
	panic("Algo salió mal")
}

func cleanup() {
	// Aquí puedes incluir acciones de limpieza, como cerrar conexiones de red o archivos abiertos
	fmt.Println("Realizando limpieza...")
}
