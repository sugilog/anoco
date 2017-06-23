package agent

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/vmihailenco/msgpack"
)

func JobSubscriber(worker chan Job) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			MethodNotAllowed(w)
			return
		}

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			ErrorResult(w, err)
			return
		}

		var params JobParams
		err = msgpack.Unmarshal(body, &params)

		if err != nil {
			ErrorResult(w, err)
			return
		}

		if len(params.Command) == 0 {
			ErrorResult(w, errors.New("Required 'Command' parameter not given."))
			return
		}

		job := Job{
			// FIXME
			Id: 0,
			// FIXME
			Command: params.Command,
		}
		worker <- job
		SuccessResult(w, Result{Id: job.Id, Status: "STARTED"})
	}
}
