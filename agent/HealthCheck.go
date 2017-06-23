package agent

import (
	"net/http"

	"github.com/vmihailenco/msgpack"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// No case to return error for OK
	b, _ := msgpack.Marshal(&Result{Status: "OK"})
	w.WriteHeader(200)
	w.Write(b)
}
