package api_utils

import (
	"encoding/json"
	"os"
)

type apiConfigData struct {
	OpenWeatherApiKey string `json:"OpenWeatherApiKey`
}

func LoadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}
	var c apiConfigData
	err = json.Unmarshal(bytes, &c)

	if err != nil {
		return apiConfigData{}, err
	}

	return c, nil
}
