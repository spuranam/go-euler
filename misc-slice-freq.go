package main

import (
	"fmt"
)

func freq(input *[]int) map[int]int {
	result := make(map[int]int,0)

	for _, v := range *input {
		if _, ok := result[v]; ok {
			result[v] += 1
			continue 
		}
		result[v] = 1
	}

	return result
}


func main(){
	data := []int{1,3,5,5,7,12,2,3}

	f := freq(&data)

	fmt.Printf("%#v", f)
}
