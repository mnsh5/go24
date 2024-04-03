package main

import (
	"fmt"
	"reflect"
)

func reverse(s string) string {
	if reflect.TypeOf(s).Kind() != reflect.String || s == "" || len(s) < 2 {
		return s
	}

	backwards := make([]byte, 0, len(s)) // Inicializar un slice de bytes vacío con capacidad suficiente
	for i := len(s) - 1; i >= 0; i-- {
		backwards = append(backwards, s[i]) // Agregar caracteres al final del slice
	}

	return string(backwards) // Convertir el slice de bytes en una cadena
}

func main() {
	// Cadena original
	str1 := "Leverkusen"
	str2 := "Buenos Aires"
	str3 := "Augsburgo"

	// Revertir la cadena
	reversed1 := reverse(str1)
	reversed2 := reverse(str2)
	reversed3 := reverse(str3)

	// Imprimir la cadena revertida
	fmt.Println("Cadena original:", str1)
	fmt.Println("Cadena revertida:", reversed1)

	fmt.Println("Cadena original:", str2)
	fmt.Println("Cadena revertida:", reversed2)

	fmt.Println("Cadena original:", str3)
	fmt.Println("Cadena revertida:", reversed3)
}
