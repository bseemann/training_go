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
	fmt.Sprintf("%v", bigger(12))
	fmt.Sprintf("%v", 12)
	fmt.Sprintf("%v", 12.0)
}
