package main

import (
	"fmt"
	"time"
)

func sumSquares(n int) int {
	ssquares := 0
	for i := 1; i <= n; i++ {
		ssquares += i * i
	}
	return ssquares
}

func squaresSum(n int) int {
	sSum := 0
	for i := 0; i <= n; i++ {
		sSum += i
	}

	return sSum * sSum
}

func main() {
	n := 100

	now := time.Now()
	fmt.Println(squaresSum(n) - sumSquares(n))
	fmt.Printf("%s\n", time.Since(now))
}
