package weather_api_service

import (
	"encoding/json"
	"fmt"
	"net/http"
	api_utils "weather-api/src/utils"
)

type weatherData struct {
	Coord struct {
		Lon float64 `json:"lon`
		Lat float64 `json:"lat`
	} `json:"coord`
	Wind struct {
		Speed float64 `json:"speed`
		Deg   float64 `json:"deg`
	} `json:"wind`
	Name string `json:"name`
	Main struct {
		Temp       float64 `json:"temp`
		Feels_Like float64 `json:"feels_like`
		Temp_Min   float64 `json:"temp_min`
		Temp_Max   float64 `json:"temp_max`
		Pressure   float64 `json:"pressure`
		Humidity   float64 `json:"humidity`
	} `json:"main`
}

func GetWeatherFromCity(cityname string) (weatherData, error) {
	apiConfigData, err := api_utils.LoadApiConfig(".env")
	fmt.Println(apiConfigData.OpenWeatherApiKey)
	if err != nil {
		return weatherData{}, err
	}
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?appid=" + apiConfigData.OpenWeatherApiKey + "&units=metric" + "&q=" + cityname)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var response weatherData
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return weatherData{}, err
	}
	return response, nil
}
