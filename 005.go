package main

import (
	"fmt"
)

func isDivisible(i, min, max int) bool {

	for j := min; j <= max; j++ {
		if i%j != 0 {
			return false
		}
	}

	return true

}

func main() {

	min, max := 1, 20
	for i := 20; ; i++ {
		if isDivisible(i, min, max) {
			fmt.Printf("smallest positive is %d\n", i)
			break
		}
	}

}
