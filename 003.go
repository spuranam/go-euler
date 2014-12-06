package main

import (
	"fmt"
	"math"
)

func isPrime(n int64) bool {

	if n == 2 {
		return true
	} else if n < 2 {
		return false
	}

	var i int64

	//sqr := int(math.Ceil(math.Sqrt(float64(n))))

	//for i:=2; i<sqr; i++ {
	for i=2; i<n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var number int64
	number = 600851475143

	i := int64(math.Sqrt(float64(number)))

	for ; i > 1 ; i-- {
		if number % i == 0 && isPrime(i) {
			fmt.Printf("%d\n", i)
			break
		}
	}
}
