package repository

import (
	"context"

	"github.com/waltherx/honda-backend/domain"
	"github.com/waltherx/honda-backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type clientRepository struct {
	database   mongo.Database
	collection string
}

func NewClientRepository(db mongo.Database, collection string) domain.ClientRepository {
	return &clientRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *clientRepository) Create(c context.Context, client *domain.Client) error {
	collection := tr.database.Collection(tr.collection)

	_, err := collection.InsertOne(c, client)

	return err
}

func (tr *clientRepository) FetchByUserID(c context.Context, userID string) ([]domain.Client, error) {
	collection := tr.database.Collection(tr.collection)

	var clients []domain.Client

	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return clients, err
	}

	cursor, err := collection.Find(c, bson.M{"userID": idHex})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &clients)
	if clients == nil {
		return []domain.Client{}, err
	}

	return clients, err
}
