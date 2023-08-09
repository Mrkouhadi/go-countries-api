package main

import (
	"encoding/json"
	"os"
	"strings"
)

func main() {
	EditCountriesFlags()
}

func EditCountriesFlags() {
	file, _ := os.ReadFile("./data/countries-with-cities.json")
	countries := []NewCountryWithCities{}
	_ = json.Unmarshal(file, &countries)

	for i := 0; i < len(countries); i++ {
		cc := strings.ToLower(countries[i].Code)
		countries[i].Flag = "./flags-svg/" + cc + ".svg"
	}
	// write a file of json format
	jsonfile, _ := json.MarshalIndent(countries, "", "    ")
	_ = os.WriteFile("./data/countries-with-svg-flags-cities.json", jsonfile, 0644)
}
