package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	weather_api_service "weather-api/src/services"

	"github.com/gin-gonic/gin"
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

func (c *UserController) ServeUser(gin_context *gin.Context) {
	gin_context.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	user_name := gin_context.Param("userName")
	var foundedUser User
	findResult := c.collection.FindOne(context.Background(),
		bson.M{"name": user_name})

	err := findResult.Decode(&foundedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(gin_context.Writer, err.Error(), http.StatusNotFound)
			return

		}
		http.Error(gin_context.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(gin_context.Writer).Encode(foundedUser)
}

type WeatherUser struct {
	User    User                            `json:"user"`
	Weather weather_api_service.WeatherData `json:"weather"`
}

func (c *UserController) ServeUserWeather(gin_context *gin.Context) {
	gin_context.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	userName := gin_context.Param("userName")
	fmt.Print(userName)
	var foundedUser User

	err := c.collection.FindOne(context.Background(), bson.M{"name": userName}).Decode(&foundedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(gin_context.Writer, err.Error(), http.StatusNotFound)
			return

		}
		http.Error(gin_context.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print("ici")

	userWeather, err := weather_api_service.GetWeatherFromCity(foundedUser.Location)

	if err != nil {
		http.Error(gin_context.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(gin_context.Writer).Encode(WeatherUser{User: foundedUser, Weather: userWeather})

}
