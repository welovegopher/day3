package main

import (
	"fmt"
)

//Given a list of numbers, search a given number
func main(){

	var incoming = []int{1,5, 6, 12, 6}
	var numberToBeSearch  = 6

	for index, el := range incoming {

		if el == numberToBeSearch {
			fmt.Println("index ", index)
		}
	}
}
