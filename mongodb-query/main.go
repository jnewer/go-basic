package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func initDB() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	co := options.Client().ApplyURI("mongodb://localhost:27017")
	mongo.Connect(context.TODO(), co)
	// 连接到MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	client.Ping(context.TODO(), nil)
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func find() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := client.Database("go_db").Collection("student")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("result: %v\n", result)
		fmt.Printf("result.Map(): %v\n", result.Map()["name"])
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB()
	find()
}
