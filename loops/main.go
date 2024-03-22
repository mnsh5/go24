package loops

import "fmt"

func Loops() {
	countries := []string{"Polska", "Deutschland", "Argentina"}
	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(nums))
}

func Sum(num []int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}
