package main

import "strconv"

func EvenOrOdd(n int) string {
	s := strconv.Itoa(n)
	// case integer n is odd, the mod function by 2 returns a number always bigger than 0
	if (n % 2) == 0 {
		s += " é par"
	} else {
		s += " é ímpar"
	}
	return s
}
