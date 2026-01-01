package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetBalanceFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to read account balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to convert string to float64.")
	}

	return balance, nil
}

func WriteBalanceToFile(fileName string, balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}
