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

	// For
	countries := []string{"Polska", "Deutschland", "Argentina"}
	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}

	for _, country := range countries {
		switch country {
		case "Polska":
			fmt.Println("Polish is spoken in", country)
		case "Deutschland":
			fmt.Println("German is spoken in", country)
		case "Argentina":
			fmt.Println("Spanish is spoken in", country)
		default:
			fmt.Println("Language for", country, "is unknown")
		}
	}
}
