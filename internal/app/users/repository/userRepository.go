package repository

import (
	"context"
	"encoding/json"
	"go-mongodb/internal/app/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return UserRepository{
		Collection: collection,
	}
}

func (r UserRepository) FindAll() ([]byte, error) {
	cursor, err := r.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}

	return json.MarshalIndent(result, "", "	")
}

func (r UserRepository) FindOne(id string) ([]byte, error) {
	return nil, nil
}

func (r UserRepository) Create(user domain.User) (*domain.User, error) {
	_, err := r.Collection.InsertOne(context.TODO(), user.GetBson())
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepository) UpdateOne(id string, user domain.User) error {
	return nil
}

func (r UserRepository) Delete(id string) error {
	return nil
}
