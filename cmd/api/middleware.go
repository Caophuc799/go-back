package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pascaldekloe/jwt"
)

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		next.ServeHTTP(w, r)
	})
}

func (app *application) checkToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {

		}

		headerPars := strings.Split(authHeader, " ")

		if len(headerPars) != 2 {
			app.errorJSON(w, errors.New("Invalid auth header"))
			return
		}

		if headerPars[0] != "Bearer" {
			app.errorJSON(w, errors.New("UnAuthorized - no bearer"))
			return
		}

		token := headerPars[1]

		claims, err := jwt.HMACCheck([]byte(token), []byte(app.config.jwt.secret))
		if err != nil {
			app.errorJSON(w, errors.New("UnAuthorized - failed hmac check"))
			return
		}

		if !claims.Valid(time.Now()) {
			app.errorJSON(w, errors.New("UnAuthorized - token expired"))
			return
		}

		if !claims.AcceptAudience("mydomain.com") {
			app.errorJSON(w, errors.New("UnAuthorized - invalid audience"))
			return
		}

		if claims.Issuer != "mydomain.com" {
			app.errorJSON(w, errors.New("UnAuthorized - invalid issuer"))
			return
		}

		userID, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			app.errorJSON(w, errors.New("UnAuthorized"))
			return
		}

		println("User id ", userID)

		next.ServeHTTP(w, r)
	})
}
