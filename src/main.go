package main

import (
	"net/http"
	"weather-api/src/controllers"

	"github.com/gorilla/mux"
)

type User struct {
	Name       string `json:"fullName"`
	locationId string `json:"locationId"`
}

type Location struct {
	Name string `json:"name"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.ServeHome)
	router.HandleFunc("/weather", controllers.ServeWeather)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
