package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrkouhadi/go-countries-api/handlers"
)

func main() {
	// set up a router
	router := chi.NewRouter()

	// cities
	router.Get("/cities", handlers.Allcities)
	router.Get("/cities/{name}", handlers.GetCityByName)
	router.Get("/cities/by-country/{name}", handlers.GetCitiesByCountryName)
	router.Get("/cities/by-geonameid/{name}", handlers.GetCityByGeoNameID)

	//countries
	router.Get("/countries", handlers.AllCountires)
	router.Get("/countries/{name}", handlers.GetCountryByName)
	router.Get("/countries/by-country/{name}", handlers.GetCitiesByCountryName)
	router.Get("/countries/by-continent-code/{code}", handlers.GetCountriesByContinentCode)
	router.Get("/countries/by-capital-city/{capital}", handlers.GetCountryByCapitalCity)

	// run the server
	http.ListenAndServe(":8080", router)
}
