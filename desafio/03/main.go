package main

import "strconv"

func DivisibleByFive(n int) string {
	s := strconv.Itoa(n)
	// case integer n is divisible by 5, the mod function by 5 returns a number always bigger than 0
	if (n % 5) == 0 {
		s += " é divisível por 5"
	} else {
		s += " não é divisível por 5"
	}
	return s
}
