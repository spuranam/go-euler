package main

import (
	"fmt"
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)

	i := 0
	j := len(runes) - 1

	for ; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func isPalindrome(p string) bool {
	return p == reverse(p)
}

func main() {
start:
	for i := 999; i > 800; i-- {
		for j := 999; j > 800; j-- {
			n := i * j
			if isPalindrome(strconv.Itoa(n)) {
				fmt.Printf("Greatest product to largest three disgit (%d, %d) is %d\n", i, j, n)
				break start
			}
		}
	}

}
