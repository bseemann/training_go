package main

import "fmt"

const LoanPassed = "Empréstimo pode ser concedido"
const LoanNotPassed = "Empréstimo não pode ser concedido"

func Loan(salary float64, credit float64) string {
	maxLoan := salary * 0.3
	if credit <= maxLoan {
		return LoanPassed
	}
	return LoanNotPassed
}


func main() {
	fmt.Println(Loan(3000,800))
	fmt.Println(Loan(3000,950))
}