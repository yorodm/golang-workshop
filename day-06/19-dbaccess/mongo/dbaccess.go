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
}

func NewMongoDB(ctx context.Context) (*MongoDB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	return &MongoDB{client}, nil
}

func (m *MongoDB) Disconnect(ctx context.Context) {
	defer m.Client.Disconnect(ctx)
}

func (m *MongoDB) CreateItem(ctx context.Context, newItem Item) string {

	collection := m.Database("todo").Collection("items")
	result, _ := collection.InsertOne(ctx, newItem)

	return result.InsertedID.(primitive.ObjectID).Hex()
}

func (m *MongoDB) UpdateItem(ctx context.Context, item Item) {
	update := bson.M{"$set": bson.M{"title": item.Title, "isdone": item.IsDone}}

	collection := m.Database("todo").Collection("items")
	collection.UpdateOne(ctx, Item{ID: item.ID}, update)
}

func (m *MongoDB) GetItems(ctx context.Context) (items []Item) {
	collection := m.Database("todo").Collection("items")
	cursor, _ := collection.Find(ctx, bson.M{})

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var oneItem Item
		cursor.Decode(&oneItem)
		items = append(items, oneItem)
	}

	return
}

func (m *MongoDB) GetItem(ctx context.Context, id int) (item Item) {

	collection := m.Database("todo").Collection("items")
	collection.FindOne(ctx, Item{ID: id}).Decode(&item)
	return
}

func (m *MongoDB) DeleteItem(ctx context.Context, id int) {
	collection := m.Database("todo").Collection("items")
	collection.DeleteMany(ctx, Item{ID: id})
	return
}
