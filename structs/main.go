package main

import "fmt"

type User struct {
	Username string
	Country  string
	City     string
}

func main() {
	user := User{Username: "BHoffman", Country: "Polska", City: "Lima"}
	// Acceder a cada campo del struct
	fmt.Println("Username:", user.Username)
	fmt.Println("Country:", user.Country)
	fmt.Println("City:", user.City)
}
