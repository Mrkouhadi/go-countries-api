package models

type CityType struct {
	Name       string `json:"city_name"`
	Country    string `json:"country_name"`
	Subcountry string `json:"province_Or_state"`
	Geonameid  string `json:"geonameid"`
}
type CountryType struct {
	Name                      string       `json:"country_name"`
	Code                      string       `json:"country_code"`
	Capital                   string       `json:"country_capital_city"`
	Continent                 string       `json:"continent_name"`
	Continent_code            string       `json:"continent_code"`
	Phone_code                string       `json:"phone_code"`
	Three_letter_country_code string       `json:"three_letter_country_code"`
	Currency                  CurrencyType `json:"currency"`
	Language                  LanguageType `json:"language"`
	Flag                      string       `json:"flag"`
}
type CurrencyType struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
type LanguageType struct {
	Code       string `json:"code"`
	Iso639_2   string `json:"Iso639_2"`
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

type ApiError struct {
	Message string `json:"error"`
}
