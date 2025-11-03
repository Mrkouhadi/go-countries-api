package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/mrkouhadi/go-countries-api/models"
)

// getAll gets data from the json files
func GetAllCountries(fname string) ([]models.CountryType, error) {
	file, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	countries := []models.CountryType{}
	err = json.Unmarshal(file, &countries)
	if err != nil {
		return nil, err
	}
	// Filter out empty countries
	var cleanCountries []models.CountryType
	for _, country := range countries {
		if country.Code != "" && country.Name != "" {
			cleanCountries = append(cleanCountries, country)
		}
	}
	return cleanCountries, nil
}

func GetAllCities(fname string) ([]models.CityType, error) {
	file, _ := os.ReadFile(fname)
	countries := []models.CityType{}
	err := json.Unmarshal(file, &countries)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

// WriteJSON helps writing json data to the client side
func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// Enabling CORS
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*") // * means open all possible origins
	// (*w).Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500") // open to only http://127.0.0.1:5500
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
}
