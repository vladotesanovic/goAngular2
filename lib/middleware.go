package lib

import (
	"encoding/json"
	"net/http"
)

// JSONMiddleware converts all request to be application/json
func JSONMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	rw.Header().Set("Content-Type", "application/json")

	next(rw, r)
}

// NotFound middleware function
func NotFound(w http.ResponseWriter, r *http.Request) {

	response := make(map[string]string)
	response["message"] = "Requested resources does not exists"

	v, _ := json.Marshal(response)

	w.Write(v)
}
