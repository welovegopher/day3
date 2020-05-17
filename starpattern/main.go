package main

import "fmt"

//Given a level, generate triangle of star
func main() {

	for i := 1; i <= 8; i++ {
		for j := 1; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
