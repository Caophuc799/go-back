package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()

	secure := alice.New(app.checkToken)

	router.HandlerFunc(http.MethodPost, "/v1/graphql/list", app.moviesGraphQL)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/v1/signin", app.Signin)

	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.getOneMovie)

	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/v1/genres/:genre_id/movies", app.getAllMovie)

	router.PUT("/v1/admin/movies/:id", app.wrap(secure.ThenFunc(app.updateMovie)))
	router.DELETE("/v1/admin/movies/:id", app.wrap(secure.ThenFunc(app.deleteMovie)))
	router.POST("/v1/admin/movies", app.wrap(secure.ThenFunc(app.insertMovie)))
	// router.HandlerFunc(http.MethodPut, "/v1/admin/movies/:id", app.wrap(secure.ThenFunc(app.updateMovie))
	// router.HandlerFunc(http.MethodDelete, "/v1/admin/movies/:id", app.deleteMovie)
	// router.HandlerFunc(http.MethodPost, "/v1/admin/movies", app.insertMovie)

	return app.enableCORS(router)
}
