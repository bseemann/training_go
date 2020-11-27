package main

import "fmt"

func MyFunc(p []int) int {
	bigger := p[0]
	for _, num := range p {
		if num > bigger {
			bigger = num
		}
	}
	return bigger
}

func main() {
	fmt.Println(MyFunc([]int{10,340,43,233,4332}))
	fmt.Println(MyFunc([]int{10,24}))
	fmt.Println(MyFunc([]int{34,34}))
	fmt.Println(MyFunc([]int{44}))
}
