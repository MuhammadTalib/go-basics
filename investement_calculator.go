package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investemtAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	// investemtAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5 more consice

	futureValue := investemtAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
