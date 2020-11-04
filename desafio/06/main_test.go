package main

import "testing"

func TestSUM(t *testing.T) {
	t.Run("Case numbers are all equal", func(t *testing.T) {
		resultado := Sum([]int{10, 10, 10})
		esperado := 30

		if resultado != esperado {
			t.Errorf("Resultado '%v' é diferente do esperado '%v'", resultado, esperado)
		}
	})

	t.Run("Case two numbers are equal", func(t *testing.T) {
		resultado := Sum([]int{9, 10, 9, 10, 9, 10})
		esperado := 57

		if resultado != esperado {
			t.Errorf("Resultado '%v' é diferente do esperado '%v'", resultado, esperado)
		}
	})

	t.Run("Case numbers are all diferent", func(t *testing.T) {
		resultado := Sum([]int{10, 23, 43, 9, 8, 45})
		esperado := 138

		if resultado != esperado {
			t.Errorf("Resultado '%v' é diferente do esperado '%v'", resultado, esperado)
		}
	})
}
