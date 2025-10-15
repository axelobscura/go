package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var investmentDuration float64 = 10

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100*float64(investmentDuration)), investmentDuration)
	fmt.Println("Future Value of Investment:", futureValue)
}
