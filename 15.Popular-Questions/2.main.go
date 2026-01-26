// The main problem discussed in the video is to
//  write a program that prints numbers from 1 to 10 (0:48).
// The key constraint is to achieve this using two goroutines in sequence

// package main

// import "fmt"

// func main() {
// 	fmt.Println("hello")
// 	oddchan := make(chan bool)
// 	evenchan := make(chan bool)
// 	donechan := make(chan bool)

// 	go func() {
// 		for i := 1; i <= 9; i = i + 2 {
// 			<-oddchan
// 			fmt.Println(i)
// 			evenchan <- true

// 		}

// 	}()
// 	go func() {
// 		for i := 2; i <= 10; i = i + 2 {
// 			<-evenchan
// 			fmt.Println(i)
// 			if i == 10 {
// 				donechan <- true
// 				return
// 			}
// 			oddchan <- true

// 		}

// 	}()
// 	oddchan <- true

// 	<-donechan

// }
