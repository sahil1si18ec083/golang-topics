package cmdmanager

import "fmt"

type Cmdmanager struct {
}

func (cmd *Cmdmanager) ReadFile() ([]string, error) {
	fmt.Println("Please enter your prices confirm each price with Enter")
	var prices []string
	for {
		var price string
		fmt.Print("Price :")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil

}

func (cmd *Cmdmanager) WriteJSON(data interface{}) error {
	fmt.Println(data)

	return nil
}
