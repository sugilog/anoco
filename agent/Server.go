package agent

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
)

func Server(port int) {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/health", HealthCheck)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
