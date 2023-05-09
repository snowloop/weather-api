package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	weather_api_service "weather-api/src/services"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

const userCollectionName = "users"

type UserController struct {
	collection *mongo.Collection
}

func NewUserController(mongoDatabase *mongo.Database) *UserController {
	return &UserController{mongoDatabase.Collection(userCollectionName)}
}

func (c *UserController) ServeUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	requestParams := mux.Vars(r)
	var foundedUser User
	findResult := c.collection.FindOne(context.Background(),
		bson.M{"name": requestParams["userName"]})

	err := findResult.Decode(&foundedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, err.Error(), http.StatusNotFound)
			return

		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(foundedUser)
}

type WeatherUser struct {
	User    User                            `json:"user"`
	Weather weather_api_service.WeatherData `json:"weather"`
}

func (c *UserController) ServeUserWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	userName := mux.Vars(r)["userName"]
	fmt.Print(userName)
	var foundedUser User

	err := c.collection.FindOne(context.Background(), bson.M{"name": userName}).Decode(&foundedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, err.Error(), http.StatusNotFound)
			return

		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print("ici")

	userWeather, err := weather_api_service.GetWeatherFromCity(foundedUser.Location)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(WeatherUser{User: foundedUser, Weather: userWeather})

}
