package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/Edilberto-Vazquez/golang-rest-and-websockets/database"
	"github.com/Edilberto-Vazquez/golang-rest-and-websockets/repository"
	"github.com/rs/cors"

	websocket "github.com/Edilberto-Vazquez/golang-rest-and-websockets/websocket"
	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DataBaseUrl string
}

type Server interface {
	Config() *Config
	Hub() *websocket.Hub
}

type Broker struct {
	config *Config
	router *mux.Router
	hub    *websocket.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("jwt secret is required")
	}
	if config.DataBaseUrl == "" {
		return nil, errors.New("database url is required")
	}
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub:    websocket.NewHub(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	handler := cors.AllowAll().Handler(b.router)
	repo, err := database.NewPostgresRepository(b.config.DataBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	go b.hub.Run()
	repository.SetRepository(repo)
	log.Println("starting server on port", b.config.Port)
	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Println("error starting server:", err)
	} else {
		log.Fatalf("server stopped")
	}
}
