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
	debug  = false
	ultimo = "0"
	url    = ""
)

func main() {
	json, err := ReadJson("megasena.json")
	if err != nil {
		log.Fatalf(ErrReadConfigFile)
	}
	url = gjson.Get(json, "url").String()
	if url == "" {
		log.Fatalf(ErrURLNotDefined)
	}
	debug = gjson.Get(json, "debug").Bool()

	s := gocron.NewScheduler(time.UTC)
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatalf("Error loading location: %s", err)
	}
	s.ChangeLocation(location)

	job, err := s.Every(1).Day().Wednesday().Saturday().At("20:20;20:30").Do(task)
	if debug {
		job, err = s.Every(10).Seconds().Do(taskTest)
	}
	if err != nil {
		log.Fatalf("Error creating job: %v", err)
	}
	go func() {
		for {
			time.Sleep(5 * time.Second)
			log.Println("Next run: ", job.NextRun())
			time.Sleep(12 * time.Hour)
		}
	}()
	s.StartBlocking()
}
