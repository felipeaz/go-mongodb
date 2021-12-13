package repository

import (
	"go-mongodb/internal/app/domain"
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
	return nil, nil
}

func (r UserRepository) FindOne(id string) ([]byte, error) {
	return nil, nil
}

func (r UserRepository) Create(user domain.User) ([]byte, error) {
	return nil, nil
}

func (r UserRepository) UpdateOne(id string, user domain.User) error {
	return nil
}

func (r UserRepository) Delete(id string) error {
	return nil
}
