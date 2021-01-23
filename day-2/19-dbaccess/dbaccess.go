package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"	
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
	ID     int    `json:"id,omitempty bson:"id,omitempty"`
	Title  string `json:"title,omitempty" bson:"title,omitempty"`
	IsDone bool   `json:"isdone,omitempty" bson:"isdone,omitempty"`
}

func CreateItem(newItem Item) string {
	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	result, _ := collection.InsertOne(ctx, newItem)

	disconnect(ctx,client)

	return result.InsertedID.(primitive.ObjectID).Hex()
}

func UpdateItem(item Item) {
	update := bson.M{"$set": bson.M{"title": item.Title,"isdone": item.IsDone }}

	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	collection.UpdateOne(ctx, Item{ID: item.ID}, update)

	disconnect(ctx,client)
}

func GetItems() (items []Item) {
	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	cursor, _ := collection.Find(ctx, bson.M{})

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var oneItem Item
		cursor.Decode(&oneItem)
		items = append(items, oneItem)
	}

	disconnect(ctx,client)

	return
}

func GetItem(id int) (item Item) {
	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	collection.FindOne(ctx, Item{ID: id}).Decode(&item)
	
	disconnect(ctx,client)

	return
}

func DeleteItem(id int) {
	ctx, client := connnect()

	collection := client.Database("todo").Collection("items")
	collection.DeleteMany(ctx, Item{ID: id})
	
	disconnect(ctx,client)

	return
}

func connnect() (context.Context, *mongo.Client) {
	ctx := context.Background()
	client, _ := mongo.Connect(ctx,options.Client().ApplyURI("mongodb://localhost:27017"))

	return ctx, client
}

func disconnect(ctx context.Context, client *mongo.Client) {
	defer client.Disconnect(ctx)
}
