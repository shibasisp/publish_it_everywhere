package main

import (
	"fmt"
	"net/http"
	"os"

	"publish_it_everywhere/api/linkedin"
	"publish_it_everywhere/config"
	"publish_it_everywhere/db"
	appmiddleware "publish_it_everywhere/middleware"

	log "github.com/sirupsen/logrus"

	"publish_it_everywhere/api"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func main() {
	// load all config/env's
	config.Initialize(os.Args[1:]...)
	linkedin.Initialize()
	db.Initialize()
	router := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
		AllowedHeaders: []string{
			"Origin", "Authorization", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Header", "Accept",
			"Content-Type", "X-CSRF-Token",
		},
		ExposedHeaders: []string{
			"Content-Length", "Access-Control-Allow-Origin", "Origin",
		},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// cross & loger middleware
	router.Use(cors.Handler)
	router.Use(
		appmiddleware.LogHandler,
		middleware.Logger,
		appmiddleware.Recoverer,
	)
	// router.Get("/", api.IndexHandeler)
	router.Route("/api", api.Init)

	log.Infoln("Starting server on port:", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router)
}
