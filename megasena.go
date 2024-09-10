package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

var (
	aposta   = ""
	debug    = false
	horarios = "21:00"
	ultimo   = "0"
	url      = ""
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

	//job, err := s.Every(1).Day().Wednesday().Saturday().At(horarios).Do(task)
	job, err := s.Every(1).Day().At(horarios).Do(task)
	if debug {
		//job, err = s.Every(10).Seconds().Do(taskTest)
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
