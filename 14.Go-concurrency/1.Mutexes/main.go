package main

import (
	"fmt"
	"sync"
)

func BuyTicket(wg *sync.WaitGroup, userid int, tickets *int, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()

	if *tickets > 0 {
		*tickets-- // user purchases
		fmt.Println("user purchsed tickets left", *tickets)

	} else {
		fmt.Println("No ticket left")
	}

}

func main() {

	var tickets int = 500
	var mu sync.Mutex

	var wg sync.WaitGroup

	for userId := 0; userId < 505; userId++ {
		wg.Add(1)

		// buy ticket

		go BuyTicket(&wg, userId, &tickets, &mu)
	}

	wg.Wait()

	fmt.Println("EOP")

}

// Note always send waitground variable and mutex variable  to  a  function  as a pointer and not pass by value
