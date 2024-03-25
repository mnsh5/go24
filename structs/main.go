package main

import "fmt"

type User struct {
	Username string
	Country  string
	City     string
}

func main() {
	users := []User{
		{Username: "BHoffman", Country: "UK", City: "Liverpool"},
		{Username: "LCommendatore", Country: "AR", City: "Buenos Aires"},
	}

	// Recorrer el slice de usuarios
	for _, user := range users {
		fmt.Println("Username:", user.Username)
		fmt.Println("Country:", user.Country)
		fmt.Println("City:", user.City)
		fmt.Println()
	}

}
