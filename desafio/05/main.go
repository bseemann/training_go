package main

func MyFunc(p []int) int {
	bigger := p[0]
	for _, num := range p {
		if num > bigger {
			bigger = num
		}
	}
	return bigger
}
