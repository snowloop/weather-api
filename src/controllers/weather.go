package controllers

import (
	"encoding/json"
	"net/http"
	weather_api_service "weather-api/src/services"

	"github.com/gin-gonic/gin"
)

func ServeWeather(gin_context *gin.Context) {
	cityName := gin_context.Param("cityName")
	apiResponse, err := weather_api_service.GetWeatherFromCity(cityName)

	if err != nil {
		http.Error(gin_context.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	gin_context.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(gin_context.Writer).Encode(apiResponse)
}
