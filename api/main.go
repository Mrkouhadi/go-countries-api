package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrkouhadi/go-countries-api/handlers"
)

func main() {
	router := chi.NewRouter()

	router.Get("/countries", handlers.AllCountires)
	router.Get("/countries/{name}", handlers.GetCountryByName)
	router.Get("/cities/{name}", handlers.GetCityByName)

	router.Get("/cities", handlers.Allcities)

	http.ListenAndServe(":8080", router)
}
