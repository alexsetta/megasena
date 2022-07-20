package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexsetta/httputils"
	"github.com/alexsetta/telegram"
	"log"
	"time"
)

const (
	url   = "https://servicebus2.caixa.gov.br/portaldeloterias/api/megasena"
	sleep = 60 * time.Minute
)

func main() {
	for {
		data, err := httputils.GetBody(url, false)
		if err != nil {
			log.Fatal(err)
		}

		var res map[string]interface{}
		json.Unmarshal([]byte(data), &res)

		log.Fatal(res["dataApuracao"])

		config, err := telegram.ReadConfig("telegram.json")
		if err != nil {
			log.Fatal(err)
		}

		if err := telegram.SendMessage(config.ID, config.Token, "Resultados Mega-Sena: \n\r"+fmt.Sprintf(" %v", res["listaDezenas"])); err != nil {
			log.Fatal(err)
		}

		log.Println("Aguardando próxima execução")
		time.Sleep(sleep)
	}
}
