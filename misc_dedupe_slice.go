package main

import (
	"fmt"
	"sort"
)

func dedupe(input *[]int) []int {

	result := make([]int,0)
	seen := make(map[int]int, 0)

	for _, v := range *input {
		if _, ok := seen[v] ; !ok {
			result = append(result, v)
			seen[v] = v
		}
	}

	sort.Ints(result)

	return result
}

func main() {
	data := []int{1,2,5,4,4,2,6,7,3}
	// {1,2,3,4,5,6,7}

	clean := dedupe(&data)
	fmt.Printf("%#v\n", clean)

	fmt.Printf("%#v\n", data)
}