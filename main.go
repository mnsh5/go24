package main

import (
	"fmt"

	anomfuncclosure "github.com/mnsh5/go24/anom-func-closure"
	controlstructures "github.com/mnsh5/go24/control_structures"
	"github.com/mnsh5/go24/functions"
	"github.com/mnsh5/go24/loops"
	"github.com/mnsh5/go24/variables"
)

func main() {
	fmt.Println("Hello, World!")
	variables.DisplayVariables()

	h := functions.Hello("Binyamin Hoffman", "Polska")
	fmt.Println(h)

	nums := []int{1, 2, 3, 4, 5}
	res := functions.Sum(nums)
	fmt.Println("The sum of the numbers is:", res)

	controlstructures.ControlStructures()
	loops.Loops()

	// files.Files()
	anomfuncclosure.AnomFunc()
	anomfuncclosure.Closure()
}
