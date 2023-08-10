package main

import (
	"net/http"
)

// ////////////////////////////////////////////////////// the whole list

func AllCountires(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, error := getAll("./data/countries-with-svg_flags.json")
		if error != nil {
			WriteJSON(w, http.StatusNotFound, ApiError{Message: error.Error()})
			return
		}
		enableCors(&w)
		WriteJSON(w, http.StatusOK, data)
	} else {
		WriteJSON(w, http.StatusMethodNotAllowed, ApiError{Message: "method not allowed"})
	}
}

func Allcities(w http.ResponseWriter, r *http.Request) {
	data, error := getAll("./data/cities.json")
	if error != nil {
		WriteJSON(w, http.StatusNotFound, ApiError{Message: error.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, data)
}

// ////////////////////////////////////////////////////// single item
// func GetCountryByID(w http.ResponseWriter, r *http.Request) {

// }
// func GetCityByID(w http.ResponseWriter, r *http.Request) {

// }
