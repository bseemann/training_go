package main

import (
	"fmt"
)

func EvenOrOdd(n int) string {
	// case integer n is odd, the mod function by 2 returns a number always bigger than 0
	if (n % 2) == 0 {
		return fmt.Sprintf("%v é par", n)
	}
	return fmt.Sprintf("%v é ímpar", n)
}

func main() {
	fmt.Println(EvenOrOdd(10))
	fmt.Println(EvenOrOdd(5))
}
