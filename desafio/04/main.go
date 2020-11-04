package main

import "fmt"

func MyFunc(n, m, p int) int {
	bigger := n
	if m > bigger {
		bigger = m
	}
	if p > bigger {
		bigger = p
	}
	return bigger
}

func main() {
	fmt.Println(MyFunc(10,2,32))
	fmt.Println(MyFunc(10,50,32))
	fmt.Println(MyFunc(433,50,32))
}
