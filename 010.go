package main

import (
	"fmt"
	"math"
	"time"
)

func isPrime(n int64) bool {

	if n == 2 {
		return true
	} else if n < 2 {
		return false
	}

	var i int64

	sqr := int64(math.Sqrt(float64(n)))

	for i = 2; i <= sqr; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	now := time.Now()
	var psum, i int64 = 0, 0
	for i = 2; i < 2e6; i++ {
		if isPrime(i) {
			psum += i
		}
	}
	fmt.Printf("%s\n", time.Since(now))
	fmt.Printf("%d\n", psum)

}
