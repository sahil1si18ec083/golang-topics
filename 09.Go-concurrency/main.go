package main

import (
	"fmt"
	"time"
)

func greet(phrase string, donech chan<- bool) {
	donech <- true
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, donech chan<- bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	donech <- true

}

func main() {
	donech := make(chan bool)

	go greet("Nice to meet you!", donech)
	go greet("How are you?", donech)
	go slowGreet("How ... are ... you ...?", donech)
	go greet("I hope you're liking the course!", donech)
	for i := 0; i < 4; i++ {
		<-donech

	}
}
