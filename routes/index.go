package routes

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// IndexRoutes is collection of routes
func IndexRoutes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	v := make(map[string]string)
	v["page"] = "Index route"
	data, _ := json.Marshal(v)

	w.Write(data)
}
