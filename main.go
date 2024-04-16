package main

import (
	"fmt"
	"net/http"
	"todolist-api/configs"
	"todolist-api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedHeaders:   []string{"Content-Type", "Accept"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	r.Use(c.Handler)

	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
