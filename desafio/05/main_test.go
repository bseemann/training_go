package main

import "testing"

func TestMyFunc(t *testing.T) {
	t.Run("Case numbers are all equal", func(t *testing.T) {
		resultado := MyFunc([]int{10, 10, 10})
		esperado := 10

		if resultado != esperado {
			t.Errorf("Resultado '%v' é diferente do esperado '%v'", resultado, esperado)
		}
	})

	t.Run("Case two numbers are equal", func(t *testing.T) {
		resultado := MyFunc([]int{9, 10, 9})
		esperado := 10

		if resultado != esperado {
			t.Errorf("Resultado '%v' é diferente do esperado '%v'", resultado, esperado)
		}
	})

	t.Run("Case numbers are all diferent", func(t *testing.T) {
		resultado := MyFunc([]int{4, 6, 2, 8, 1})
		esperado := 8

		if resultado != esperado {
			t.Errorf("Resultado '%v' é diferente do esperado '%v'", resultado, esperado)
		}
	})
}
