package functions

import "fmt"

func Hello(name string, country string) string {
	return fmt.Sprintf("Witaj %s, witaj w %s", name, country)
}
