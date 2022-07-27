package main

import (
	"fmt"
	"strings"
)

func verifica(aposta, resultado string) string {
	resultado = resultado[1 : len(resultado)-1]
	res := ""

	a := strings.Split(aposta, " ")
	r := strings.Split(resultado, " ")

	for _, e := range a {
		if contains(r, e) {
			res += fmt.Sprintf("%s ", e)
		}
	}

	return res
}
