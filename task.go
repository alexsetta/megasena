package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexsetta/httputils"
	"log"
)

func taskTeste() {
	log.Println("ok")
}

func task() {
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
}
