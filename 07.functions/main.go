package main

import "fmt"

func transformNumbers(nums *[]int, fn func(int) int) []int {
	res := []int{}
	for _, val := range *nums {
		res = append(res, fn(val))
	}
	return res

}
func double(a int) int {
	return a * 2

}
func triple(a int) int {
	return a * 4

}
func sum(a int) (f func(int) int) {
	return func(b int) int {
		return a + b
	}

}
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)

}
func variablic(nums ...int) int {
	res := 0
	for _, val := range nums {
		res = res + val
	}
	return res

}
func main() {
	nums := []int{1, 2, 3, 4}

	res_double := transformNumbers(&nums, double)
	res_triple := transformNumbers(&nums, triple)
	fmt.Println(res_double)
	fmt.Println(res_triple)

	f := sum(3)
	d := f(4)
	fmt.Println(d)

	r := fact(5)
	fmt.Println(r)
	temp := []int{1, 2, 3}
	sum := variablic(3, 4, 5)
	sum2 := variablic(2, 3, 4)

	sum3 := variablic(temp...)
	fmt.Println(sum, sum2, sum3)

}
