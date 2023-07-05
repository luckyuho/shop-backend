package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	// 设置MongoDB连接选项
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接是否成功
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// 获取对特定数据库的句柄
	database := client.Database("mydatabase")

	// 获取对特定集合的句柄
	collection = database.Collection("mycollection")

	// // 插入文档
	// doc := bson.D{
	// 	{"name", "John"},
	// 	{"age", 30},
	// 	{"city", "New York"},
	// }

	// insertResult, err := collection.InsertOne(context.TODO(), doc)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 获取插入的文档ID
	// insertedID := insertResult.InsertedID
	// fmt.Println("Inserted document ID:", insertedID)

	// // 查询文档
	// filter := bson.D{{"name", "John"}}

	// var result bson.M
	// err = collection.FindOne(context.TODO(), filter).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Found document:", result)

	// // 更新文档
	// update := bson.D{{"$set", bson.D{{"age", 31}}}}

	// updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Updated", updateResult.ModifiedCount, "document(s)")

	// // 删除文档
	// deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Deleted", deleteResult.DeletedCount, "document(s)")

	// // 断开与MongoDB的连接
	// err = client.Disconnect(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Disconnected from MongoDB!")
}

func GetMongo() *mongo.Collection {
	return collection
}
