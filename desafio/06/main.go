package main

import "fmt"

func Sum(p []int) int {
	summed := 0
	for _, num := range p {
		summed += num
	}
	return summed
}


func main() {
	fmt.Println(Sum([]int{10,32,10}))
}