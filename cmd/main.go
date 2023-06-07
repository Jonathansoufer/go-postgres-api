package main

import (
	"fmt"
	"net/http"

	"github.com/Jonathansoufer/go-postgres-api/configs"
	"github.com/Jonathansoufer/go-postgres-api/handlers"
	"github.com/go-chi/chi"
)

func main(){
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Get("/todos", handlers.List)
	r.Get("/todos/{id}", handlers.Get)
	r.Post("/todos", handlers.Create)
	r.Put("/todos/{id}", handlers.Update)
	r.Delete("/todos/{id}", handlers.Delete)


	http.ListenAndServe(fmt.Sprintf(":%s",configs.GetAPI().Port), r)
}