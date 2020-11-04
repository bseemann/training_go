package main

import (
	"fmt"
)

func DivisibleByFive(n int) string {
	// case integer n is divisible by 5, the mod function by 5 returns a number always bigger than 0
	if (n % 5) == 0 {
		return  fmt.Sprintf("%v é divisível por 5", n)
	}
	return  fmt.Sprintf("%v não é divisível por 5", n)
}

func main() {
	fmt.Println(DivisibleByFive(5))
	fmt.Println(DivisibleByFive(25))
	fmt.Println(DivisibleByFive(10))
	fmt.Println(DivisibleByFive(18))
}
