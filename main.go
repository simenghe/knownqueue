package main

import (
	"encoding/json"
	"fullservice/db"
	"fullservice/middlewares"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
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

	// Middleware Initialization
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middlewares.Json)

	// Connect Cache + Databases
	cache := db.New()
	defer cache.Close()
	val, err := cache.Get("Hello")
	log.Println(val, err == redis.Nil)

	// Route Definitions
	r.Get("/", HomePage)
	r.Mount("/admin", AdminRouter())

	log.Printf("Serving http server on port %s", PORT)
	http.ListenAndServe(PORT, r)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(CompositeKey{1, 2})
}
