package controlstructures

import (
	"fmt"
	"runtime"
)

func ControlStructures() {
	// Example of if
	country := "Polska"
	if country == "Polska" {
		fmt.Println("Stolicą Polski jest Warszawa")
	} else {
		fmt.Println("You are from Argentina")
	}

	if os := runtime.GOOS; os == "linux" {
		fmt.Println("Linux!")
	} else {
		fmt.Println("Any OS!")
	}
}
