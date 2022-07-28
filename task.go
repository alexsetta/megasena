package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexsetta/httputils"
	"log"
)

func taskTest() {
	log.Printf("Debug: %t, Último: %s", debug, ultimo)
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
	listaDezenas := fmt.Sprintf("%v", res["listaDezenas"])

	if debug {
		log.Printf("Concurso: %s, Último: %s, Dezenas sorteadas: %s", concurso, ultimo, listaDezenas)
	}
	if ultimo != concurso {
		ultimo = concurso
		msg := fmt.Sprintf("Mega-Sena concurso %s\n\r", concurso)
		msg += fmt.Sprintf("Resultado: %s\n\r", listaDezenas)
		msg += fmt.Sprintf("Aposta: %s\n\r", aposta)
		msg += "\n\r"

		acertos := verifica(aposta, listaDezenas)
		if len(acertos) >= 12 {
			msg += "*** PARABENS ***\n\r"
		}
		msg += fmt.Sprintf("Acertos (%d): %s", len(acertos)/3, acertos)
		notifica(msg)
	}
	return fmt.Sprintf("%v", res["listaDezenas"]), nil
}
