package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/mrkouhadi/go-countries-api/handlers"
	"github.com/mrkouhadi/go-countries-api/models"
	"github.com/mrkouhadi/go-countries-api/utils"
	"golang.org/x/time/rate"
)

func main() {
	// set up a router
	router := chi.NewRouter()
	router.Use(RateLimiter)
	// cities
	router.Get("/cities", handlers.Allcities)
	router.Get("/cities/{name}", handlers.GetCityByName)
	router.Get("/cities/by-country/{name}", handlers.GetCitiesByCountryName)
	router.Get("/cities/by-geonameid/{name}", handlers.GetCityByGeoNameID)

	//countries
	router.Get("/countries", handlers.AllCountries)
	router.Get("/countries/{name}", handlers.GetCountryByName)
	router.Get("/countries/by-continent-code/{code}", handlers.GetCountriesByContinentCode)
	router.Get("/countries/by-capital-city/{capital}", handlers.GetCountryByCapitalCity)

	// setting up env vars
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// run the server
	http.ListenAndServe(":"+port, router)
}

// middlewares
func RateLimiter(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(1, 3) // max of 3 requests and then 1 more requests per second
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			message := models.ApiError{
				Message: "Request Failed! Because The API is at capacity, try again later.",
			}
			utils.WriteJSON(w, http.StatusTooManyRequests, message)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
