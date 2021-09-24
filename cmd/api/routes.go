package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovie)

	router.HandlerFunc(http.MethodGet, "/v1/movies/:genre_id", app.getAllMovie)

	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)

	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)

	return app.enableCORS(router)
}
