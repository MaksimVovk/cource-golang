package main

import (
	"fmt"
	"math"
)

const investmentAmount = 1000.0

func main() {
	val, val2 := calculateFutureValues(10.0, 10.0)
	outPutText("Test\n")
	fmt.Println(val)
	fmt.Println(val2)
}

func outPutText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(expectedReturnRate, years float64) (float64, float64) {
	val1 := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	val2 := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	return val1, val2
}
