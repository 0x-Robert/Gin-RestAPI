package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/bson"
)

type Person struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
	Pnum string `bson:"pnum"`
}

func main() {
	newPerson := Person{Name: "inTest3", Age: 55, Pnum: "0318894561"}
	Mongo_URL := "mongodb://127.0.0.1:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))
	//db := client.Database("go-ready") // database 접속
	//col := db.Collection("tPerson")   // collection 접속

	// collection으로 바로 접근
	coll := client.Database("go-ready").Collection("tPerson")
	result, err := coll.InsertOne(context.TODO(), newPerson)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
}
