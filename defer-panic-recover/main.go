package main

import "fmt"

func main() {
	// Utilizamos defer para asegurarnos de que la función cleanup() se llame justo antes de que ocurra el pánico
	defer cleanup()

	fmt.Println("Inicio del programa")

	// Simulamos una operación que podría causar un pánico
	if err := doSomething(); err != nil {
		panic(err)
	}

	fmt.Println("Fin del programa")
}

func doSomething() error {
	// Aquí puedes realizar alguna operación que podría causar un pánico, por ejemplo, una división por cero
	// Para este ejemplo, solo simulamos un error
	return fmt.Errorf("algo salió mal")
}

func cleanup() {
	// Aquí puedes incluir acciones de limpieza, como cerrar conexiones de red o archivos abiertos
	fmt.Println("Realizando limpieza...")
}
