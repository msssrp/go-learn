package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
}

func NewUserCollection(c *mongo.Client) *mongo.Collection {
	return c.Database("myDatabase").Collection("Users")
}

func (u *User) InsertUser(c *mongo.Client) (*mongo.InsertOneResult, error) {

	u.ID = primitive.NewObjectID()

	collection := NewUserCollection(c)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	res, err := collection.InsertOne(ctx, u)

	if err != nil {
		panic(err)
	}

	return res, nil
}
