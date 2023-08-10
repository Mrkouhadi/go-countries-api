package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/mrkouhadi/go-countries-api/models"
	"github.com/mrkouhadi/go-countries-api/utils"
)

// ////////////////////////////////////////////////////// the whole list

func Allcities(w http.ResponseWriter, r *http.Request) {
	data, error := utils.GetAllCities("./data/cities.json")
	if error != nil {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: error.Error()})
		return
	}
	utils.WriteJSON(w, http.StatusOK, data)
}

// ////////////////////////////////////////////////////// single item

func GetCityByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	//FIXME:
	//
	name = strings.Title(name)
	//
	//
	fmt.Println(name)
	data, error := utils.GetAllCities("./data/cities.json")
	if error != nil {
		utils.WriteJSON(w, http.StatusNotFound, models.ApiError{Message: error.Error()})
		return
	}
	requestedCountry := models.CityType{}
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
