package main

import "testing"

func TestLoan(t *testing.T) {
	t.Run("Case loan is less than 30% of salary", func(t *testing.T) {
		resultado := Loan(3000, 800)
		esperado := "Empréstimo pode ser concedido"

		if resultado != esperado {
			t.Errorf("Resultado '%v' é diferente do esperado '%v'", resultado, esperado)
		}
	})

	t.Run("Case loan is exactly 30% of salary", func(t *testing.T) {
		resultado := Loan(3000, 900)
		esperado := "Empréstimo pode ser concedido"

		if resultado != esperado {
			t.Errorf("Resultado '%v' é diferente do esperado '%v'", resultado, esperado)
		}
	})

	t.Run("Case loan is greater than 30% of salary", func(t *testing.T) {
		resultado := Loan(3000, 950)
		esperado := "Empréstimo não pode ser concedido"

		if resultado != esperado {
			t.Errorf("Resultado '%v' é diferente do esperado '%v'", resultado, esperado)
		}
	})
}
