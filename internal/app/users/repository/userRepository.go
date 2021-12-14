package repository

import (
	"context"
	"encoding/json"
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
	var result bson.M
	if err := r.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result); err != nil {
		return nil, err
	}
	return json.MarshalIndent(result, "", "	")
}

func (r UserRepository) Create(input map[string]interface{}) (interface{}, error) {
	_, err := r.Collection.InsertOne(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r UserRepository) UpdateOne(id string, input map[string]interface{}) error {
	_, err := r.Collection.UpdateOne(context.TODO(), bson.M{"_id": id}, input)
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepository) Delete(id string) error {
	_, err := r.Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
