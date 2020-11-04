package main

import "fmt"

func bigger(n float64) float64 {
	if n > 10 {
		return 2 * n
	} else if n == 10 {
		return n
	}
	return n / 2
}

func main() {
	fmt.Printf("%v -> %v\n", 12, bigger(12))
	fmt.Printf("%v -> %v\n", 10, bigger(10))
	fmt.Printf("%v -> %v\n", 9, bigger(9))
}
