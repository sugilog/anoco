package agent

import (
	"fmt"
	"net/http"

	"github.com/vmihailenco/msgpack"
)

func SuccessResult(w http.ResponseWriter, result Result) {
	b, err := msgpack.Marshal(&result)

	if err != nil {
		ErrorResult(w, err)
	} else {
		w.WriteHeader(200)
		w.Write(b)
	}
}

func ErrorResult(w http.ResponseWriter, err error) {
	b, _ := msgpack.Marshal(&Result{
		Error: fmt.Sprintf("%s", err),
	})

	w.WriteHeader(500)
	w.Write(b)
}

func MethodNotAllowed(w http.ResponseWriter) {
	b, _ := msgpack.Marshal(&Result{
		Error: "Method Not Allowed",
	})

	w.WriteHeader(405)
	w.Write(b)
}
