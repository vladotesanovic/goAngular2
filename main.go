package main

import (
	"encoding/json"
	"net/http"

	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"github.com/vladotesanovic/goangular2/lib"
	"github.com/vladotesanovic/goangular2/routes"
)

func main() {
	router := httprouter.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		v := make(map[string]string)
		v["page"] = "Welcome to Go Angular 2"
		data, _ := json.Marshal(v)

		w.Write(data)

	})

	// colllection of routes
	router.GET("/index", routes.IndexRoutes)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(lib.JSONMiddleware))
	router.NotFound = http.HandlerFunc(lib.NotFound)

	// serve client files
	router.ServeFiles("/static/*filepath", http.Dir("./client"))

	n.UseHandler(router)

	http.ListenAndServe(":"+port, n)
}
