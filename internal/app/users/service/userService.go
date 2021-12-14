package service

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go-mongodb/internal/app/domain"
)

type UserService struct {
	UserRepository domain.UserRepository
	Log            *logrus.Logger
}

func NewUserService(repository domain.UserRepository, log *logrus.Logger) UserService {
	return UserService{
		UserRepository: repository,
		Log:            log,
	}
}

func (s UserService) FindAll() ([]interface{}, error) {
	var users []interface{}

	resp, err := s.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s UserService) FindOne(id string) (interface{}, error) {
	var user interface{}

	resp, err := s.UserRepository.FindOne(id)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s UserService) Create(input map[string]interface{}) (interface{}, error) {
	input["_id"] = uuid.NewString()
	resp, err := s.UserRepository.Create(input)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s UserService) UpdateOne(id string, input map[string]interface{}) error {
	return s.UserRepository.UpdateOne(id, input)
}

func (s UserService) Delete(id string) error {
	return s.UserRepository.Delete(id)
}
