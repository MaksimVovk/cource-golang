package main

import "fmt"

func main() {
	var revenue int
	var expenses int
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)

	ebt := float64(revenue - expenses)
	profit := ebt * (1 - taxRate/100)
	ration := ebt / profit

	fmt.Println("EBT: ", ebt)
	fmt.Println("Ration: ", ration)
	fmt.Println("Profit: ", profit)
}
