package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionClient = "client"
)

type Client struct {
	ID       primitive.ObjectID `bson:"_id" json:"-"`
	Ci       string             `bson:"ci" form:"ci" binding:"required" json:"ci"`
	FullName string             `bson:"fullname" form:"fullname" binding:"required" json:"fullname"`
	Address  string             `bson:"address" form:"address" json:"address"`
	Phone    string             `bson:"phone" form:"phone" json:"phone"`
	//Created_at
	UserID primitive.ObjectID `bson:"userID" json:"-"`
}

type ClientRepository interface {
	Create(c context.Context, client *Client) error
	FetchByUserID(c context.Context, userID string) ([]Client, error)
}

type ClientUsecase interface {
	Create(c context.Context, client *Client) error
	FetchByUserID(c context.Context, userID string) ([]Client, error)
}
