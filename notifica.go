package main

import (
	"fmt"
	"github.com/alexsetta/telegram"
	"log"
)

func notifica(concurso, listaDezenas string) {
	config, err := telegram.ReadConfig("telegram.json")
	if err != nil {
		log.Println(err)
	}

	msg := fmt.Sprintf("Resultados Mega-Sena (%s): \n\r%s", concurso, listaDezenas)
	log.Println(msg)

	if err := telegram.SendMessage(config.ID, config.Token, msg); err != nil {
		log.Println(err)
	}
}
