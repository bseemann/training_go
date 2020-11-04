package main

func Sum(p []int) int {
	summed := 0
	for _, num := range p {
		summed += num
	}
	return summed
}
