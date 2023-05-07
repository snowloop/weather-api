package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	weather_api_service "weather-api/src/services"
)

func ServeWeather(w http.ResponseWriter, r *http.Request) {

	city := strings.SplitN(r.URL.Path, "/", 3)[2]

	apiResponse, err := weather_api_service.GetWeatherFromCity(city)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(apiResponse)
}
