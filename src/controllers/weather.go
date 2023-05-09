package controllers

import (
	"encoding/json"
	"net/http"
	weather_api_service "weather-api/src/services"

	"github.com/gorilla/mux"
)

func ServeWeather(w http.ResponseWriter, r *http.Request) {
	requestParams := mux.Vars(r)
	apiResponse, err := weather_api_service.GetWeatherFromCity(requestParams["cityName"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(apiResponse)
}
