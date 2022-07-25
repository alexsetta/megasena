package main

import (
	"github.com/alexsetta/telegram"
	"log"
)

func notifica(msg string) {
	config, err := telegram.ReadConfig("telegram.json")
	if err != nil {
		log.Println(err)
	}

	log.Println(msg)
	if err := telegram.SendMessage(config.ID, config.Token, msg); err != nil {
		log.Println(err)
	}
}
