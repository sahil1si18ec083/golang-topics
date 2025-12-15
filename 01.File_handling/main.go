package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getUserInput(balance *float64) error {
	fmt.Scan(balance)
	if *balance < 0 || *balance == 0 {
		return errors.New("invalid Input")
	}

	return nil

}
func main() {
	fileName := "balance.txt"

	var balance float64
	err := getUserInput(&balance)
	if err != nil {
		fmt.Print("User input for balance is invalid")
		return
	}

	balanceText := fmt.Sprintf("%f", balance)
	permissions := os.FileMode(0644)
	data := []byte(balanceText)
	os.WriteFile(fileName, data, permissions)

	bytearr, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)

	}
	balanceString := string(bytearr)
	finalValue, err := strconv.ParseFloat(balanceString, 64)
	if err != nil {
		panic("cannot continue")
	}
	fmt.Printf("%f,%T", finalValue, finalValue)
}
