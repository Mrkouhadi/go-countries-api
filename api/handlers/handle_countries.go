package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/mrkouhadi/go-countries-api/models"
	"github.com/mrkouhadi/go-countries-api/utils"
)

func AllCountires(w http.ResponseWriter, r *http.Request) {
	fmt.Println(chi.URLParam(r, "id")) // get the ID
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
