package main

import (
	"encoding/json"
	"fullservice/middlewares"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

const (
	PORT = ":5000"
)

type CompositeKey struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func Square(n int) int {
	return n * n
}

func main() {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middlewares.Json)

	r.Get("/", HomePage)
	r.Mount("/admin", AdminRouter())

	log.Printf("Serving http server on port %s", PORT)
	http.ListenAndServe(PORT, r)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := CompositeKey{1, 2}
	json.NewEncoder(w).Encode(key)
}
