package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexsetta/httputils"
	"log"
)

func taskTest() {
	log.Printf("Debug: %s, Último: %s", debug, ultimo)
}

func task() (string, error) {
	data, err := httputils.GetBody(url, false)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var res map[string]interface{}
	json.Unmarshal([]byte(data), &res)

	concurso := fmt.Sprintf("%v", res["numero"])

	if debug {
		log.Printf("Concurso: %s, Último: %s", concurso, ultimo)
	}
	if ultimo != concurso {
		ultimo = concurso
		notifica(concurso, fmt.Sprintf(" %v", res["listaDezenas"]))
	}
	return fmt.Sprintf("%v", res["listaDezenas"]), nil
}
