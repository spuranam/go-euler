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
	n := 10001
	c := 0
	var i int64 = 3
	for {
		if isPrime(i) {
			c++
			if c == n {
				fmt.Printf("The %dth prime is %d\n", n, i)
				break
			}
		}
		i++
	}
	fmt.Printf("%s\n", time.Since(now))

}
