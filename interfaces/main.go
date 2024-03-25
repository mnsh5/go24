package main

import (
	"fmt"
)

// Definimos una interfaz llamada Cat con un método Hello()
type Cat interface {
	Hello() string
}

// Creamos una estructura llamada Kitty que implementa la interfaz Cat
type Kitty struct{}

// Implementamos el método Hello() de la interfaz Cat para la estructura Kitty
func (k Kitty) Hello() string {
	return "Kitty says Meow"
}

func main() {
	kitty := Kitty{}
	fmt.Println(kitty.Hello())
}
