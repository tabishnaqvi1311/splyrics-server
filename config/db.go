package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func ConnectToMongo() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(getMongoUri()))

	if err != nil{
		log.Fatal(err)
	}

	return client
}

var DB *mongo.Client = ConnectToMongo()

func GetCollection(client *mongo.Client) *mongo.Collection{
	coll := client.Database("Spotify-Lyrics").Collection("lyrics")
	return coll
}