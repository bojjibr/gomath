package calc

import "testing"

func TestSoma(t *testing.T) {
	t.Run("Soma sem dígito", func(t *testing.T) {
		got := Sum(3, 2)
		if got != 5 {
			t.Errorf("Soma(3, 2) = %v; want 5", got)
		}
	})

	t.Run("Soma com dígito", func(t *testing.T) {
		got := Sum(2.5, 2.5)
		if got != 5 {
			t.Errorf("Soma(2.5, 2.5) = %v; want 5", got)
		}
	})
}

func TestSub(t *testing.T) {
	t.Run("Subtração sem dígito", func(t *testing.T) {
		got := Sub(3, 2)
		if got != 1 {
			t.Errorf("Subtração(3, 2) = %v; want 1", got)
		}
	})

	t.Run("Subtração com dígito", func(t *testing.T) {
		got := Sub(5.7, 2.5)
		if got != 3.2 {
			t.Errorf("Subtração(5.7, 2.5) = %v; want 3.2", got)
		}
	})
}

func TestMult(t *testing.T) {
	t.Run("Multiplicação sem dígito", func(t *testing.T) {
		got := Mult(3, 2)
		if got != 6 {
			t.Errorf("Multiplicação(3, 2) = %v; want 6", got)
		}
	})

	t.Run("Multiplicação com dígito", func(t *testing.T) {
		got := Mult(5.7, 2.5)
		if got != 14.25 {
			t.Errorf("Multiplicação(5.7, 2.5) = %v; want 14.25", got)
		}
	})
}
