package agent

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK!")
}
