package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexsetta/httputils"
	"log"
	"strings"
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

	if concLimite != "0" && concurso >= concLimite {
		notifica(fmt.Sprintf("Refaça as apostas, concurso %s é maior ou igual ao limite %s", concurso, concLimite))
	}

	if ultimo != concurso {
		ultimo = concurso
		msg := fmt.Sprintf("Mega-Sena concurso %s\n\r", concurso)
		msg += fmt.Sprintf("Concurso limite: %s\n\r", concLimite)
		msg += fmt.Sprintf("Resultado: %s\n\r", listaDezenas)
		msg += formataApostas(apostas, listaDezenas)
		notifica(msg)
	}
	return fmt.Sprintf("%v", res["listaDezenas"]), nil
}

// formataApostas gera o trecho da mensagem com o resultado de cada cartão.
func formataApostas(apostas []string, listaDezenas string) string {
	msg := ""
	sena := false
	for i, aposta := range apostas {
		acertos := verifica(aposta, listaDezenas)
		qtd := len(strings.Fields(acertos))

		msg += "\n\r"
		msg += fmt.Sprintf("Cartão %d: %s\n\r", i+1, aposta)
		msg += fmt.Sprintf("Acertos (%d): %s", qtd, strings.TrimSpace(acertos))
		switch {
		case qtd == 6:
			sena = true
			msg += " <-- SENA"
		case qtd == 5:
			msg += " <-- QUINA"
		case qtd == 4:
			msg += " <-- QUADRA"
		}
		msg += "\n\r"
	}
	if sena {
		msg += "\n\r*** PARABENS ***\n\r"
	}
	return msg
}
