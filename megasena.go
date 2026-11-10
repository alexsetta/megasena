package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

var (
	apostas    []string
	debug      = false
	horarios   = "21:30"
	ultimo     = "0"
	concLimite = "0"
	url        = ""
)

func main() {
	fmt.Println("Iniciando")
	readConfig()

	s := gocron.NewScheduler(time.UTC)
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatalf("Error loading location: %s", err)
	}
	s.ChangeLocation(location)

	_, _ = task()

	job, err := s.Every(1).Week().Tuesday().Thursday().Saturday().At(horarios).Do(task)

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
