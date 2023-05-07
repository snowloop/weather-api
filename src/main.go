package main

import (
	"context"
	"fmt"
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
	fmt.Print(mongoDatabase)

	// Creating the user controller and injecting the dependencies
	userController := controllers.NewUserController(mongoDatabase)

	// Creating the mux router
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.ServeHome)
	router.HandleFunc("/weather", controllers.ServeWeather)
	router.HandleFunc("/user", userController.ServeUser)
	http.Handle("/", router)

	// Serving the router
	http.ListenAndServe(":8080", nil)
}
