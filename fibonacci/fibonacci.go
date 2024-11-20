package main

import (	
	"fmt"
)

func main() {

	var num int = 5

	// Display the sorted slice of integers.
	r := fib(num)
	fmt.Println(r)
}

func fib(num int) (int) {
	if num <= 0 {
		return 1
	}
	return fib(num-1) + fib(num-2)
}

