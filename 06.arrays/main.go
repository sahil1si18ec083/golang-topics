package main

import "fmt"

// type Product struct {
// 	Price float32
// 	id    int
// }

func main() {
	// var arr [4]int

	// arr[0] = 99
	// arr2 := []int{1, 2, 3, 4, 5}

	// part := arr2[1:3]
	// fmt.Println(part)

	// prices := []float64{10.99, 8.99}
	// prices[1] = 9.99

	// updatedPrices := append(prices, 5.99)
	// fmt.Println(updatedPrices, prices)

	// hobbies := [3]string{"sports", "cooking", "reading"}
	// fmt.Print(hobbies[0])
	// h2 := hobbies[1:]
	// fmt.Print(h2)
	// mainhob := hobbies[:2]
	// fmt.Println(cap(mainhob))
	// mainhob = mainhob[1:]
	// fmt.Println(mainhob)
	// g := []string{"a", "b"}
	// g = append(g, "c")
	// fmt.Println(g)
	// parr := []Product{{Price: 67.99, id: 7}, {Price: 67.99, id: 7}}
	// fmt.Println(parr)

	mp := map[string]int{"a": 90, "b": 67}
	fmt.Println(mp)
	val, ok := mp["a"]
	delete(mp, "b")
	fmt.Println(val, ok)

	arr3 := make([]int, 1)
	fmt.Println(len(arr3)) // 1
	arr3[0] = 777
	for index, val := range arr3 {
		fmt.Print(index, val)
	}
	for key, val := range mp {
		fmt.Println(key, val)
	}

}
