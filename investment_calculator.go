package main

func main() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var investmentDuration = 10

	var futureValue = float64(investmentAmount) * (1 + expectedReturnRate/100*float64(investmentDuration))
}
