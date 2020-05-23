package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURL = "mongodb+srv://mongol:m0ng0l@cluster0-gt8lu.mongodb.net/test?retryWrites=true&w=majority"
)

//AddNew would insert the new document into the alert collection
func AddNew(doc *GrafanaBody) interface{} {

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

	res, err := alertsCollection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
	oid := res.InsertedID.(primitive.ObjectID).Hex()
	return oid

}

//Delete would insert the new document into the alert collection
func Delete() {

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

	if err := alertsCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}
}
