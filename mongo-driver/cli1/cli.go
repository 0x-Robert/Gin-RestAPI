package main

import (
	"context"
	"fmt"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Mongo_URL := "mongodb://127.0.0.1:27017"
	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))
	// if err == nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("client", client)

	//mongodb password가 설정되있는경우
	// credential := options.Credential{
	// 	Username: "codz",
	// 	Password: "states",
	// }

	// connect := func(dataSource string) (*mongo.Client, error) {
	// 	if client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dataSource).SetAuth(credential)); err != nil {
	// 		return nil, err
	// 	} else if err = client.Ping(context.Background(), nil); err != nil { //커넥션 유지를 위한 Ping
	// 		return nil, err
	// 	} else {
	// 		return client, nil
	// 	}
	// }

	// client, err := connect("mongodb://127.0.0.1:27017")
	// if err == nil {
	// 	fmt.Println(err)
	// }

	// //db -> collection 단계적 접근
	// db := client.Database("go-ready") // database 접속
	// col := db.Collection("tPerson")   // collection 접속

	// // collection으로 바로 접근
	// coll := client.Database("go-ready").Collection("tPerson")

	// // 한 DB내 여러 collection 접근방식
	// // db := client.Database("go-ready")
	// // colPerson := db.Collection("tPerson")
	// // colStudent := db.Collection("tStudent")
	// // colWoman := db.Collection("tWoman")
	// fmt.Println(col, coll)

}
