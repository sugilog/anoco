package agent

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	SuccessResult(w, Result{Status: "OK"})
}
