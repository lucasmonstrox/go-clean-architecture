package domain

import "testing"

func Soma(valores ...int) (total int) {
	for _, valor := range valores {
		total += valor
	}
	return
}

func TestSoma(t *testing.T) {
	test := []int{7, 2, 3, 4}
	total := Soma(test...)
	if total != 10 {
		t.Fatalf("Valor esperado: 10 - Valor retornado: %d", total)
	}
}
