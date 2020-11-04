package main

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
