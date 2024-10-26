package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5
	var investmentAmount := 1000
	expectedReturnRate := 5.5
	years := 10.0

	futureValue := (investmentAmount) * math.Pow(1+expectedReturnRate/100, (years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
