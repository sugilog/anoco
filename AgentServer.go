package anoco

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
)

func AgentServer(port int) {
	http.HandleFunc("/", IndexHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
