package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64
	years := 10.0
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(
		(1+expectedReturnRate/100),
		years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future Value: %v\n", futureValue)

	val := fmt.Sprintf("Future value %.1f\n", futureValue)
	fmt.Println("futureRealValue", futureRealValue)
	fmt.Printf(`
	Future Real Value: %v
	%v`, futureValue, val)
}
