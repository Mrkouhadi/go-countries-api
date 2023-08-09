package main

import (
	"encoding/json"
	"os"
)

func MergeCountriesWithCities() {
	// get data from ./data/countries.json
	file, _ := os.ReadFile("./data/countries.json")
	countriesData := []NewCountry{}
	_ = json.Unmarshal(file, &countriesData)

	// get data from ./data/cities.json
	NewfileCities, _ := os.ReadFile("./data/cities.json")
	CitiesData := []CityJSON{}
	_ = json.Unmarshal(NewfileCities, &CitiesData)

	allData := []NewCountryWithCities{}

	for i := 0; i < len(countriesData); i++ {
		NewcountryWithCities := NewCountryWithCities{
			Name:                      countriesData[i].Name,
			Code:                      countriesData[i].Code,
			Capital:                   countriesData[i].Capital,
			Continent_name:            countriesData[i].Continent_name,
			Continent_code:            countriesData[i].Continent_code,
			Phone_code:                countriesData[i].Phone_code,
			Three_letter_country_code: countriesData[i].Three_letter_country_code,
			Currency:                  countriesData[i].Currency,
			Language:                  countriesData[i].Language,
			Flag:                      countriesData[i].Flag,
			// Cities: ,
		}
		for j := 0; j < len(CitiesData); j++ {
			if countriesData[i].Name == CitiesData[j].Country {
				NewcountryWithCities.Cities = append(NewcountryWithCities.Cities, CitiesData[j])
			}
		}
		allData = append(allData, NewcountryWithCities)
	}

	// convert the data from a struct type to json
	jsonfile, _ := json.MarshalIndent(allData, "", "    ")
	// insert it to a json file
	_ = os.WriteFile("./data/countries-with-cities.json", jsonfile, 0644)
}
