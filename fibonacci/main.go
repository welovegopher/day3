package main

import (
	"fmt"
)

//Write a program to generate fibonacci series
// 1, 2, 3, 5, 8, 13 ...

func main(){

	var a = 1
	var b = 1
	var sum = 0

	for i := 0; i < 10; i++ {
		sum = a+b
		a = b
		b = sum
		fmt.Println(a)
	}
}

//returns multiple variables
func sumMultipleReturns(a int, b int, limit int)  (int, bool) {
	sum :=  SumFn4(a, b)
	if sum > limit {
		return 0, true
	}
	return sum, false
}

//simple function
func sumFn1(a int, b int)  int {
	return  a+b
}

//parameters with same type
func sumFn2(a, b int)  int {
	return  a+b
}

//parameters with named-value
func sumFn3(a, b int)  (sum int) {
	sum =  a+b
	return sum
}

//parameters with named-value
func SumFn4(a... int)  (sum int) {
	 for _, element := range a {
	 	sum = sum + element
	 }
	 return sum
}

//function with function as parameter
func sumFn5(a int, b int, sumFunc func(a, b int) int)  (sum int) {
	return sumFunc(a, b)
}


