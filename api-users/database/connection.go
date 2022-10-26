package database

import (
	"api-users/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectDB() *mongo.Database {
	ctx := context.TODO()
	conf := config.GetConfig()

	connection := options.Client().ApplyURI(conf.Mongo.Server)

	client, err := mongo.Connect(ctx, connection)
	if err != nil {
		panic(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return client.Database(conf.Mongo.Database)
}

func UserCollection(nameColl string) *mongo.Collection {
	coll := ConnectDB().Collection(nameColl)
	return coll
}
