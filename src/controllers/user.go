package controllers

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"

var collection *mongo.Collection

func initMongoClient() {
	clientOption := options.Client().ApplyURI(connectionString)

	mongo.Connect(context.TODO(), clientOption)
}
