package main

import "testing"

func TestDivisibleByFive(t *testing.T) {
	t.Run("Caso divisível por 5", func(t *testing.T) {
		resultado := DivisibleByFive(10)
		esperado := "10 é divisível por 5"

		if resultado != esperado {
			t.Errorf("Resultado '%v' é diferente do esperado '%v'", resultado, esperado)
		}
	})
	t.Run("Caso indivisível por 5", func(t *testing.T) {
		resultado := DivisibleByFive(11)
		esperado := "11 não é divisível por 5"

		if resultado != esperado {
			t.Errorf("Resultado '%v' é diferente do esperado '%v'", resultado, esperado)
		}
	})
}
