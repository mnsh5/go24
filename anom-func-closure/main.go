package anomfuncclosure

import "fmt"

func AnomFunc() {
	// Definición de una función anónima que imprime un mensaje
	saludar := func() {
		fmt.Println("¡Hola desde la función anónima!")
	}
	// Llamada a la función anónima
	saludar()
}

func Closure() {
	mensaje := "Hola desde el closure!"

	// Definición del closure que captura la variable mensaje
	imprimirMensaje := func() {
		fmt.Println(mensaje)
	}

	// Llamada al closure
	imprimirMensaje()
}
