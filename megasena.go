package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexsetta/httputils"
	"log"
	"time"
)

const (
	url   = "https://servicebus2.caixa.gov.br/portaldeloterias/api/megasena"
	sleep = 60 * time.Minute
)

var (
	ultimo = ""
)

func main() {
	for {
		data, err := httputils.GetBody(url, false)
		if err != nil {
			log.Fatal(err)
		}

		var res map[string]interface{}
		json.Unmarshal([]byte(data), &res)

		concurso := fmt.Sprintf("%v", res["numero"])
		if ultimo != concurso {
			ultimo = concurso
			notifica(concurso, fmt.Sprintf(" %v", res["listaDezenas"]))
		}

		log.Println("Aguardando próxima execução")
		time.Sleep(sleep)
	}
}
