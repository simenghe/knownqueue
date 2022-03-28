package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AdminRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.BasicAuth("BasicAuth", map[string]string{
		"Admin": "Password",
	}))
	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Admin Router Reached"))
	})

	router.Get("/panel", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Panel Reached"))
	})
	return router
}
