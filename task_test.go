package main

import (
	"github.com/tidwall/gjson"
	"log"
	"testing"
)

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
