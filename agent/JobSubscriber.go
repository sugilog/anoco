package agent

import (
	"net/http"

	"github.com/vmihailenco/msgpack"
)

func JobSubscriber(worker chan Job) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		job := Job{
			// FIXME
			Id: 0,
			// FIXME
			Command: "sleep 2; echo 12345",
		}
		worker <- job

		// No case to return error for OK
		b, _ := msgpack.Marshal(&Result{
			Id:     job.Id,
			Status: "STARTED",
		})
		w.WriteHeader(200)
		w.Write(b)
	}
}
