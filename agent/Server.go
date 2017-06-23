package agent

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
)

func Server(port int) {
	worker := make(chan Job)
	reporter := make(chan Result)

	go func() {
		JobWorker(worker, reporter)
	}()

	// FIXME
	go func() {
		for {
			select {
			case result := <-reporter:
				log.Print(result)
			}
		}
	}()

	http.HandleFunc("/health", HealthCheck)
	http.HandleFunc("/job", JobSubscriber(worker))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
