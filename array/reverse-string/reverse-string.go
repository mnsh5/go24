package main

import (
	"fmt"
)

func reverseString(s string) string {
	// Convertir la cadena a un slice de bytes
	bytes := []byte(s)

	// Obtener la longitud de la cadena
	length := len(bytes)

	// Intercambiar los elementos del slice de bytes
	for i := 0; i < length/2; i++ {
		// Intercambiar los elementos en posiciones i y length-i-1
		bytes[i], bytes[length-i-1] = bytes[length-i-1], bytes[i]
	}

	// Convertir el slice de bytes invertido de nuevo a una cadena
	return string(bytes)
}

func main() {
	// Cadena original
	str := "München"

	// Revertir la cadena
	reversed := reverseString(str)

	// Imprimir la cadena revertida
	fmt.Println("Cadena original:", str)
	fmt.Println("Cadena revertida:", reversed)
}
