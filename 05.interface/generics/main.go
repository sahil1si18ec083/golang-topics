package main

import "fmt"

// func sum(a interface{}, b interface{}) interface{} {
// 	aint, aisint := a.(int)
// 	bint, bisint := b.(int)
// 	if aisint && bisint {
// 		return aint + bint
// 	}
// 	astring, aistring := a.(string)
// 	bstring, bistring := b.(string)
// 	if aistring && bistring {
// 		return astring + bstring
// 	}
// 	return nil

// }
func sum[T int | float64 | string](a T, b T) T {
	return a + b

}
func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum("4", "9"))

}
