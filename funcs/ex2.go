package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getInput("Revenue")
	if err != nil {
		fmt.Printf("Error while get revenue: %v\n", err)
		return
	}

	expenses, err := getInput("Expenses")
	if err != nil {
		fmt.Printf("Error while get expenses: %v\n", err)
		return
	}

	taxRate, err := getInput("Tax rate")
	if err != nil {
		fmt.Printf("Error while get tax rate: %v\n", err)
		return
	}

	ebt := getEbt(revenue, expenses)
	profit := getProfit(ebt, taxRate)
	ration := getRation(ebt, profit)
	writreData(ebt, profit, ration)

	fmt.Println("EBT:", ebt)
	fmt.Println("Ration:", ration)
	fmt.Println("Profit:", profit)
}

func writreData(ebt, ration, profit float64) {
	data := fmt.Sprintf("ebt: %v\nration: %v\nprofit: %v", ebt, ration, profit)
	os.WriteFile("calculatedData.txt", []byte(data), 0644)
}

func getInput(text string) (float64, error) {
	var val float64
	fmt.Printf("%v: ", text)
	fmt.Scan(&val)

	if val < 1 {
		return 0, errors.New("Invalid value.")
	}
	return val, nil
}

func getEbt(revenue float64, expenses float64) float64 {
	return revenue - expenses
}

func getProfit(ebt float64, taxRate float64) float64 {
	return ebt * (1 - taxRate/100)
}

func getRation(ebt, profit float64) float64 {
	return ebt / profit
}
