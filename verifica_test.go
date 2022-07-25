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
