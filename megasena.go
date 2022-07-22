package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexsetta/httputils"
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

const (
	url = "https://servicebus2.caixa.gov.br/portaldeloterias/api/megasena"
)

var ultimo = ""

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

func main() {
	s := gocron.NewScheduler(time.UTC)
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatalf("Error loading location: %s", err)
	}
	s.ChangeLocation(location)

	job, err := s.Every(1).Day().Wednesday().Saturday().At("20:15;20:30").Do(task)
	//job, err := s.Every(10).Seconds().Do(taskTeste)
	if err != nil {
		log.Fatalf("Error creating job: %v", err)
	}
	go func() {
		for {
			time.Sleep(10 * time.Second)
			log.Println("Next run", job.NextRun())
			time.Sleep(12 * time.Hour)
		}
	}()
	s.StartBlocking()
}
