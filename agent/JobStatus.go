package agent

import (
	"net/http"
)

func JobStatus(checker chan Job, checked chan string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// FIXME
		jobid := "1"
		checker <- Job{UUID: jobid}

		for {
			select {
			case status := <-checked:
				SuccessResult(w, Result{UUID: jobid, Status: status})
				return
			}
		}
	}
}
