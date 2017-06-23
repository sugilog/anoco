package agent

import (
	"net/http"
)

func JobStatus(checker chan Job, checked chan string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// FIXME
		jobid := 1
		checker <- Job{Id: jobid}

		for {
			select {
			case status := <-checked:
				// FIXME
				SuccessResult(w, Result{Id: jobid, Status: status})
				return
			}
		}
	}
}
