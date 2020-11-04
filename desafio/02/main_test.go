package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	t.Run("Caso par", func(t *testing.T) {
		resultado := EvenOrOdd(4)
		esperado := "4 é par"

		if resultado != esperado {
			t.Errorf("Resultado '%v' é diferente do esperado '%v'", resultado, esperado)
		}
	})

	t.Run("Caso ímpar", func(t *testing.T) {
		resultado := EvenOrOdd(7)
		esperado := "7 é ímpar"

		if resultado != esperado {
			t.Errorf("Resultado '%v' é diferente do esperado '%v'", resultado, esperado)
		}
	})
}
