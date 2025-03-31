package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func writeBalanceToFile(balance float64) {
	bytes := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(bytes), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return 1000.0, errors.New("Faild to read file.")
	}
	balanceText := string(data)
	val, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000.0, errors.New("Faild to parse stored balance value.")
	}
	return val, nil
}

func main() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Printf("Error %v\n", err)
	}

	writeBalanceToFile(accountBalance)

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// isCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			amount, err := changeBalance("deposit")

			if err != nil {
				fmt.Println(err)
				continue
			}

			accountBalance += amount
			fmt.Println("Your balance now is", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			amount, err := changeBalance("withdraw")

			if err != nil {
				fmt.Println(err)
				continue
			}

			if amount > accountBalance {
				fmt.Println("You don't have enough money.")
				continue
			}

			accountBalance -= amount
			fmt.Println("Your balance now is", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is", accountBalance)
		// } else if choice == 2 {
		// 	amount := changeBalance("deposit")

		// 	if amount <= 0 {
		// 		fmt.Println("Invalid amount.")
		// 		continue
		// 	}

		// 	accountBalance += amount
		// 	fmt.Println("Your balance now is", accountBalance)
		// } else if choice == 3 {
		// 	amount := changeBalance("withdraw")

		// 	if amount <= 0 {
		// 		fmt.Println("Invalid amount.")
		// 		continue
		// 	}

		// 	if amount > accountBalance {
		// 		fmt.Println("You don't have enough money.")
		// 		continue
		// 	}

		// 	accountBalance -= amount
		// 	fmt.Println("Your balance now is", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }
	}
}

func changeBalance(operationType string) (float64, error) {
	var amount float64
	fmt.Printf("What amount would you like to %v? ", operationType)
	fmt.Scan(&amount)
	if amount < 1 {
		return 0, errors.New("Value is less than 1")
	}
	return amount, nil
}
