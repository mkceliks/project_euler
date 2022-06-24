package main

import (
	"fmt"
	"strconv"
)

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func main() {
	var temp int = 0
	for i := 0; i <= 33; i++ {
		fmt.Print(strconv.Itoa(FibonacciRecursion(i)) + " ")
		if FibonacciRecursion(i)%2 == 0 {
			temp = temp + FibonacciRecursion(i)
		}
	}
	fmt.Println("")
	fmt.Println(temp)
}
