package main

import "fmt"

//Given a number, generate the factorial of the number
func main() {

	var factorialNumber = 5
	var factorialValue = 1;

	for i := 1; i <= factorialNumber; i++ {
		factorialValue = factorialValue*i
	}

	fmt.Println(factorialValue)
}
