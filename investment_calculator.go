package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var investmentDuration float64 = 10

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100*float64(investmentDuration)), investmentDuration)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), investmentDuration)

	fmt.Println("Future Real Value of Investment:", futureRealValue)
	fmt.Println("Future Value of Investment:", futureValue)

	var salarioBase float64 = 1500

	fmt.Printf("Ingresar el Salario Base: ")
	fmt.Scan(&salarioBase)

	fmt.Printf("El Salario Base es: %.2f\n", salarioBase)
	salarioMenosImpuestos := salarioBase - (salarioBase * 0.30)
	fmt.Printf("El Salario Base menos impuestos es: %.2f\n", salarioMenosImpuestos)
}
