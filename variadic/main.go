package main

import "fmt"

func main() {
	variadic("a", "b", "c", "d", "c", "e")
	variadic("a", "b", "c", "d")
	variadicSum(1,2,4,4,5,6)
}

func variadic(a... string)  {
	for _, el := range a {
		fmt.Print(el)
	}
}

func variadicSum(b... int)  {

	var sum = 0
	for _, el := range b {
		sum = sum + el
	}

	fmt.Println(sum)
}
