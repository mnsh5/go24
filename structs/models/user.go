package models

type User struct {
	Username string
	Country  string
	City     string
}

// Agregar un nuevo usuario al slice existente
func (u *User) AddUser(users []User) []User {
	// Agregar el nuevo usuario al slice
	users = append(users, *u)
	// Devolver el slice actualizado
	return users
}
