package functions

import "fmt"

func Hello(name string, country string) string {
	return fmt.Sprintf("Witaj %s, witaj w %s", name, country)
}

func Sum(num []int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}
