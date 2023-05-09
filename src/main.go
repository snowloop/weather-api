package main

import (
	"context"
	"net/http"
	"weather-api/src/controllers"
	api_utils "weather-api/src/utils"

	"github.com/gorilla/mux"
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
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.ServeHome).Methods("GET")
	router.HandleFunc("/weather/{cityName}", controllers.ServeWeather).Methods("GET")
	router.HandleFunc("/user/weather/{userName}", userController.ServeUserWeather).Methods("GET")
	router.HandleFunc("/user/{userName}", userController.ServeUser).Methods("GET")

	http.Handle("/", router)

	// Serving the router
	http.ListenAndServe(":8080", nil)
}
