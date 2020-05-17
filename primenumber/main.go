package main

import "fmt"

//given a number, check if it is a prime number or not.
func main()  {

	var numberToBeChecked = 121
	var isPrimeNumber = true
	var half = numberToBeChecked / 2

	for i := 2; i < half; i++ {
		if numberToBeChecked % i == 0 {
			fmt.Println("divisible by ", i)
			isPrimeNumber = false
			break
		}
	}

	if isPrimeNumber {
		fmt.Println("is a prime number")
	} else {
		fmt.Println("is not a prime number")
	}

}
