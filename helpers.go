package main

import (
	"encoding/json"
	"net/http"
	"os"
)

// getAll gets data from the json files
func getAll(fname string) ([]CountryType, error) {
	file, _ := os.ReadFile(fname)
	countries := []CountryType{}
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

// enabling CORS
func enableCors(w *http.ResponseWriter) {
	// WE CHOOSE ONLY ONE OF THESE BELOW NOT BOTH...
	// (*w).Header().Set("Access-Control-Allow-Origin", "*") // this means open all possible origins
	(*w).Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500") // open to only http://localhost:3000
}
