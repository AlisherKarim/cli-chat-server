package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alisherkarim/cli-chat-server/api/v1/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func handleJoin(responseWriter http.ResponseWriter, req *http.Request) {
	fmt.Print(req)
	data, _ := json.Marshal("Hello from server")
	responseWriter.Write(data)
}

func initRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	return router
}

func main()  {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env no loaded correctly")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not loaded")
	}

	indexRouter := initRouter()

	// router for paths /v1/*
	routes.RegisterRoutes(indexRouter)

	server := &http.Server{
		Addr: ":" + portString,
		Handler: indexRouter,
	}
	
	log.Printf("Started custom server on port=%s", portString)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
}
