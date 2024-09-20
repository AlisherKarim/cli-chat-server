package main

import (
	"log"
	"net/http"
	"os"

	v1Routes "github.com/alisherkarim/cli-chat-server/api/v1/routes"
	"github.com/alisherkarim/cli-chat-server/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type Env struct {
	db db.Storage
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

	
	db, err := db.NewPostgreStorage()
	if err != nil {
		log.Fatal(err.Error())
	}
	
	if err := db.Init(); err != nil {
		log.Fatal(err.Error())
	}
	
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not loaded")
	}

	indexRouter := initRouter()
	// router for paths /v1/*
	indexRouter.Mount("/api/v1", v1Routes.NewRouter(db));

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
