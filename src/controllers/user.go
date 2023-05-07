package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name       string `json:"name"`
	locationId string
}

type Location struct {
	Name string
}

const userCollectionName = "collection1"

type UserController struct {
	mongoDatabase *mongo.Database
}

func (c *UserController) ServeUser(w http.ResponseWriter, r *http.Request) {

	findResult := c.mongoDatabase.Collection(userCollectionName).FindOne(context.Background(), bson.M{"myField": "Alex"})
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(findResult)
}

func NewUserController(mongoDatabase *mongo.Database) *UserController {
	return &UserController{mongoDatabase}
}
