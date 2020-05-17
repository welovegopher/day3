package main

import (
	"fmt"
)

//given number, reverse of the number should match the original number
func main() {
	var originalNumber = 121

	if originalNumber == reverseFunc(originalNumber) {
		fmt.Println("it is a palindrome")
	} else {
		fmt.Println("it is not a palindrome")
	}
}

func reverseFunc(originalNumber int) int {
	var reverse = 0
	for originalNumber > 0 {
		r := originalNumber % 10
		reverse = reverse*10 + r
		originalNumber = originalNumber / 10
	}
	return reverse
}
