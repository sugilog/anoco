package agent

import (
	"net/http"
)

func JobSubscriber(worker chan Job) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			MethodNotAllowed(w)
			return
		}

		job := Job{
			// FIXME
			Id: 0,
			// FIXME
			Command: "sleep 2; echo " + r.Method,
		}
		worker <- job
		SuccessResult(w, Result{Id: job.Id, Status: "STARTED"})
	}
}
