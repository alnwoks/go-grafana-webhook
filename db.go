package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//AddNew would insert the new document into the alert collection
func AddNew(doc *GrafanaBody) {
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

	fmt.Println(alertresult)

}
