package utils

import (
	"fmt"
	"os"
)

func WriteBalanceToFile(accountBalanceFile string, balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}
