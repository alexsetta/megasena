package main

import (
	"log"
	"strings"

	"github.com/tidwall/gjson"
)

const (
	ErrURLNotDefined  = "'URL' not defined in config file"
	ErrReadConfigFile = "Error read config file"
)

func readConfig() {
	json, err := ReadJson("megasena.json")
	if err != nil {
		log.Fatalf(ErrReadConfigFile)
	}
	url = gjson.Get(json, "url").String()
	if url == "" {
		log.Fatalf(ErrURLNotDefined)
	}
	apostas = parseApostas(gjson.Get(json, "aposta"))
	debug = gjson.Get(json, "debug").Bool()
	if h := gjson.Get(json, "horarios").String(); h != "" {
		horarios = h
	}
	concLimite = gjson.Get(json, "concurso").String()
}

// parseApostas aceita "aposta" como string única ou array de strings (um cartão por item).
func parseApostas(v gjson.Result) []string {
	itens := []string{v.String()}
	if v.IsArray() {
		itens = nil
		for _, e := range v.Array() {
			itens = append(itens, e.String())
		}
	}

	var res []string
	for _, a := range itens {
		if a = strings.TrimSpace(a); a != "" {
			res = append(res, a)
		}
	}
	return res
}
