package main

import (
	"fmt"

	"github.com/mnsh5/go24/functions"
	"github.com/mnsh5/go24/variables"
)

func main() {
	fmt.Println("Hello, World!")
	variables.DisplayVariables()

	h := functions.Hello("Binyamin Hoffman", "Polsce")
	fmt.Println(h)
}
