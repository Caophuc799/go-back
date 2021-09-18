package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovie)

	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.getOneMovie)

	return router
}
