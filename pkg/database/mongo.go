package database

import (
	"context"
	"strconv"

	"github.com/kev1nandreas/go-rest-api-template/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongoDB() *mongo.Collection {
	mongo_host := env.GetEnvString("MONGO_HOST", "localhost")
	mongo_port := env.GetEnvInt("MONGO_PORT", 27017)

	mongo_database := env.GetEnvString("MONGO_DATABASE", "logging")
	mongo_collection := env.GetEnvString("MONGO_COLLECTION", "logs")

	marshalledURI := "mongodb://" + mongo_host + ":" + strconv.Itoa(mongo_port)

	clientOptions := options.Client().ApplyURI(marshalledURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	return client.Database(mongo_database).Collection(mongo_collection)
}
