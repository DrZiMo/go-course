package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetBalanceFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Printf("Successfully deposited $%.2f. New balance is: $%.2f\n", depositAmount, accountBalance)
			fileops.WriteBalanceToFile(accountBalanceFile, accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds!")
				continue
			}

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Printf("Successfully withdrew $%.2f. New balance is: $%.2f\n", withdrawAmount, accountBalance)
			fileops.WriteBalanceToFile(accountBalanceFile, accountBalance)
		case 4:
			fmt.Println("Thank you for banking with us. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		// if choice == 1 {
		// 	fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
		// } else if choice == 2 {
		// 	var depositAmount float64
		// 	fmt.Print("Enter amount to deposit: ")
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		continue
		// 	}

		// 	accountBalance += depositAmount
		// 	fmt.Printf("Successfully deposited $%.2f. New balance is: $%.2f\n", depositAmount, accountBalance)
		// } else if choice == 3 {
		// 	var withdrawAmount float64
		// 	fmt.Print("Enter amount to withdraw: ")
		// 	fmt.Scan(&withdrawAmount)
		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Insufficient funds!")
		// 		continue
		// 	}

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawAmount
		// 	fmt.Printf("Successfully withdrew $%.2f. New balance is: $%.2f\n", withdrawAmount, accountBalance)
		// } else if choice == 4 {
		// 	fmt.Println("Thank you for banking with us. Goodbye!")
		// 	break
		// } else {
		// 	fmt.Println("Invalid choice. Please try again.")
		// }
	}
}
