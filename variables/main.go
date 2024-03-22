package variables

import "fmt"

func DisplayVariables() {
	// declaracion
	var polska string = "Polska"

	// asignacion
	binyamin := "Binyamin Hoffman"

	message := fmt.Sprintf("Hello, I'm %s from %s", binyamin, polska)
	fmt.Println(message)
}
