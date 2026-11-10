package main

import (
	"testing"
)

func Test_Verifica(t *testing.T) {
	retornado := verifica("04 10 23 30 49 55", "[16 20 21 39 44 55]")
	esperado := "55 "

	if esperado != retornado {
		t.Errorf("expected: %s, returned: %s", esperado, retornado)
	}
}

func Test_formataApostas(t *testing.T) {
	retornado := formataApostas([]string{"04 10 23 30 49 55", "16 20 21 39 44 01"}, "[16 20 21 39 44 55]")
	esperado := "\n\rCartão 1: 04 10 23 30 49 55\n\rAcertos (1): 55\n\r" +
		"\n\rCartão 2: 16 20 21 39 44 01\n\rAcertos (5): 16 20 21 39 44 <-- QUINA\n\r"

	if esperado != retornado {
		t.Errorf("expected: %q, returned: %q", esperado, retornado)
	}
}
