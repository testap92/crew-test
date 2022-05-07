package mongoClient

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoClientInterface interface {
	Storage(collection string) *mongo.Collection
}

type MongoClient struct {
	DBName string
	Cli    *mongo.Client
}

func (r *MongoClient) Storage(collection string) *mongo.Collection {
	return r.Cli.Database(r.DBName).Collection(collection)
}

func Connect(dbName, DSN string) *MongoClient {
	client, err := mongo.NewClient(options.Client().ApplyURI(DSN))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoClient{
		DBName: dbName,
		Cli:    client,
	}
}
