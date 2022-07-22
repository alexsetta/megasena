package main

import (
	"github.com/go-co-op/gocron"
	"github.com/tidwall/gjson"
	"log"
	"time"
)

const (
	ErrURLNotDefined  = "'URL' not defined in config file"
	ErrReadConfigFile = "Error read config file"
)

var (
	ultimo = ""
	url    = ""
)

func main() {
	json, err := ReadJson("megasena.json")
	if err != nil {
		log.Fatalf(ErrReadConfigFile)
	}
	value := gjson.Get(json, "url")
	url = value.String()
	if url == "" {
		log.Fatalf(ErrURLNotDefined)
	}

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
