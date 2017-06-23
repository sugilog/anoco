package agent

import (
	"fmt"
	"net/http"

	"github.com/vmihailenco/msgpack"
)

func SuccessResult(w http.ResponseWriter, result Result) {
	b, err := msgpack.Marshal(&result)

	if err != nil {
		ErrorResult(w, result.Id, err)
	} else {
		w.WriteHeader(200)
		w.Write(b)
	}
}

func ErrorResult(w http.ResponseWriter, id int, err error) {
	b, _ := msgpack.Marshal(&Result{
		Id:    id,
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
