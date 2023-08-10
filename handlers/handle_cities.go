package handlers

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/mrkouhadi/go-countries-api/models"
	"github.com/mrkouhadi/go-countries-api/utils"
)

// ////////////////////////////////////////////////////// the whole list of cities
func Allcities(w http.ResponseWriter, r *http.Request) {
	data, error := utils.GetAllCities("./data/cities.json")
	if error != nil {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: error.Error()})
		return
	}
	utils.WriteJSON(w, http.StatusOK, data)
}

// ////////////////////////////////////////////////////// single items

func GetCityByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	//FIXME:
	//
	name = strings.Title(name)
	//
	//
	data, error := utils.GetAllCities("./data/cities.json")
	if error != nil {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: error.Error()})
		return
	}
	requestedCity := models.CityType{}
	for _, val := range data {
		if val.Name == name {
			requestedCity = val
			break
		}
	}
	if requestedCity.Name == "" {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: "Not found, please double check your parameters"})
		return
	}
	utils.EnableCors(&w)
	utils.WriteJSON(w, http.StatusOK, requestedCity)
}

// ////////////////////////////////////////////////////// get a city by continent name
func GetCitiesByCountryName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	//FIXME:
	//
	name = strings.Title(name)
	//
	//
	data, error := utils.GetAllCities("./data/cities.json")
	if error != nil {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: error.Error()})
		return
	}
	requestedCities := []models.CityType{}
	for _, val := range data {
		if val.Country == name {
			requestedCities = append(requestedCities, val)
		}
	}
	if len(requestedCities) == 0 {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: "Not found, please double check your parameters"})
		return
	}
	utils.EnableCors(&w)
	utils.WriteJSON(w, http.StatusOK, requestedCities)
}

func GetCityByGeoNameID(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	//FIXME:
	//
	name = strings.Title(name)
	//
	//
	data, error := utils.GetAllCities("./data/cities.json")
	if error != nil {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: error.Error()})
		return
	}
	requestedCities := models.CityType{}
	for _, val := range data {
		if val.Geonameid == name {
			requestedCities = val
		}
	}
	if requestedCities.Name == "" {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: "Not found, please double check your parameters"})
		return
	}
	utils.EnableCors(&w)
	utils.WriteJSON(w, http.StatusOK, requestedCities)
}
