package handlers

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/mrkouhadi/go-countries-api/models"
	"github.com/mrkouhadi/go-countries-api/utils"
)

func AllCountires(w http.ResponseWriter, r *http.Request) {
	data, error := utils.GetAllCountries("./data/countries-with-svg_flags.json")
	if error != nil {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: error.Error()})
		return
	}
	utils.EnableCors(&w)
	utils.WriteJSON(w, http.StatusOK, data)
}

func GetCountryByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	// FIXME: This code is depricated and need to be updated
	//
	name = strings.Title(name)
	//
	//
	data, error := utils.GetAllCountries("./data/countries-with-svg_flags.json")
	if error != nil {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: error.Error()})
		return
	}
	requestedCountry := models.CountryType{}
	for _, val := range data {
		if val.Name == name {
			requestedCountry = val
			break
		}
	}
	if requestedCountry.Name == "" {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: "Not found, please double check your parameters"})
		return
	}
	utils.EnableCors(&w)
	utils.WriteJSON(w, http.StatusOK, requestedCountry)
}

// GetCountriesByContinentCode get a list of countries that has the same continent code
func GetCountriesByContinentCode(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "code")
	// FIXME: This code is depricated and need to be updated
	//
	name = strings.ToUpper(name)
	//
	//
	data, error := utils.GetAllCountries("./data/countries-with-svg_flags.json")
	if error != nil {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: error.Error()})
		return
	}
	requestedCountries := []models.CountryType{}
	for _, val := range data {
		if val.Continent_code == name {
			requestedCountries = append(requestedCountries, val)
		}
	}
	if len(requestedCountries) == 0 {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: "Not found, please double check your parameters"})
		return
	}
	utils.EnableCors(&w)
	utils.WriteJSON(w, http.StatusOK, requestedCountries)
}

// GetCountriesByContinentCode get a list of countries that has the same continent code
func GetCountryByCapitalCity(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "capital")
	// FIXME: This code is depricated and need to be updated
	//
	name = strings.Title(name)
	//
	//
	data, error := utils.GetAllCountries("./data/countries-with-svg_flags.json")
	if error != nil {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: error.Error()})
		return
	}
	requestedCountry := models.CountryType{}
	for _, val := range data {
		if val.Capital == name {
			requestedCountry = val
		}
	}
	if requestedCountry.Name == "" {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: "Not found, please double check your parameters"})
		return
	}
	utils.EnableCors(&w)
	utils.WriteJSON(w, http.StatusOK, requestedCountry)
}

