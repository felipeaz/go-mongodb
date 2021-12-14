package domain

import (
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Id        string `json:"_id" bson:"_id"`
	Email     string `json:"email,omitempty" bson:"email"`
	FirstName string `json:"firstName,omitempty" bson:"firstName"`
	LastName  string `json:"lastName,omitempty" bson:"lastName"`
	Phone     string `json:"phone,omitempty" bson:"phone"`
}

func (u User) GetBson() bson.M {
	return bson.M{
		"_id":       u.Id,
		"email":     u.Email,
		"firstName": u.FirstName,
		"lastName":  u.LastName,
		"phone":     u.Phone,
	}
}

type UserService interface {
	FindAll() ([]User, error)
	FindOne(id string) (*User, error)
	Create(user User) (*User, error)
	UpdateOne(id string, user User) error
	Delete(id string) error
}

type UserRepository interface {
	FindAll() ([]byte, error)
	FindOne(id string) ([]byte, error)
	Create(user User) (*User, error)
	UpdateOne(id string, user User) error
	Delete(id string) error
}
