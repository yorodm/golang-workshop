package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
	ID     int    `json:"id,omitempty" bson:"id,omitempty"`
	Title  string `json:"title,omitempty" bson:"title,omitempty"`
	IsDone bool   `json:"isdone,omitempty" bson:"isdone,omitempty"`
}

type MongoDB struct {
	*mongo.Client
	ctx context.Context
}

func NewMongoDB() (*MongoDB, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	return &MongoDB{client, ctx}, nil
}

func (m *MongoDB) Disconnect() {
	defer m.Client.Disconnect(m.ctx)
}

func CreateItem(m *MongoDB, newItem Item) string {

	collection := m.Database("todo").Collection("items")
	result, _ := collection.InsertOne(m.ctx, newItem)

	return result.InsertedID.(primitive.ObjectID).Hex()
}

func UpdateItem(m *MongoDB, item Item) {
	update := bson.M{"$set": bson.M{"title": item.Title, "isdone": item.IsDone}}

	collection := m.Database("todo").Collection("items")
	collection.UpdateOne(m.ctx, Item{ID: item.ID}, update)
}

func GetItems(m *MongoDB) (items []Item) {
	collection := m.Database("todo").Collection("items")
	cursor, _ := collection.Find(m.ctx, bson.M{})

	defer cursor.Close(m.ctx)
	for cursor.Next(m.ctx) {
		var oneItem Item
		cursor.Decode(&oneItem)
		items = append(items, oneItem)
	}

	return
}

func GetItem(m *MongoDB, id int) (item Item) {

	collection := m.Database("todo").Collection("items")
	collection.FindOne(m.ctx, Item{ID: id}).Decode(&item)
	return
}

func DeleteItem(m *MongoDB, id int) {
	collection := m.Database("todo").Collection("items")
	collection.DeleteMany(m.ctx, Item{ID: id})
	return
}
