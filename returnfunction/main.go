package main

import (
	"fmt"
	"math"
)

//take a function that takes integer, if float is less than 50 return sin otherwise return cos
func main() {

	var (
		num             = 51
		outputFunc      func(float64) float64
		floatTypeNumber = float64(num)
	)

	outputFunc = getFunction(num)
	output := outputFunc(floatTypeNumber)

	fmt.Println(output)
}

func getFunction(a int) func(float64) float64 {
	if a < 50 {
		return math.Sin
	} else {
		return math.Cos
	}
}
