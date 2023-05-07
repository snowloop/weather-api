package api_utils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type apiConfigData struct {
	OpenWeatherApiKey     string `json:"OpenWeatherApiKey`
	MongoConnectionString string `json:"MongoConnectionString"`
	DatabaseName          string `json:"DatabaseName"`
}

func LoadApiConfig(filename string) (apiConfigData, error) {
	if filename == "" {
		filename = ".env"
	}
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

func InitMongoClient(apiConfig *apiConfigData) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOption := options.Client().ApplyURI(apiConfig.MongoConnectionString)

	mongoClient, err := mongo.Connect(ctx, clientOption)

	if err != nil {
		panic(err)
	} else {
		fmt.Print("Connected to the database!")
	}

	return mongoClient
}
