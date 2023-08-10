package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/countries", AllCountires)
	http.HandleFunc("/cities", Allcities)

	http.ListenAndServe(":8080", nil)
}
