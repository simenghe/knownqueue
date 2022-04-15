package main

import (
	"encoding/json"
	"fullservice/db"
	"fullservice/middlewares"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
)

const (
	PORT = ":8080"
)

type CompositeKey struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Server struct {
	Mux       chi.Mux
	Cache     db.CacheStore
	Mutex     sync.RWMutex
	mutexOnce *sync.Once
}

var cache = db.New()

func NewServer() *Server {
	return &Server{
		Mux:   *chi.NewRouter(),
		Cache: db.New(),
	}
}

func main() {
	server := NewServer()
	// Middleware Initialization
	server.Mux.Use(middleware.RequestID)
	server.Mux.Use(middleware.RealIP)
	server.Mux.Use(middleware.Logger)
	server.Mux.Use(middleware.Recoverer)
	server.Mux.Use(middlewares.Json)

	// Connect Cache + Databases
	defer cache.Close()
	val, err := cache.Get("Hello")
	log.Println(val, err == redis.Nil)

	// Route Definitions
	server.Mux.Get("/", HomePageHandler)
	server.Mux.Get("/setex", SetExpiryHandler)
	server.Mux.Get("/get", GetFromCacheHandler)
	server.Mux.Mount("/admin", AdminRouter())

	log.Printf("Serving http server on port %s", PORT)
	err = http.ListenAndServe(PORT, server)
	if err != nil {
		log.Fatalln(err)
	}

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(CompositeKey{1, 2})
}

func SetExpiryHandler(w http.ResponseWriter, r *http.Request) {
	err := cache.SetEx("test", CompositeKey{1, 3}, 30*time.Second)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(CompositeKey{1, 2})
}

func GetFromCacheHandler(w http.ResponseWriter, r *http.Request) {
	val, err := cache.Get("test")

	if err == redis.Nil {
		log.Println(err)
		w.WriteHeader(http.StatusPartialContent)
		return
	}

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(val)
}
