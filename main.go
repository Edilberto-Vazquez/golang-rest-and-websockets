package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Edilberto-Vazquez/golang-rest-and-websockets/handlers"
	"github.com/Edilberto-Vazquez/golang-rest-and-websockets/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DataBaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)

}
