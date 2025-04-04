package main

import (
	"log"

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
	aposta = gjson.Get(json, "aposta").String()
	debug = gjson.Get(json, "debug").Bool()
	horarios = gjson.Get(json, "horarios").String()
	concLimite = gjson.Get(json, "concurso").String()
}
