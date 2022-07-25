package main

import (
	"github.com/tidwall/gjson"
	"log"
	"testing"
)

func Test_Verifica(t *testing.T) {
	aposta := "04 10 23 30 49 55"
	resultado := "[16 20 21 39 44 55]"
	esperado := "55"
	retornado := ""

	if aposta != resultado {
		t.Errorf("expected: %s, returned: %s", esperado, retornado)
	}
}

func Test_task(t *testing.T) {
	json, err := ReadJson("megasena.json")
	if err != nil {
		log.Fatalf(ErrReadConfigFile)
	}
	value := gjson.Get(json, "url")
	url = value.String()
	if url == "" {
		log.Fatalf(ErrURLNotDefined)
	}

	resultado, err := task()
	if err != nil {
		t.Errorf("erro: %v", err)
	}

	println(resultado)

	if len(resultado) == 0 {
		t.Errorf("resultado não encontrado")
	}
}
