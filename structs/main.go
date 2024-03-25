package main

import (
	"fmt"

	"github.com/mnsh5/go24/structs/models"
)

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

	// Convertir el slice de User a []models.User
	var modelsUsers []models.User
	for _, user := range users {
		modelsUsers = append(modelsUsers, models.User(user))
	}

	// Crear una nueva instancia de models.User
	u := new(models.User)

	// Agregar usuarios al slice existente en la instancia de models.User
	u.AddUser(modelsUsers)

	// Imprimir el slice de usuarios actualizado
	fmt.Println("Usuarios actualizados:", modelsUsers)

}
