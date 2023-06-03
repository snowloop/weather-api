package main

import (
	"context"
	"weather-api/src/controllers"
	api_utils "weather-api/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	apiConfig, err := api_utils.LoadApiConfig("")
	if err != nil {
		panic(err)
	}

	mongoClient := api_utils.InitMongoClient(&apiConfig)

	defer mongoClient.Disconnect(context.Background())

	mongoDatabase := mongoClient.Database(apiConfig.DatabaseName)

	// Creating the user controller and injecting the dependencies
	userController := controllers.NewUserController(mongoDatabase)

	// Creating the mux router
	router := gin.Default()
	router.GET("", controllers.ServeHome)
	router.GET("/weathers/:cityName", controllers.ServeWeather)
	router.GET("/users/weather/:userName", userController.ServeUserWeather)
	router.GET("/users/:userName", userController.ServeUser)

	router.Run()
}
