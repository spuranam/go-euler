package main

import (
	"fmt"
)

func fib(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)

}

func main() {

	sum := 0
	for i := 0; i < 4e6; i++ {
		c := fib(i)
		if c%2 == 0 {
			sum += c
		}

		if c > 4e6 {
			break
		}
	}

	fmt.Println(sum)

}
