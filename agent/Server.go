package agent

import (
	"log"
	"net/http"
	"strconv"
)

func Server(port int) {
	publisher := make(chan Job)
	reporter := make(chan Result)
	checker := make(chan Job)
	checked := make(chan string)

	go func() {
		JobWorker(publisher, reporter, checker, checked)
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
	http.HandleFunc("/job/status", JobStatus(checker, checked))
	http.HandleFunc("/job", JobSubscriber(publisher))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
