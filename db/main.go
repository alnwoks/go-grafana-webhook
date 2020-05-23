package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	mongoURL := "mongodb+srv://mongol:m0ng0l@cluster0-gt8lu.mongodb.net/test?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}

	//create timeout context
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	//connecting client to mongo db with timeout context
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//ping mongo db
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("pong.")
	}

	grafanaDatabase := client.Database("grafana")
	messageCollection := grafanaDatabase.Collection("messages")

	result, err := messageCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "The Polyglot Developer Podcast"},
		{Key: "author", Value: "Nic Raboy"},
	})
	fmt.Println("Result: ", *result)

	//get list of databases
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(databases)
	}

}

//AddNew would insert the new document into the alert collection
func AddNew(doc *AlertDocument) {
	mongoURL := "mongodb+srv://mongol:m0ng0l@cluster0-gt8lu.mongodb.net/test?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatal(err)
	}

	//create timeout context
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	//connecting client to mongo db with timeout context
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	grafanaDatabase := client.Database("grafana")
	alertsCollection := grafanaDatabase.Collection("alerts")

	alertresult, err := alertsCollection.InsertOne(ctx, doc)

	fmt.Println("ID: ", alertresult.InsertedID)
}
