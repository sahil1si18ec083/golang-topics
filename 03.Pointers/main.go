package main

import "fmt"

func calculateAdultAge(age *int) int {
	*age = *age - 18
	return *age

}
func increaseAgeBy(ageptr *int, fact int) {
	*ageptr = *ageptr + fact
}

func main() {
	age := 90

	ageptr := &age

	increaseAgeBy(ageptr, 2)

	ans := calculateAdultAge(ageptr)
	fmt.Println(ans)

}
